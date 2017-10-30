package cmd_run

import (
	"../api_v1"
	"../utils"
	"errors"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"os"
	"time"
)

type CommandRun struct {
	/**
	 * アーティファクト出力パス
	 * len() == 0 の場合、アーティファクトは回収しない
	 */
	ArtifactPath string

	/**
	 * リポジトリURL
	 */
	Repository string

	/**
	 * リポジトリリビジョン
	 */
	Revision string

	/**
	 * for github
	 * リポジトリアクセストークン
	 */
	GithubAccessToken string

	/**
	 * timebooster.yaml path
	 */
	ConfigPath string

	/**
	 * timebooster endpoint
	 */
	Endpoint string

	/**
	 * timebooster api key
	 */
	ApiKey string

	Context *cli.Context

	/**
	 * コピー対象の環境変数一覧
	 */
	Environments []string
}

func NewInstance(ctx *cli.Context) *CommandRun {
	result := &CommandRun{
		Context:      ctx,
		Environments: []string{},
	}
	fmt.Printf("Check CLI Options...\n")

	{
		art := ctx.String("artifact")
		if len(art) > 0 {
			fmt.Printf(" -artifact %v\n", art)
			result.ArtifactPath = art
		}
	}

	for _, env := range ctx.StringSlice("env") {
		result.Environments = append(result.Environments, env)
	}
	if ctx.BoolT("env-from-circleci") {
		result.Environments = append(result.Environments,
			"CIRCLECI",
			"CI",
			"CIRCLE_PROJECT_USERNAME",
			"CIRCLE_PROJECT_REPONAME",
			"CIRCLE_BRANCH",
			"CIRCLE_TAG",
			"CIRCLE_SHA1",
			"CIRCLE_REPOSITORY_URL",
			"CIRCLE_COMPARE_URL",
			"CIRCLE_BUILD_URL",
			"CIRCLE_BUILD_NUM",
			"CIRCLE_PREVIOUS_BUILD_NUM",
			"CI_PULL_REQUESTS",
			"CI_PULL_REQUEST",
			"CIRCLE_USERNAME",
			"CIRCLE_PR_USERNAME",
			"CIRCLE_PR_REPONAME",
			"CIRCLE_PR_NUMBER",
			"CIRCLE_NODE_TOTAL",
			"CIRCLE_NODE_INDEX",
			"CIRCLE_BUILD_IMAGE",
			//"CIRCLE_ARTIFACTS",		// conflict!
			//"CIRCLE_TEST_REPORTS",	// conflict!
		)
	}

	result.ConfigPath = getConfigFilePath(ctx)
	result.ApiKey = getApiKey(ctx)
	result.Endpoint = getServerEndpoint(ctx)
	result.Repository = ctx.String("repository")
	result.Revision = ctx.String("revision")
	result.GithubAccessToken = ctx.String("github-access-token")

	// githubアクセストークンをURL変換する
	if len(result.GithubAccessToken) > 0 {
		result.Repository = utils.GetGithubRepositoryPath(result.Repository, result.GithubAccessToken)
	}
	if ctx.BoolT("print-repository") {
		fmt.Printf(" -repository %v %v\n", result.Repository, result.Revision)
	}

	fmt.Printf(" -api-key Hash(%v)\n", utils.ToMD5(result.Repository+result.ApiKey))
	fmt.Printf(" -endpoint Hash(%v)\n", utils.ToMD5(result.Repository+result.Endpoint))
	fmt.Printf(" -config %v\n", result.ConfigPath)
	if len(result.Environments) > 0 {
		fmt.Printf(" Environments [")
		for _, name := range result.Environments {
			fmt.Printf("%v, ", name)
		}
		fmt.Printf("]\n")
	}
	return result
}

/**
 * ビルドタスクを投げる
 */
func (it *CommandRun) requestNewBuildTask() (*api_v1.BuildInfo, error) {
	if len(it.Endpoint) == 0 {
		return nil, errors.New("timebooster -endpoint not set")
	}

	request := &api_v1.BuildApiBuildsPostRequest{
		Key:     swag.String(it.ApiKey),
		Payload: &api_v1.BuildRequest{},
	}

	// コンフィグをロードしてビルドリクエストを投げる
	if config, err := loadConfigFile(it.Context); err != nil {
		fmt.Printf("config load error[%v]", err.Error())
		return nil, err
	} else {
		request.Payload.Config = swag.String(config)

		if it.Context.BoolT("print-config") {
			fmt.Printf("Config file\n")
			fmt.Printf("============================================\n")
			fmt.Printf("%v\n", config)
			fmt.Printf("============================================\n")
		}

	}

	// リポジトリを転記する
	if len(it.Repository) > 0 || len(it.Revision) > 0 {
		request.Payload.Repository = &api_v1.BuildRepository{}
		if len(it.Repository) > 0 {
			request.Payload.Repository.Git = swag.String(it.Repository)
		}
		if len(it.Revision) > 0 {
			request.Payload.Repository.GitRevision = swag.String(it.Revision)
		}
	}

	// 環境変数をロードする
	{
		envArray := api_v1.EnvironmentValueArray{}

		for _, envName := range it.Environments {
			value := os.Getenv(envName)

			// debug print
			if it.Context.BoolT("print-env") {
				fmt.Printf("Environment link [%v] = [%v]\n", envName, value)
			}
			if len(value) > 0 {
				envArray = append(envArray, api_v1.EnvironmentValue{
					Name:  swag.String(envName),
					Value: swag.String(value),
				})
			}
		}

		if len(envArray) > 0 {
			request.Payload.Environment = &envArray
		}
	}

	client := swagger_utils.NewFetchClient(it.Endpoint, &http.Client{Timeout: (time.Duration(30) * time.Second)})
	api := api_v1.NewBuildApi()

	result := api_v1.BuildInfo{}
	if err := api.BuildsPost(client, request, &result); err != nil {
		fmt.Printf("Task create error[%v]", err.Error())
		return nil, err
	}

	fmt.Printf("New task ID[%v]\n", *result.Id)

	return &result, nil
}

func (it *CommandRun) awaitBuild(info *api_v1.BuildInfo) (*api_v1.BuildInfo, error) {
	var oldState api_v1.BuildState
	oldStateTime := time.Now()

	api := api_v1.NewBuildApi()

	// ビルドが完了するまで待つ
	for true {
		client := swagger_utils.NewFetchClient(it.Endpoint, &http.Client{Timeout: (time.Duration(30) * time.Second)})
		req := &api_v1.BuildApiBuildsBuildidGetRequest{
			Key:     swag.String(it.ApiKey),
			BuildId: info.Id,
		}
		if err := api.BuildsBuildidGet(client, req, info); err != nil {
			fmt.Printf("Task sync error[%v]\n", err.Error())
			return nil, err
		}

		if oldState != info.State.Value() {
			// ステートが変更された
			if len(oldState) > 0 {
				// かかった時間を表示
				now := time.Now()
				diff := now.Sub(oldStateTime)
				fmt.Printf(" %v seconds", int64(diff/time.Second))
				oldStateTime = now
			}

			if *info.State == api_v1.BuildState_Pending {
				fmt.Printf("\nTask [%v] SpinUp", *info.Id)
			} else {
				fmt.Printf("\nTask [%v] State Changed [%v]", *info.Id, *info.State)
			}
			oldState = *info.State

			// 状態チェック
			switch *info.State {
			case api_v1.BuildState_Timeout, api_v1.BuildState_Failed, api_v1.BuildState_Completed:
				{
					// ビルド状態が確定した
					fmt.Printf("\n")
					return info, nil
				}
			}
		} else {
			// ステート待ち
			fmt.Printf(".")
			time.Sleep(time.Duration(5) * time.Second)
		}
	}
	return nil, nil
}

func (it *CommandRun) downloadArtifact(info *api_v1.BuildInfo) error {
	os.MkdirAll(it.ArtifactPath, os.ModePerm)

	api := api_v1.NewBuildApi()
	request := &api_v1.BuildApiBuildsBuildidArtifactGetRequest{
		Key:     swag.String(getApiKey(it.Context)),
		BuildId: info.Id,
	}

	downloadPath := fmt.Sprintf("%v/%v.zip", it.ArtifactPath, *info.Id)

	client := swagger_utils.NewFetchClient(it.Endpoint, &http.Client{Timeout: (time.Duration(30) * time.Second)})
	client.CustomFetch = func(client *swagger_utils.BasicFetchClient, resultPtr interface{}) error {
		fmt.Printf("Download Artifacts...\n")

		resp, err := client.Client.Do(client.Request)
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			return nil
		}

		file, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			return err
		}

		bytes, err := io.Copy(file, resp.Body)
		fmt.Printf("Download [%v] %v bytes", downloadPath, bytes)
		return err
	}

	return api.BuildsBuildidArtifactGet(client, request, nil)
}

func (it *CommandRun) Execute() {
	task, err := it.requestNewBuildTask()
	if err != nil {
		os.Exit(1)
		return
	}

	task, err = it.awaitBuild(task)
	if err != nil {
		os.Exit(2)
		return
	}

	// アーティファクトダウンロードを行う
	if len(it.ArtifactPath) > 0 {
		// ダウンロードする
		err = it.downloadArtifact(task)
		if err != nil {
			os.Exit(3)
			return
		}
	}

	if *task.State != api_v1.BuildState_Completed {
		// ビルド失敗した
		os.Exit(4)
		return
	}

}

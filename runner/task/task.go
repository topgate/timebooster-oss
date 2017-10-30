package task

import (
	"../api_v1"
	"../utils"
	"context"
	"errors"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/**
 * timeboosterで実行されるタスクインターフェース
 */
type Task interface {
	/**
	 * 実行を行わせる
	 */
	Exec(ctx context.Context, env *TaskEnvironment) error
}

/**
 * タスク実行環境を定義する
 */
type TaskEnvironment struct {
	/**
	 * 制御ファイル本体
	 */
	Config Configure

	/**
	 * 元のビルド情報
	 */
	info *api_v1.BuildInfo

	/**
	 * 標準ログ
	 */
	stdLogText string

	/**
	 * 標準エラー
	 */
	errorLogText string
}

func (it *TaskEnvironment) GetBuildId() string {
	return *it.info.Id
}

/**
 * ワークスペースパスを取得する
 */
func (it *TaskEnvironment) GetWorkspacePath() string {
	result := fmt.Sprintf("private/%v/work", it.GetBuildId())
	os.MkdirAll(result, os.ModePerm)
	return result
}

/**
 * gitリポジトリ名を取得する
 */
func (it *TaskEnvironment) GetRepositoryName() string {
	repo := it.Config.Env.Repository

	if it.info.Repository != nil && it.info.Repository.Git != nil {
		repo = swag.StringValue(it.info.Repository.Git)
	}

	_, repoName := utils.SplitGithubRepository(repo)
	return repoName
}

/**
 * Workspace配下のリポジトリパスを取得する
 * これは ${GetWorkspacePath()}/${リポジトリ名} に一致する
 */
func (it *TaskEnvironment) GetRepositoryPath() string {
	return it.GetWorkspacePath() + "/" + it.GetRepositoryName()
}

/**
 * 成果物回収パスを取得する
 */
func (it *TaskEnvironment) GetArtifactPath() string {
	result := fmt.Sprintf("private/%v/art", it.GetBuildId())
	os.MkdirAll(result, os.ModePerm)
	return result
}

func (it *TaskEnvironment) NewExternalCommand(cmd string, args ...string) *utils.ExternalCommand {
	result := &utils.ExternalCommand{}

	result.Commands = []string{cmd}
	for _, v := range args {
		result.Commands = append(result.Commands, v)
	}
	result.Chdir = it.GetWorkspacePath()

	result.Stdout = func(line string) {
		it.Stdout("%v\n", line)
	}
	result.Stderr = func(line string) {
		it.Stderror("%v\n", line)
	}

	return result
}

/**
 * リポジトリを所定箇所にcloneする
 */
func (it *TaskEnvironment) Clone() error {
	repository := it.Config.Env.Repository
	revision := it.Config.Env.Revision

	if it.info.Repository != nil {
		// gitリポジトリの指定がある
		if it.info.Repository.Git != nil {
			repository = swag.StringValue(it.info.Repository.Git)
		}

		// リビジョン指定がある
		if it.info.Repository.GitRevision != nil {
			revision = swag.StringValue(it.info.Repository.GitRevision)
		}
	}

	// clone repo
	{
		cmd := it.NewExternalCommand("git", "clone", repository)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if len(revision) > 0 {
		it.Stdout("Checkout Revision[%v]\n", revision)
		cmd := it.NewExternalCommand("git", "checkout", "-f", revision)
		cmd.Chdir = it.GetRepositoryPath()
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

/**
 * Dockerfileをビルドする
 */
func (it *TaskEnvironment) buildDockerFile(dockerfilePath string) error {
	dockerBuildMutex.Lock()
	defer dockerBuildMutex.Unlock()

	//it.Stdout("repository | %v\n", it.GetRepositoryPath())
	it.Stdout("dockerfile | %v\n", dockerfilePath)

	// コンテナをビルドする
	cmd := it.NewExternalCommand("docker", "build", "-t", it.GetDockerImageTag(dockerfilePath), "-f", dockerfilePath, ".")
	cmd.Chdir = it.GetRepositoryPath()

	return cmd.Run()
}

func (it *TaskEnvironment) pullDockerImage(dockerImage string) error {
	dockerBuildMutex.Lock()
	defer dockerBuildMutex.Unlock()

	it.Stdout("pull docker image | %v\n", dockerImage)

	// imageをpullさせる
	cmd := it.NewExternalCommand("gcloud", "docker", "--", "pull", dockerImage)
	cmd.Chdir = it.GetRepositoryPath()

	return cmd.Run()
}

/**
 * 相対パスからDocker実行用の絶対パスに変換する
 */
func toDockerPath(path string) string {
	fullPath, _ := filepath.Abs(path)

	// windows ?
	if strings.Contains(fullPath, ":") {
		fullPath = strings.Replace(fullPath, ":\\", "/", -1)
		fullPath = strings.Replace(fullPath, "\\", "/", -1)
		fullPath = "/" + fullPath
	}

	return fullPath
}

/**
 * Docker imageに指定されるタグを取得する
 */
func (it *TaskEnvironment) GetDockerImageTag(dockerfilePath string) string {
	buf, _ := ioutil.ReadFile(it.GetRepositoryPath() + "/" + dockerfilePath)
	return utils.ToMD5(string(buf))
}

/**
 * dockerコンテナ内で処理を実行する
 */
func (it *TaskEnvironment) runCommands(taskIndex int, dockerfilePath string, dockerImagePath string, cmdList []string) error {
	// 実行ファイルを生成
	os.MkdirAll(it.GetWorkspacePath()+"/runner", os.ModePerm)
	{
		shell := "#! /bin/bash -eu\n\n"
		for _, line := range cmdList {
			shell += line + "\n"
		}
		ioutil.WriteFile(it.GetWorkspacePath()+"/runner"+"/.timebooster-internal.sh", []byte(shell), os.ModeExclusive)
	}
	{
		shell := "#! /bin/bash -eu\n\n"
		shell += "chmod 755 /runner/.timebooster-internal.sh\n"
		shell += "/runner/.timebooster-internal.sh &> /artifacts/"
		shell += fmt.Sprintf("exec-%v.txt\n", taskIndex)
		ioutil.WriteFile(it.GetWorkspacePath()+"/runner"+"/.timebooster.sh", []byte(shell), os.ModeExclusive)
	}

	// 実行する
	cmd := it.NewExternalCommand("docker", "run", "--rm")

	// 環境変数設定
	cmd.Commands = append(cmd.Commands, "-e", "TIMEBOOSTER_ARTIFACTS=/artifacts")  // 環境変数追加
	cmd.Commands = append(cmd.Commands, "-e", "TIMEBOOSTER_BUILD_ID="+*it.info.Id) // ビルドID
	for key, value := range it.Config.Env.Variable {
		arg := fmt.Sprintf("%v", value)
		if len(key) > 0 && len(arg) > 0 {
			cmd.Commands = append(cmd.Commands, "-e", key+"="+arg) // 環境変数追加
		}
	}
	// サーバーからの環境変数引き継ぎ
	if it.info.Environment != nil {
		for _, env := range *it.info.Environment {
			cmd.Commands = append(cmd.Commands, "-e", *env.Name+"="+*env.Value) // サーバーから環境変数引き継ぎ
		}
	}

	// ボリュームマウント
	cmd.Commands = append(cmd.Commands, "-v", toDockerPath(it.GetArtifactPath())+":/artifacts")         // 成果物回収ディレクトリ
	cmd.Commands = append(cmd.Commands, "-v", toDockerPath(it.GetWorkspacePath())+":/work")             // ビルドワークスペース
	cmd.Commands = append(cmd.Commands, "-v", toDockerPath(it.GetWorkspacePath()+"/runner")+":/runner") // タスク実行シェルパス
	cmd.Commands = append(cmd.Commands, "-w", "/work/"+it.GetRepositoryName())                          // 初期カレントディレクトリ

	// キャッシュディレクトリをマウント
	for _, cache := range it.Config.Env.Cache {
		if len(cache) > 0 {
			cmd.Commands = append(cmd.Commands, "-v", toDockerPath("private/mnt"+cache)+":"+cache)
		}
	}

	if len(dockerfilePath) > 0 {
		cmd.Commands = append(cmd.Commands, it.GetDockerImageTag(dockerfilePath)) // 実行コンテナ
	} else {
		cmd.Commands = append(cmd.Commands, dockerImagePath) // 実行コンテナ
	}
	cmd.Commands = append(cmd.Commands, "/bin/bash", "-c", "chmod 755 /runner/.timebooster.sh; /runner/.timebooster.sh") // 実行コマンド

	execText := "Run... [ "
	for _, v := range cmd.Commands {
		execText += (v + " ")
	}
	execText += "]\n"
	it.Stdout(execText)

	return cmd.Run()
}

/**
 * 成果物を改修する
 */
func (it *TaskEnvironment) collectArtifacts() {
	// ログを出力する
	if len(it.stdLogText) > 0 {
		ioutil.WriteFile(it.GetArtifactPath()+"/timebooster-stdout.txt", []byte(it.stdLogText), os.ModePerm)
	}
	if len(it.errorLogText) > 0 {
		ioutil.WriteFile(it.GetArtifactPath()+"/timebooster-stderror.txt", []byte(it.errorLogText), os.ModePerm)
	}

	// zipする
	ZIP_FILE := "artifacts.zip"
	{
		cmd := it.NewExternalCommand("zip", "-r", "work/"+ZIP_FILE, "art/")
		cmd.Chdir = it.GetWorkspacePath() + "/../"
		cmd.Run()
	}
	// アップロード
	{
		cmd := it.NewExternalCommand("gsutil", "cp", it.GetWorkspacePath()+"/"+ZIP_FILE, "gs://"+utils.GetTimeboosterProjectId()+".appspot.com/artifacts/"+it.GetBuildId()+"/")
		cmd.Chdir = ""
		cmd.Run()
	}
}

/**
 * タスクを実行する
 */
func (it *TaskEnvironment) Execute() error {
	// 成果物は必ず回収する
	defer it.collectArtifacts()

	for index, exec := range it.Config.Task.Exec {
		// docker imageを生成
		if len(exec.Dockerfile) > 0 {
			if err := it.buildDockerFile(exec.Dockerfile); err != nil {
				fmt.Printf("docker build error %v", err.Error())
				return err
			}
		} else if len(exec.DockerImage) > 0 {
			if err := it.pullDockerImage(exec.DockerImage); err != nil {
				fmt.Printf("docker pull error %v", err.Error())
				return err
			}
		} else {
			return errors.New("docker file/image not found.")
		}

		// コマンドを実行
		if err := it.runCommands(index, exec.Dockerfile, exec.DockerImage, exec.Cmd); err != nil {
			return err
		}
	}

	return nil
}

/**
 * 標準出力を行う
 *
 * 標準出力の内容はGAE/Goのサーバーに送られ、Circle CIからポーリングされる
 */
func (it *TaskEnvironment) Stdout(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	if len(args) > 0 {
		it.stdLogText += fmt.Sprintf(format, args)
	} else {
		it.stdLogText += format
	}
}

/**
 * 標準出力を行う
 *
 * 標準出力の内容はGAE/Goのサーバーに送られ、Circle CIからポーリングされる
 */
func (it *TaskEnvironment) Stderror(format string, args ...interface{}) {
	fmt.Errorf(format, args...)
	if len(args) > 0 {
		it.errorLogText += fmt.Sprintf(format, args)
	} else {
		it.errorLogText += format
	}
}

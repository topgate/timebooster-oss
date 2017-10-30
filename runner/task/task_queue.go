package task

import (
	"../api_v1"
	"../utils"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"github.com/go-yaml/yaml"
	"net/http"
	"time"
)

type TaskQueue struct {
	/**
	 * 実行されるタスク一覧
	 */
	Tasks api_v1.BuildInfoArray
}

/**
 * ステートを切り替える
 */
func (it *TaskQueue) patchTask(task *api_v1.BuildInfo) error {
	client := swagger_utils.NewFetchClient(utils.GetServerEndpoint(), &http.Client{Timeout: (time.Duration(30) * time.Second)})
	api := api_v1.NewBuildApi()

	req := &api_v1.BuildApiBuildsBuildidPatchRequest{
		Key:       swag.String(utils.GetApiKey()),
		BuildId:   task.Id,
		NewObject: task,
	}
	return api.BuildsBuildidPatch(client, req, task)
}

/**
 * タスクを実行する
 */
func (it *TaskQueue) runTask(task *api_v1.BuildInfo) {
	fmt.Printf("Task Start[%v]\n", task.Id)
	defer func() {
		fmt.Printf("Task Finish[%v] State[%v]\n", *task.Id, *task.State)
		// 最終結果に書き換える
		it.patchTask(task)
	}()

	task.State = api_v1.BuildState_Building.Ptr()
	if err := it.patchTask(task); err != nil {
		task.State = api_v1.BuildState_Failed.Ptr()
		return
	}

	env := &TaskEnvironment{
		info: task,
	}

	if err := yaml.Unmarshal([]byte(*task.Config), &env.Config); err != nil {
		fmt.Printf("Yaml error %v\n", err.Error())
		task.State = api_v1.BuildState_Failed.Ptr()
		return
	}

	env.Stdout("Build ID[%v]\n", env.GetBuildId())
	//env.Stdout("Build Repo[%v]\n", strings.Trim(env.Config.Env.Repository, "\n"))

	// リポジトリをCloneする
	if err := env.Clone(); err != nil {
		fmt.Printf("Clone error %v\n", err.Error())
		task.State = api_v1.BuildState_Failed.Ptr()
		return
	}

	// Dockerで指定ビルドを行う
	if err := env.Execute(); err != nil {
		fmt.Printf("Execute error %v\n", err.Error())
		task.State = api_v1.BuildState_Failed.Ptr()
		return

	}

	// 全タスク成功
	task.State = api_v1.BuildState_Completed.Ptr()
}

func (it *TaskQueue) run(context *MachineContext, task *api_v1.BuildInfo) {
	// ビルド中にタスク状態を変更する
	task.State = api_v1.BuildState_Building.Ptr()
	if err := it.patchTask(task); err != nil {
		fmt.Printf("BuildTask start failed[%v] error[%v]\n", *task.Id, err.Error())
	} else {
		fmt.Printf("Go BuildTask start [%v]\n", *task.Id)
		// 以後は並列実行させる
		go func(task *api_v1.BuildInfo) {
			defer context.TaskDone()

			// 並列実行
			it.runTask(task)
		}(task)
	}
}

/**
 * 指定した並列数で実行させる
 */
func (it *TaskQueue) Run(context *MachineContext) {
	// 全てのタスクを並列実行
	for _, item := range it.Tasks {
		context.TaskAdd()
		it.run(context, &item)
	}
}

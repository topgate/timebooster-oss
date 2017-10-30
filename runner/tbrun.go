package main

import (
	"./api_v1"
	"./task"
	"./utils"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"time"
)

/**
 * 実行対象のビルドタスクを全て取得する
 */
func fetchBuildTasks() api_v1.BuildInfoArray {
	client := swagger_utils.NewFetchClient(utils.GetServerEndpoint(), &http.Client{Timeout: (time.Duration(30) * time.Second)})
	api := api_v1.NewBuildApi()

	req := &api_v1.BuildApiBuildsGetRequest{
		State: swag.String(string(api_v1.BuildState_Pending)),
		Key:   swag.String(utils.GetApiKey()),
	}
	builds := api_v1.BuildInfoArray{}

	if err := api.BuildsGet(client, req, &builds); err != nil {
		fmt.Printf("FetchError[%v]", err.Error())
	}

	return builds
}

func main() {

	context := task.NewContext()

	for !context.IsShutdownTime() {
		builds := fetchBuildTasks()
		fmt.Printf("Polling tasks[%v]\n", len(builds))

		if len(builds) > 0 {
			fmt.Printf("Run %v tasks\n", len(builds))
			// Run Tasks
			// デフォルトで4並列実行としておく
			queue := &task.TaskQueue{
				Tasks: builds,
			}
			queue.Run(context)
		}

		// 次のポーリングまで待機
		time.Sleep(time.Second * 5)
	}

	fmt.Printf("Shutdown Time[%v]\n", context.ShutdownTime)
}

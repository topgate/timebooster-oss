package api_build

import (
	"api_v1"
	"data"
	"datastore/model"
	"datastore/service"
	"fmt"
	"github.com/eaglesakura/swagger-go-core"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"utils"
)

/**
 * 指定条件のビルド一覧を取得する
 */
func BuildsGet(req swagger.RequestContext, params *api_v1.BuildsGetParams) swagger.Responder {
	ctx, _ := req.(data.Context)
	buildService := service.NewBuildService(ctx)

	var temp model.BuildInfoArray
	if params.State != nil {
		// ステート条件で検索する
		temp = buildService.ListBuildTasks(api_v1.BuildState(*params.State))
	}

	ctx.LogDebug("Found builds[%v]", len(temp))

	result := temp.ToApiModel()
	return &result
}

/**
 * 新規ビルドを開始する
 */
func BuildsPost(req swagger.RequestContext, params *api_v1.BuildsPostParams) swagger.Responder {
	ctx, _ := req.(data.Context)
	buildService := service.NewBuildService(ctx)

	task, err := buildService.NewBuildTask(params.Payload)
	if err != nil {
		ctx.LogError("Task create failed[%v]", err.Error())
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_ParameterError, "")
	}

	ctx.LogInfo("NewBuild task[%v]", task.Id)

	// マシンを起動させる
	machineService := service.NewMachineService(ctx)
	if err := machineService.Start(); err != nil {
		ctx.LogError("Build machine start failed[%v]", err.Error())
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_DataModifyFailed, "")
	}

	result := task.ToApiModel()
	return &result
}

/**
 * 指定ビルドの状態を取得する
 */
func BuildsBuildIdGet(req swagger.RequestContext, params *api_v1.BuildsBuildidGetParams) swagger.Responder {
	ctx, _ := req.(data.Context)
	buildService := service.NewBuildService(ctx)

	task := buildService.GetBuildTask(*params.BuildId)
	if task == nil {
		return &swagger_utils.RawBufferResponse{
			StatusCode: http.StatusNotFound,
		}
	}

	result := task.ToApiModel()
	return &result
}

/**
 * 指定ビルドの状態を更新する
 */
func BuildsBuildIdPatch(req swagger.RequestContext, params *api_v1.BuildsBuildidPatchParams) swagger.Responder {
	ctx, _ := req.(data.Context)
	buildService := service.NewBuildService(ctx)

	task, err := buildService.PatchBuildTask(*params.BuildId, params.NewObject.State)
	if err != nil {
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_DataModifyFailed, fmt.Sprintf("Build modify error[%v]", *params.BuildId))
	}

	result := task.ToApiModel()
	return &result
}

/**
 * 指定ビルドの状態を更新する
 */
func BuildsBuildIdArtifactGet(req swagger.RequestContext, params *api_v1.BuildsBuildidArtifactGetParams) swagger.Responder {
	ctx, _ := req.(data.Context)
	buildService := service.NewBuildService(ctx)

	link := buildService.GetArtifactsLink(*params.BuildId)

	if len(link) == 0 {
		//return &swagger_utils.RawBufferResponse{
		//	StatusCode:http.StatusNotFound,
		//}

		return &swagger_utils.RedirectResponse{
			Location: "https://google.com",
		}
	} else {
		return &swagger_utils.RedirectResponse{
			Location: link,
		}
	}
}

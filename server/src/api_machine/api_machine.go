package api_machine

import (
	"api_v1"
	"data"
	"datastore/service"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/swag"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"utils"
)

/**
 * マシン状態を取得する
 */
func MachineGet(req swagger.RequestContext, params *api_v1.MachineGetParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)

	info := machineService.LoadMachineInfo()

	result := &api_v1.MachineInfo{}
	result.Id = swag.String(info.Id)
	result.State = machineService.GetMachineStatus().Ptr()
	result.Zone = swag.String(info.Zone)

	return result
}

/**
 * ビルドマシンを生成する
 */
func MachinePost(req swagger.RequestContext, params *api_v1.MachinePostParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)
	if info, err := machineService.Create(params.Payload); err != nil {
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_Unknown, err.Error())
	} else {
		return &api_v1.MachineInfo{
			Id:    swag.String(info.Id),
			State: api_v1.MachineState_Shutdown.Ptr(),
			Zone:  swag.String(info.Zone),
		}
	}
}

/**
 * ビルドマシンを削除する
 */
func MachineDelete(req swagger.RequestContext, params *api_v1.MachineDeleteParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)
	err := machineService.Delete()
	if err != nil {
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_Unknown, err.Error())
	} else {
		return &swagger_utils.RawBufferResponse{
			StatusCode: http.StatusOK,
		}
	}
}

/**
 * ビルドマシンを起動する
 */
func MachineBoot(req swagger.RequestContext, params *api_v1.MachineBootPostParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)

	if err := machineService.Start(); err != nil {
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_Unknown, err.Error())
	}

	// 起動成功
	return &swagger_utils.RawBufferResponse{
		StatusCode: http.StatusOK,
	}
}

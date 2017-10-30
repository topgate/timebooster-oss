package api_machine

import (
	"api_v1"
	"data"
	"datastore/service"
	"github.com/eaglesakura/swagger-go-core"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"utils"
)

/**
 * startup-scriptを取得する
 */
func MachineStartscriptGet(req swagger.RequestContext, params *api_v1.MachineStartupscriptGetParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)
	machine := machineService.LoadMachineInfo()

	return &swagger_utils.RawBufferResponse{
		StatusCode:  200,
		ContentType: "text/plain",
		Payload:     []byte(machine.StartupScript),
	}
}

/**
 * startup-scriptを更新する
 */
func MachineStartscriptPost(req swagger.RequestContext, params *api_v1.MachineStartupscriptPostParams) swagger.Responder {
	ctx, _ := req.(data.Context)

	machineService := service.NewMachineService(ctx)
	if err := machineService.SetStartupscript(*params.Script); err != nil {
		ctx.LogError("Failed startup script[%v]", err.Error())
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_DataModifyFailed, "Update failed")
	}

	return &swagger_utils.RawBufferResponse{
		StatusCode: 200,
	}
}

package appserver

import (
	"api_build"
	"api_machine"
	"api_v1"
	"data"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"server"
)

func init() {

	server := server.NewApplication()
	mapper := utils.NewHandleMapper()

	// debug api
	{
		api := api_v1.NewDebugApiController()
		api.HandleDebugSwaggerGet(DebugSwagger)
		api.MapHandlers(mapper)
	}

	// build api
	{
		api := api_v1.NewBuildApiController()
		api.HandleBuildsGet(api_build.BuildsGet)
		api.HandleBuildsPost(api_build.BuildsPost)
		api.HandleBuildsBuildidGet(api_build.BuildsBuildIdGet)
		api.HandleBuildsBuildidPatch(api_build.BuildsBuildIdPatch)
		api.HandleBuildsBuildidArtifactGet(api_build.BuildsBuildIdArtifactGet)
		api.MapHandlers(mapper)
	}
	// Machine api
	{
		api := api_v1.NewMachineApiController()
		api.HandleMachineGet(api_machine.MachineGet)
		api.HandleMachinePost(api_machine.MachinePost)
		api.HandleMachineDelete(api_machine.MachineDelete)

		// マシン設定変更
		api.HandleMachineStartupscriptGet(api_machine.MachineStartscriptGet)
		api.HandleMachineStartupscriptPost(api_machine.MachineStartscriptPost)

		// マシン起動
		api.HandleMachineBootPost(api_machine.MachineBoot)
		api.MapHandlers(mapper)
	}

	http.Handle("/api/v1/", mapper.NewRouter(server))
}

func DebugSwagger(_ctx swagger.RequestContext, param *api_v1.DebugSwaggerGetParams) swagger.Responder {

	ctx := _ctx.(data.Context)
	swagger, _ := ctx.GetAssets().LoadFile("assets/swagger.json")

	return &utils.RawBufferResponse{
		StatusCode:  200,
		ContentType: "application/json",
		Payload:     swagger,
	}
}

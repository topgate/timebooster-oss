package service

import (
	"api_v1"
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/stretchr/testify/assert"
	"server"
	"testing"
)

const testApiKey string = "AIzaSyAEj97VicJ8OQqMNzYiWk1HQhOr21PQl_Y"

func TestMachineService_GetMachineId(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	service := NewMachineService(ctx)
	assert.NotEqual(t, service.GetMachineId(), "")
}

func TestMachineService_GetMachineState_NotCreated(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey + "this,is,error,key"),
	}
	defer ctx.Context.Close()

	service := NewMachineService(ctx)
	assert.Equal(t, service.GetMachineStatus(), api_v1.MachineState_None)
}

//func TestMachineService_Start(t *testing.T) {
//	t.Parallel()
//	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
//	ctx.Options.Auth = &gaefire.AuthenticationInfo{
//		ApiKey: swag.String(testApiKey),
//	}
//	defer ctx.Context.Close()
//
//	service := NewMachineService(ctx)
//	assert.Nil(t, service.Start())
//}

//func TestMachineService_Delete(t *testing.T) {
//	t.Parallel()
//	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
//	ctx.Options.Auth = &gaefire.AuthenticationInfo{
//		ApiKey: swag.String(testApiKey),
//	}
//	defer ctx.Context.Close()
//
//	service := NewMachineService(ctx)
//	assert.Nil(t, service.Delete())
//}

//func TestMachineService_Create(t *testing.T) {
//	t.Parallel()
//	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
//	ctx.Options.Auth = &gaefire.AuthenticationInfo{
//		ApiKey: swag.String(testApiKey),
//	}
//	defer ctx.Context.Close()
//
//	service := NewMachineService(ctx)
//	req := &api_v1.MachineRequest{
//		Zone:swag.String("us-east1-d"),
//		Cpu:swag.Int32(4),
//		Ram:swag.Float32(16.0),
//		Storage:swag.Float32(24.0),
//	}
//	machine, err := service.Create(req)
//	assert.Nil(t, err)
//	assert.NotNil(t, machine)
//	assert.Equal(t, machine.Zone, "us-east1-d")
//}

func TestMachineService_SetBootscript(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	{
		service := NewMachineService(ctx)
		info := service.LoadMachineInfo()
		assert.NotEqual(t, len(info.StartupScript), 0)
		service.SetStartupscript(`#this,is,test,script`)
	}

	{
		service := NewMachineService(ctx)
		info := service.LoadMachineInfo()
		assert.Equal(t, info.StartupScript, `#this,is,test,script`)
	}
}

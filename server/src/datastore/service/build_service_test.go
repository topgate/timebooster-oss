package service

import (
	"api_v1"
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/stretchr/testify/assert"
	"server"
	"testing"
)

func TestBuildService_NewBuildTask(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	service := NewBuildService(ctx)

	var buildId string
	// 新規にタスク生成できる
	{
		req := &api_v1.BuildRequest{
			Config: swag.String("#this,is,test"),
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		buildId = task.Id
	}

	// タスクがロードできる
	{
		task := service.GetBuildTask(buildId)
		assert.NotNil(t, task)
		assert.Equal(t, task.Id, buildId)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		// タスクを変換できる
		model := task.ToApiModel()
		assert.Equal(t, task.Id, *model.Id)
		assert.Nil(t, model.Environment)
		assert.Nil(t, model.Repository)
	}

}

func TestBuildService_NewBuildTaskWithRepository(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	service := NewBuildService(ctx)

	var buildId string
	// 新規にタスク生成できる
	{
		req := &api_v1.BuildRequest{
			Config: swag.String("#this,is,test"),
			Repository: &api_v1.BuildRepository{
				Git: swag.String("git@github.com:user/repo.git"),
			},
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		buildId = task.Id
	}

	// タスクがロードできる
	{
		task := service.GetBuildTask(buildId)
		assert.NotNil(t, task)
		assert.Equal(t, task.Id, buildId)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		// タスクを変換できる
		model := task.ToApiModel()
		assert.Equal(t, task.Id, *model.Id)
		assert.Equal(t, swag.StringValue(model.Repository.Git), "git@github.com:user/repo.git")
		assert.Nil(t, model.Repository.GitRevision)
	}
}

func TestBuildService_NewBuildTaskWithRepositoryRevision(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	service := NewBuildService(ctx)

	var buildId string
	// 新規にタスク生成できる
	{
		req := &api_v1.BuildRequest{
			Config: swag.String("#this,is,test"),
			Repository: &api_v1.BuildRepository{
				Git:         swag.String("git@github.com:user/repo.git"),
				GitRevision: swag.String("develop"),
			},
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		buildId = task.Id
	}

	// タスクがロードできる
	{
		task := service.GetBuildTask(buildId)
		assert.NotNil(t, task)
		assert.Equal(t, task.Id, buildId)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		// タスクを変換できる
		model := task.ToApiModel()
		assert.Equal(t, task.Id, *model.Id)
		assert.Equal(t, swag.StringValue(model.Repository.Git), "git@github.com:user/repo.git")
		assert.Equal(t, swag.StringValue(model.Repository.GitRevision), "develop")
	}
}

func TestBuildService_NewBuildTaskWithEnv(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	service := NewBuildService(ctx)

	var buildId string
	// 環境変数付きのタスクを登録できる
	{
		env := &api_v1.EnvironmentValueArray{
			api_v1.EnvironmentValue{
				Name:  swag.String("Key1"),
				Value: swag.String("Value1"),
			},
			api_v1.EnvironmentValue{
				Name:  swag.String("Key2"),
				Value: swag.String("Value2"),
			},
		}
		req := &api_v1.BuildRequest{
			Config:      swag.String("#this,is,test"),
			Environment: env,
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		buildId = task.Id
	}

	// タスクがロードできる
	{
		task := service.GetBuildTask(buildId)
		assert.NotNil(t, task)
		assert.Equal(t, task.Id, buildId)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)
		assert.Equal(t, len(task.Environments), 2)
		assert.Equal(t, task.Environments[0].Name, "Key1")
		assert.Equal(t, task.Environments[0].Value, "Value1")
		assert.Equal(t, task.Environments[1].Name, "Key2")
		assert.Equal(t, task.Environments[1].Value, "Value2")

		// タスクを変換できる
		model := task.ToApiModel()
		assert.Equal(t, len(*model.Environment), 2)
		assert.Equal(t, *(*model.Environment)[0].Name, "Key1")
		assert.Equal(t, *(*model.Environment)[0].Value, "Value1")
		assert.Equal(t, *(*model.Environment)[1].Name, "Key2")
		assert.Equal(t, *(*model.Environment)[1].Value, "Value2")
	}

}

func TestBuildService_PatchBuildTask(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	var buildId string
	// 新規にタスク生成できる
	{
		service := NewBuildService(ctx)
		req := &api_v1.BuildRequest{
			Config: swag.String("#this,is,test"),
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)

		buildId = task.Id
	}

	// ステートを変更する
	{
		service := NewBuildService(ctx)
		task, err := service.PatchBuildTask(buildId, api_v1.BuildState_Building.Ptr())
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Building)
	}

	// タスクがロードできる
	{
		service := NewBuildService(ctx)
		task := service.GetBuildTask(buildId)
		assert.NotNil(t, task)
		assert.Equal(t, task.Id, buildId)
		assert.Equal(t, task.State, api_v1.BuildState_Building)
	}
}

func TestBuildService_ListBuildTask(t *testing.T) {
	t.Parallel()
	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
	ctx.Options.Auth = &gaefire.AuthenticationInfo{
		ApiKey: swag.String(testApiKey),
	}
	defer ctx.Context.Close()

	// 新規にタスク生成できる
	{
		service := NewBuildService(ctx)
		req := &api_v1.BuildRequest{
			Config: swag.String("#this,is,test"),
		}

		task, err := service.NewBuildTask(req)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, task.State, api_v1.BuildState_Pending)
	}

	// タスクを列挙する
	{
		service := NewBuildService(ctx)
		tasks := service.ListBuildTasks(api_v1.BuildState_Pending)
		assert.Equal(t, len(tasks), 1)
		assert.Equal(t, tasks[0].State, api_v1.BuildState_Pending)
	}
}

//func TestBuildService_GetArtifactsLink(t *testing.T) {
//	t.Parallel()
//	ctx, _ := server.NewApplication().NewContext(nil).(*server.ContextImpl)
//	ctx.Options.Auth = &gaefire.AuthenticationInfo{
//		ApiKey: swag.String(testApiKey),
//	}
//	defer ctx.Context.Close()
//
//	service := NewBuildService(ctx)
//	link := service.GetArtifactsLink("d02d87bf2097de8afca9f25600c60667-ccedafd04c35ac2273b7a98c1a2922c479e8b286")
//	ctx.LogDebug("Download Link[%v]", link)
//
//	assert.NotEqual(t, len(link), 0)
//	assert.Equal(t, strings.Index(link, "http"), 0)
//}

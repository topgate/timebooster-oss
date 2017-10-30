package service

import (
	"api_v1"
	"cloud.google.com/go/storage"
	"data"
	"datastore/model"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/file"
	"time"
	"utils"
)

/**
 * ビルドマシン管理用サービス
 */
type BuildService struct {
	req data.Context
	ctx context.Context
}

func NewBuildService(req data.Context) *BuildService {
	result := &BuildService{
		req: req,
		ctx: req.GetOptions().Context,
	}
	return result
}

/**
 * ビルドタスクを新規に作成する
 */
func (it *BuildService) NewBuildTask(req *api_v1.BuildRequest) (*model.BuildInfo, error) {
	machineId := getMachineId(it.req)
	now := time.Now()

	build := &model.BuildInfo{
		Id:           machineId + "-" + utils.ToSHA1(fmt.Sprintf("%v,%v", time.Now(), appengine.RequestID(it.ctx))),
		MachineId:    machineId,
		State:        api_v1.BuildState_Pending,
		CreatedDate:  now,
		ModifiedDate: now,
		Config:       *req.Config,
	}

	// repo
	if req.Repository != nil {
		build.Repository.Git = swag.StringValue(req.Repository.Git)
		build.Repository.GitRevision = swag.StringValue(req.Repository.GitRevision)
	}

	// add env
	if req.Environment != nil && len(*req.Environment) > 0 {
		build.Environments = []model.BuildEnvironment{}
		for _, item := range *req.Environment {
			build.Environments = append(build.Environments, model.BuildEnvironment{Name: *item.Name, Value: *item.Value})
		}
	}

	g := it.req.GetGoon()
	_, err := g.Put(build)
	if err != nil {
		return nil, err
	} else {
		return build, nil
	}
}

/**
 * ビルド情報一覧を取得する
 * 指定したAPIkeyに紐付いたビルドタスクのみが取得される
 */
func (it *BuildService) ListBuildTasks(state api_v1.BuildState) []model.BuildInfo {
	g := it.req.GetGoon()

	result := []model.BuildInfo{}
	g.GetAll(datastore.NewQuery("BuildInfo").Filter("State =", string(state)).Filter("MachineId =", getMachineId(it.req)), &result)

	return result
}

/**
 * 指定したビルドタスクの現在状態を取得する
 *
 * タスクが見つからない場合はnilを返却する
 */
func (it *BuildService) GetBuildTask(buildId string) *model.BuildInfo {
	g := it.req.GetGoon()
	model := &model.BuildInfo{
		Id: buildId,
	}

	if err := g.Get(model); err != nil {
		it.req.LogError("Task load error[%v]", err.Error())
		return nil
	} else {
		return model
	}
}

/**
 * アーティファクトダウンロードリンクを取得する
 */
func (it *BuildService) GetArtifactsLink(buildId string) string {
	bucket, _ := file.DefaultBucketName(it.ctx)
	path := fmt.Sprintf("artifacts/%v/artifacts.zip", buildId)
	expires := time.Now().Add(time.Second * 120)
	email, _ := appengine.ServiceAccount(it.ctx)

	url, err := storage.SignedURL(bucket, path, &storage.SignedURLOptions{
		GoogleAccessID: email,
		SignBytes: func(b []byte) ([]byte, error) {
			_, signedBytes, err := appengine.SignBytes(it.ctx, b)
			return signedBytes, err
		},
		Method:  "GET",
		Expires: expires,
	})
	if err != nil {
		it.req.LogError("SignedError[%v] email[%v]", err.Error(), email)
		return ""
	}

	return url
}

/**
 * 指定したビルドタスクの現在状態を取得する
 *
 * タスクが見つからない場合はnil,errorを返却する
 */
func (it *BuildService) PatchBuildTask(buildId string, buildState *api_v1.BuildState) (*model.BuildInfo, error) {
	g := it.req.GetGoon()
	model := &model.BuildInfo{
		Id: buildId,
	}

	err := g.RunInTransaction(func(g *goon.Goon) error {
		if err := g.Get(model); err != nil {
			return err
		}

		if buildState != nil {
			model.State = *buildState
		}

		// 更新日時を変更
		model.ModifiedDate = time.Now()
		g.Put(model)
		return nil
	}, nil)

	if err != nil {
		return nil, err
	} else {
		return model, nil
	}
}

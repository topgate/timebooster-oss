package model

import (
	"api_v1"
	"github.com/eaglesakura/swagger-go-core/swag"
	"google.golang.org/appengine/datastore"
	"time"
	"utils"
)

/**
 * ビルドのリポジトリ設定
 */
type BuildRepository struct {
	/**
	 * gitリポジトリ
	 */
	Git string `datastore:",noindex"`

	/**
	 * gitリビジョン
	 */
	GitRevision string `datastore:",noindex"`
}

/**
 * ビルドに付与する環境変数
 */
type BuildEnvironment struct {
	Name  string `datastore:",noindex"`
	Value string `datastore:",noindex"`
}

/**
 * 1ビルドごとに1生成されるEntity
 * ビルド成否等を保存する
 */
type BuildInfo struct {
	/**
	 * ランダムに割り当てられるビルドID
	 */
	Id string `datastore:"-" goon:"id"`

	/**
	 * 実行されたマシンID
	 */
	MachineId string

	/**
	 * 現在のビルド状態
	 */
	State api_v1.BuildState

	/**
	 * ビルド情報のYAML
	 */
	Config string `datastore:",noindex"`

	/**
	 * リポジトリ設定
	 */
	Repository BuildRepository

	/**
	 * 環境変数一覧
	 * Optionalなので、null許容
	 */
	Environments []BuildEnvironment `datastore:",noindex"`

	/**
	 * 作成日時
	 */
	CreatedDate time.Time

	/**
	 * 更新日時
	 */
	ModifiedDate time.Time
}

type BuildInfoArray []BuildInfo

/**
 * ビルド中に送信されたリアルタイムログ
 */
type BuildLog struct {
	/**
	 * 所属しているビルド
	 */
	Owner *datastore.Key `datastore:"-" goon:"parent"`

	/**
	 * ランダムに割り当てられるログID
	 */
	Id string `datastore:"-" goon:"id"`

	/**
	 * コマンドラインのログ結果が行ごとに表示される
	 */
	Lines []string `datstore:",noindex"`
}

func (it *BuildInfo) ToApiModel() api_v1.BuildInfo {
	result := api_v1.BuildInfo{
		Id:        swag.String(it.Id),
		State:     it.State.Ptr(),
		Config:    swag.String(it.Config),
		StartDate: swag.Int64(utils.Milliseconds(it.CreatedDate)),
	}

	// support git
	if len(it.Repository.Git) > 0 {
		result.Repository = &api_v1.BuildRepository{
			Git: swag.String(it.Repository.Git),
		}
		if len(it.Repository.GitRevision) > 0 {
			result.Repository.GitRevision = swag.String(it.Repository.GitRevision)
		}
	}

	if len(it.Environments) > 0 {
		env := make(api_v1.EnvironmentValueArray, len(it.Environments))
		for index, item := range it.Environments {
			env[index] = api_v1.EnvironmentValue{
				Name:  swag.String(item.Name),
				Value: swag.String(item.Value),
			}
		}
		result.Environment = &env
	}

	return result
}

func (it *BuildInfoArray) ToApiModel() api_v1.BuildInfoArray {
	result := api_v1.BuildInfoArray{}

	for _, item := range *it {
		result = append(result, item.ToApiModel())
	}

	return result

}

package server

import (
	"data"
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/gaefire/factory"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/go-yaml/yaml"
	"net/http"
	"utils"
)

/**
 * GAEインスタンス1つにつき1オブジェクト作られるサーバー全体のデータ
 */
type ApplicationImpl struct {
	/**
	 * サーバーのビルド情報
	 */
	Config data.Configuration

	/**
	 * データパース制御
	 */
	RequestEncoder *data.RequestEncoder

	/**
	 * Asset管理
	 */
	Assets gaefire.AssetManager

	/**
	 * 認証情報
	 */
	FirebaseServiceAccount gaefire.ServiceAccount

	/**
	 * 認証サポート
	 */
	AuthProxy gaefire.AuthenticationProxy
}

/**
 * 1ハンドリングごとのコンテキストを生成する
 */
func (it *ApplicationImpl) NewContext(request *http.Request) swagger.RequestContext {
	ctx := factory.NewContext(request)

	result := &ContextImpl{
		Context: ctx,
		App:     it,
		Options: data.RequestOptions{
			Context: ctx.GetAppengineContext(),
			Auth:    nil,
		},
	}

	return result
}

/**
 * ビルド情報を取得する
 */
func (it *ApplicationImpl) GetConfiguration() *data.Configuration {
	return &it.Config
}

/**
 * Firebase Service Account管理を取得する
 */
func (it *ApplicationImpl) GetFirebaseServiceAccount() gaefire.ServiceAccount {
	return it.FirebaseServiceAccount
}

func NewApplication() *ApplicationImpl {
	result := &ApplicationImpl{}

	// サービスアカウントの作成
	result.Assets = factory.NewAssetManager()
	result.RequestEncoder = data.NewRequestEncoder()

	if yamlBuf, err := result.Assets.LoadFile("assets/config.yaml"); err == nil {
		if err := yaml.Unmarshal(yamlBuf, &result.Config.Build); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	if json, err := result.Assets.LoadFile("assets/firebase-admin.json"); err == nil {
		result.FirebaseServiceAccount = factory.NewServiceAccount(json)
	} else {
		panic(err)
	}

	if json, err := result.Assets.LoadFile("assets/swagger.json"); err == nil {
		opt := gaefire.AuthenticationProxyOption{
			//EndpointsId:result.Config.Build.Env.GcpProjectId,
			EndpointsId: utils.GetGcpProjectId() + ".appspot.com",
		}
		result.AuthProxy = factory.NewAuthenticationProxyWithOption(result.FirebaseServiceAccount, opt, json)
	} else {
		panic(err)
	}

	return result
}

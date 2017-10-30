package server

import (
	"api_v1"
	"data"
	"fmt"
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"github.com/eaglesakura/swagger-go-core/swag"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"net/http"
	"utils"
)

const (
	_CONTENT_TYPE_APPLICATION_JSON = "application/json"
)

/**
 * 1リクエストにつき1オブジェクト生成される、リクエスト単位のContext
 */
type ContextImpl struct {
	/**
	 * GAE/Go context
	 */
	Context gaefire.Context

	/**
	 * 認証オプション
	 */
	Options data.RequestOptions

	/**
	 * Goon Obj
	 */
	Goon *goon.Goon

	/**
	 * アプリ本体データ
	 */
	App *ApplicationImpl
}

/**
 * リクエストごとのオプション情報を取得する
 * 認証チェック等はこれで行う。
 */
func (it *ContextImpl) GetOptions() *data.RequestOptions {
	return &it.Options
}

/**
 * リクエストごとのオプション情報を取得する
 * 認証チェック等はこれで行う。
 */
func (it *ContextImpl) GetGoon() *goon.Goon {
	if it.Goon == nil {
		it.Goon = goon.FromContext(it.Options.Context)
	}
	return it.Goon
}

/**
 * request -> parameterへのバインド制御インターフェースを生成する
 *
 * ex) swagger.NewRequestBinder()
 */
func (it *ContextImpl) NewRequestBinder(request *http.Request) (swagger.RequestBinder, error) {
	// 事前チェックを行う
	auth, err := it.App.AuthProxy.Verify(it.Options.Context, request)
	if err != nil {
		it.LogError("Auth verify error[%v]", err.Error())
		return nil, err
	}

	// 認証情報の出力
	it.LogDebug("Auth user authorized")
	it.LogDebug("  - Token apiKey[%v]", swag.StringValue(auth.ApiKey))

	it.Options.Auth = auth
	if swag.StringValue(auth.ApiKey) != "" {
		it.Options.MachineId = utils.ToMD5(*auth.ApiKey)
	}

	binder := swagger_utils.NewRequestBinder(request, func(contentType string) swagger.Consumer {
		return it.App.RequestEncoder.GetConsumer(contentType)
	})

	return binder, nil
}

/**
 * Request -> Parameterのバインド失敗時に呼び出される。
 */
func (it *ContextImpl) NewBindErrorResponse(err error) swagger.Responder {
	requestID := appengine.RequestID(it.GetOptions().Context)
	it.LogError("Bind Error[%v] Req[%v]", err.Error(), requestID)

	panicError, _ := err.(*errors.PanicError)
	if panicError != nil {
		return &swagger_utils.RawBufferResponse{
			StatusCode:  http.StatusInternalServerError,
			ContentType: "plain/text",
			Payload:     []byte(fmt.Sprintf("Server error ref[%v]", requestID)),
		}
	} else {
		return utils.NewApiErrorResponse(api_v1.ApiErrorCodeEnum_ParameterError, "")
	}
}

/**
 * ハンドリングの完了処理を行う。
 *
 * このメソッドは制御の最後にかならず呼び出される。
 * 必要に応じてリソースの開放処理を行う。
 */
func (it *ContextImpl) Done(writer http.ResponseWriter, response swagger.Responder) {
	if response != nil {
		response.WriteResponse(writer, it.App.RequestEncoder.GetProducer(_CONTENT_TYPE_APPLICATION_JSON))
	} else {
		writer.WriteHeader(http.StatusNotImplemented)
	}

	it.Context.Close()
	it.Context = nil
}

/**
 * アプリ本体を取得する
 */
func (it *ContextImpl) GetApp() data.Application {
	return it.App
}

func (it *ContextImpl) GetAssets() gaefire.AssetManager {
	return it.App.Assets
}

/**
 * エラーログ出力を行なう
 */
func (it *ContextImpl) LogError(fmt string, args ...interface{}) {
	it.Context.LogError(fmt, args...)
}

/**
 * デバッグログ出力を行なう
 */
func (it *ContextImpl) LogDebug(fmt string, args ...interface{}) {
	it.Context.LogDebug(fmt, args...)
}

/**
 * インフォログ出力を行なう
 */
func (it *ContextImpl) LogInfo(fmt string, args ...interface{}) {
	it.Context.LogInfo(fmt, args...)
}

package api_v1

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type MachineBootPostParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
}

/**
 *
 * APIキーに紐付いたビルドマシン用を起動する。 主にgithubのpushに反応し、事前にビルドマシンを起動するために使用する。
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @return void
 */
type MachineBootPostHandler func(context swagger.RequestContext, params *MachineBootPostParams) swagger.Responder

func (it *MachineBootPostParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachineBootPostParams(binder swagger.RequestBinder) (*MachineBootPostParams, error) {
	result := &MachineBootPostParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachineDeleteParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
}

/**
 *
 * APIキーに紐付いたビルドマシンを削除する
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @return void
 */
type MachineDeleteHandler func(context swagger.RequestContext, params *MachineDeleteParams) swagger.Responder

func (it *MachineDeleteParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachineDeleteParams(binder swagger.RequestBinder) (*MachineDeleteParams, error) {
	result := &MachineDeleteParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachineGetParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
}

/**
 *
 * APIキーに紐付いたビルドマシンを取得する
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @return MachineInfo
 */
type MachineGetHandler func(context swagger.RequestContext, params *MachineGetParams) swagger.Responder

func (it *MachineGetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachineGetParams(binder swagger.RequestBinder) (*MachineGetParams, error) {
	result := &MachineGetParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachinePostParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	//
	Payload *MachineRequest
}

/**
 *
 * APIキーに紐付いたビルドマシンを作成する 既に作成済みの場合、何も行なわない
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @param Payload
 * @return MachineInfo
 */
type MachinePostHandler func(context swagger.RequestContext, params *MachinePostParams) swagger.Responder

func (it *MachinePostParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Payload, it.Payload == nil).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachinePostParams(binder swagger.RequestBinder) (*MachinePostParams, error) {
	result := &MachinePostParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if err := binder.BindBody("MachineRequest", &result.Payload); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachineStartupscriptGetParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
}

/**
 *
 * APIキーに紐付いたビルドマシン用の起動スクリプトを取得する
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @return void
 */
type MachineStartupscriptGetHandler func(context swagger.RequestContext, params *MachineStartupscriptGetParams) swagger.Responder

func (it *MachineStartupscriptGetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachineStartupscriptGetParams(binder swagger.RequestBinder) (*MachineStartupscriptGetParams, error) {
	result := &MachineStartupscriptGetParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachineStartupscriptPostParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// 起動時の実行スクリプト
	Script *string
}

/**
 *
 * APIキーに紐付いたビルドマシン用の起動スクリプトを設定する。
 * @param Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 * @param Script 起動時の実行スクリプト
 * @return void
 */
type MachineStartupscriptPostHandler func(context swagger.RequestContext, params *MachineStartupscriptPostParams) swagger.Responder

func (it *MachineStartupscriptPostParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Script, it.Script == nil).Required(true).MinLength(1).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewMachineStartupscriptPostParams(binder swagger.RequestBinder) (*MachineStartupscriptPostParams, error) {
	result := &MachineStartupscriptPostParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if err := binder.BindForm("script", "string", &result.Script); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type MachineApiController struct {
	MachineBootPost swagger.HandleRequest

	MachineDelete swagger.HandleRequest

	MachineGet swagger.HandleRequest

	MachinePost swagger.HandleRequest

	MachineStartupscriptGet swagger.HandleRequest

	MachineStartupscriptPost swagger.HandleRequest
}

func NewMachineApiController() *MachineApiController {
	result := &MachineApiController{}

	result.MachineBootPost.Path = "/api/v1/machine/boot"
	result.MachineBootPost.Method = strings.ToUpper("Post")
	result.HandleMachineBootPost(func(context swagger.RequestContext, params *MachineBootPostParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachineBootPost"))
	})

	result.MachineDelete.Path = "/api/v1/machine"
	result.MachineDelete.Method = strings.ToUpper("Delete")
	result.HandleMachineDelete(func(context swagger.RequestContext, params *MachineDeleteParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachineDelete"))
	})

	result.MachineGet.Path = "/api/v1/machine"
	result.MachineGet.Method = strings.ToUpper("Get")
	result.HandleMachineGet(func(context swagger.RequestContext, params *MachineGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachineGet"))
	})

	result.MachinePost.Path = "/api/v1/machine"
	result.MachinePost.Method = strings.ToUpper("Post")
	result.HandleMachinePost(func(context swagger.RequestContext, params *MachinePostParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachinePost"))
	})

	result.MachineStartupscriptGet.Path = "/api/v1/machine/startupscript"
	result.MachineStartupscriptGet.Method = strings.ToUpper("Get")
	result.HandleMachineStartupscriptGet(func(context swagger.RequestContext, params *MachineStartupscriptGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachineStartupscriptGet"))
	})

	result.MachineStartupscriptPost.Path = "/api/v1/machine/startupscript"
	result.MachineStartupscriptPost.Method = strings.ToUpper("Post")
	result.HandleMachineStartupscriptPost(func(context swagger.RequestContext, params *MachineStartupscriptPostParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl MachineStartupscriptPost"))
	})

	return result
}

func (it *MachineApiController) HandleMachineBootPost(handler MachineBootPostHandler) {
	it.MachineBootPost.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachineBootPostParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) HandleMachineDelete(handler MachineDeleteHandler) {
	it.MachineDelete.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachineDeleteParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) HandleMachineGet(handler MachineGetHandler) {
	it.MachineGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachineGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) HandleMachinePost(handler MachinePostHandler) {
	it.MachinePost.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachinePostParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) HandleMachineStartupscriptGet(handler MachineStartupscriptGetHandler) {
	it.MachineStartupscriptGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachineStartupscriptGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) HandleMachineStartupscriptPost(handler MachineStartupscriptPostHandler) {
	it.MachineStartupscriptPost.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewMachineStartupscriptPostParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *MachineApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.MachineBootPost)

	mapper.PutHandler(it.MachineDelete)

	mapper.PutHandler(it.MachineGet)

	mapper.PutHandler(it.MachinePost)

	mapper.PutHandler(it.MachineStartupscriptGet)

	mapper.PutHandler(it.MachineStartupscriptPost)

}

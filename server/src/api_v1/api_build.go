package api_v1

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type BuildsBuildidArtifactGetParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// ビルドID
	BuildId *string
}

/*


指定IDのビルド成果物を取得する
 param: Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 param: BuildId ビルドID
 return: void
*/
type BuildsBuildidArtifactGetHandler func(context swagger.RequestContext, params *BuildsBuildidArtifactGetParams) swagger.Responder

func (it *BuildsBuildidArtifactGetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.BuildId, it.BuildId == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewBuildsBuildidArtifactGetParams(binder swagger.RequestBinder) (*BuildsBuildidArtifactGetParams, error) {
	result := &BuildsBuildidArtifactGetParams{}

	if err := binder.BindPath("buildId", "string", &result.BuildId); err != nil {
		return nil, err
	}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type BuildsBuildidGetParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// ビルドID
	BuildId *string
}

/*


指定IDのビルドを取得する
 param: Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 param: BuildId ビルドID
 return: BuildInfo
*/
type BuildsBuildidGetHandler func(context swagger.RequestContext, params *BuildsBuildidGetParams) swagger.Responder

func (it *BuildsBuildidGetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.BuildId, it.BuildId == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewBuildsBuildidGetParams(binder swagger.RequestBinder) (*BuildsBuildidGetParams, error) {
	result := &BuildsBuildidGetParams{}

	if err := binder.BindPath("buildId", "string", &result.BuildId); err != nil {
		return nil, err
	}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type BuildsBuildidPatchParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// ビルドID
	BuildId *string
	// 差分更新するビルド情報  値がsetされているパラメータのみを上書きする。 変更不可なパラメータに対しては何も行なわない（エラーとも扱わない）
	NewObject *BuildInfo
}

/*


指定IDのビルドを更新する
 param: Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 param: BuildId ビルドID
 param: NewObject 差分更新するビルド情報  値がsetされているパラメータのみを上書きする。 変更不可なパラメータに対しては何も行なわない（エラーとも扱わない）
 return: BuildInfo
*/
type BuildsBuildidPatchHandler func(context swagger.RequestContext, params *BuildsBuildidPatchParams) swagger.Responder

func (it *BuildsBuildidPatchParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.BuildId, it.BuildId == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.NewObject, it.NewObject == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewBuildsBuildidPatchParams(binder swagger.RequestBinder) (*BuildsBuildidPatchParams, error) {
	result := &BuildsBuildidPatchParams{}

	if err := binder.BindPath("buildId", "string", &result.BuildId); err != nil {
		return nil, err
	}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if err := binder.BindBody("BuildInfo", &result.NewObject); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type BuildsGetParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// 列挙するビルドステータス
	State *string
}

/*


指定条件のビルドを取得する
 param: Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 param: State 列挙するビルドステータス
 return: BuildInfoArray
*/
type BuildsGetHandler func(context swagger.RequestContext, params *BuildsGetParams) swagger.Responder

func (it *BuildsGetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.State, it.State == nil).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewBuildsGetParams(binder swagger.RequestBinder) (*BuildsGetParams, error) {
	result := &BuildsGetParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}
	if err := binder.BindQuery("state", "string", &result.State); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type BuildsPostParams struct {
	// クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	Key *string
	// ビルド情報
	Payload *BuildRequest
}

/*


新規にビルドを開始させる。
 param: Key クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
 param: Payload ビルド情報
 return: BuildInfo
*/
type BuildsPostHandler func(context swagger.RequestContext, params *BuildsPostParams) swagger.Responder

func (it *BuildsPostParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Key, it.Key == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Payload, it.Payload == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewBuildsPostParams(binder swagger.RequestBinder) (*BuildsPostParams, error) {
	result := &BuildsPostParams{}

	if err := binder.BindQuery("key", "string", &result.Key); err != nil {
		return nil, err
	}

	if err := binder.BindBody("BuildRequest", &result.Payload); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type BuildApiController struct {
	BuildsBuildidArtifactGet swagger.HandleRequest

	BuildsBuildidGet swagger.HandleRequest

	BuildsBuildidPatch swagger.HandleRequest

	BuildsGet swagger.HandleRequest

	BuildsPost swagger.HandleRequest
}

func NewBuildApiController() *BuildApiController {
	result := &BuildApiController{}

	result.BuildsBuildidArtifactGet.Path = "/api/v1/builds/{buildId}/artifact"
	result.BuildsBuildidArtifactGet.Method = strings.ToUpper("Get")
	result.HandleBuildsBuildidArtifactGet(func(context swagger.RequestContext, params *BuildsBuildidArtifactGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl BuildsBuildidArtifactGet"))
	})

	result.BuildsBuildidGet.Path = "/api/v1/builds/{buildId}"
	result.BuildsBuildidGet.Method = strings.ToUpper("Get")
	result.HandleBuildsBuildidGet(func(context swagger.RequestContext, params *BuildsBuildidGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl BuildsBuildidGet"))
	})

	result.BuildsBuildidPatch.Path = "/api/v1/builds/{buildId}"
	result.BuildsBuildidPatch.Method = strings.ToUpper("Patch")
	result.HandleBuildsBuildidPatch(func(context swagger.RequestContext, params *BuildsBuildidPatchParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl BuildsBuildidPatch"))
	})

	result.BuildsGet.Path = "/api/v1/builds"
	result.BuildsGet.Method = strings.ToUpper("Get")
	result.HandleBuildsGet(func(context swagger.RequestContext, params *BuildsGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl BuildsGet"))
	})

	result.BuildsPost.Path = "/api/v1/builds"
	result.BuildsPost.Method = strings.ToUpper("Post")
	result.HandleBuildsPost(func(context swagger.RequestContext, params *BuildsPostParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl BuildsPost"))
	})

	return result
}

func (it *BuildApiController) HandleBuildsBuildidArtifactGet(handler BuildsBuildidArtifactGetHandler) {
	it.BuildsBuildidArtifactGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewBuildsBuildidArtifactGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *BuildApiController) HandleBuildsBuildidGet(handler BuildsBuildidGetHandler) {
	it.BuildsBuildidGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewBuildsBuildidGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *BuildApiController) HandleBuildsBuildidPatch(handler BuildsBuildidPatchHandler) {
	it.BuildsBuildidPatch.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewBuildsBuildidPatchParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *BuildApiController) HandleBuildsGet(handler BuildsGetHandler) {
	it.BuildsGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewBuildsGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *BuildApiController) HandleBuildsPost(handler BuildsPostHandler) {
	it.BuildsPost.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewBuildsPostParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *BuildApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.BuildsBuildidArtifactGet)

	mapper.PutHandler(it.BuildsBuildidGet)

	mapper.PutHandler(it.BuildsBuildidPatch)

	mapper.PutHandler(it.BuildsGet)

	mapper.PutHandler(it.BuildsPost)

}

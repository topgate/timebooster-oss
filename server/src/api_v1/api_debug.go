package api_v1

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type DebugSwaggerGetParams struct {
}

/*


swagger.json
 return: void
*/
type DebugSwaggerGetHandler func(context swagger.RequestContext, params *DebugSwaggerGetParams) swagger.Responder

func (it *DebugSwaggerGetParams) Valid(factory swagger.ValidatorFactory) bool {

	return true
}

// Bind from request
func NewDebugSwaggerGetParams(binder swagger.RequestBinder) (*DebugSwaggerGetParams, error) {
	result := &DebugSwaggerGetParams{}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type DebugApiController struct {
	DebugSwaggerGet swagger.HandleRequest
}

func NewDebugApiController() *DebugApiController {
	result := &DebugApiController{}

	result.DebugSwaggerGet.Path = "/api/v1/debug/swagger"
	result.DebugSwaggerGet.Method = strings.ToUpper("Get")
	result.HandleDebugSwaggerGet(func(context swagger.RequestContext, params *DebugSwaggerGetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl DebugSwaggerGet"))
	})

	return result
}

func (it *DebugApiController) HandleDebugSwaggerGet(handler DebugSwaggerGetHandler) {
	it.DebugSwaggerGet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewDebugSwaggerGetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *DebugApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.DebugSwaggerGet)

}

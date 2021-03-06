package api_v1

// generated by lightweight-swagger-codegen@eaglesakura

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"github.com/eaglesakura/swagger-go-core/utils"
	"net/url"
	"strings"
)

const DebugApi_BasePath string = "/api/v1"

type DebugApi struct {
	BasePath string
}

func NewDebugApi() *DebugApi {
	return &DebugApi{
		BasePath: DebugApi_BasePath,
	}
}

/*

   swagger.json
*/
type DebugApiDebugSwaggerGetRequest struct {
}

/*

   swagger.json

     result: void
*/
func (it *DebugApi) DebugSwaggerGet(_client swagger.FetchClient, _request *DebugApiDebugSwaggerGetRequest, result interface{}) error {

	// create path and map variables
	{
		localVarPath := strings.Replace("/debug/swagger", "{format}", "json", -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Get"))
	}

	return _client.Fetch(result)
}

func (it *DebugApi) this_is_call_dummy() {
	v := url.Values{}
	v.Add("Key", "Value")

	errors.New(0, "stub")
	strings.ToUpper("")
}

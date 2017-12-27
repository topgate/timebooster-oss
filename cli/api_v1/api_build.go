package api_v1

// generated by lightweight-swagger-codegen@eaglesakura

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"github.com/eaglesakura/swagger-go-core/utils"
	"net/url"
	"strings"
)

const BuildApi_BasePath string = "/api/v1"

type BuildApi struct {
	BasePath string
}

func NewBuildApi() *BuildApi {
	return &BuildApi{
		BasePath: BuildApi_BasePath,
	}
}

/*

   指定IDのビルド成果物を取得する
*/
type BuildApiBuildsBuildidArtifactGetRequest struct {
	/*
	   クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	*/
	Key *string

	/*
	   ビルドID
	*/
	BuildId *string
}

/*

   指定IDのビルド成果物を取得する

     result: void
*/
func (it *BuildApi) BuildsBuildidArtifactGet(_client swagger.FetchClient, _request *BuildApiBuildsBuildidArtifactGetRequest, result interface{}) error {
	if !_client.NewValidator(_request.Key, _request.Key == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Key' when calling BuildsBuildidArtifactGet")
	}
	if !_client.NewValidator(_request.BuildId, _request.BuildId == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'BuildId' when calling BuildsBuildidArtifactGet")
	}

	// create path and map variables
	{
		localVarPath := strings.Replace("/builds/{buildId}/artifact", "{format}", "json", -1)
		localVarPath = strings.Replace(localVarPath, "{"+"buildId"+"}", utils.EscapeString(*_request.BuildId), -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Get"))
	}

	if _request.Key != nil {
		_client.AddQueryParam("key", utils.ParameterToString(_request.Key))
	}

	return _client.Fetch(result)
}

/*

   指定IDのビルドを取得する
*/
type BuildApiBuildsBuildidGetRequest struct {
	/*
	   クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	*/
	Key *string

	/*
	   ビルドID
	*/
	BuildId *string
}

/*

   指定IDのビルドを取得する

     result: BuildInfo
*/
func (it *BuildApi) BuildsBuildidGet(_client swagger.FetchClient, _request *BuildApiBuildsBuildidGetRequest, result *BuildInfo) error {
	if !_client.NewValidator(_request.Key, _request.Key == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Key' when calling BuildsBuildidGet")
	}
	if !_client.NewValidator(_request.BuildId, _request.BuildId == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'BuildId' when calling BuildsBuildidGet")
	}

	// create path and map variables
	{
		localVarPath := strings.Replace("/builds/{buildId}", "{format}", "json", -1)
		localVarPath = strings.Replace(localVarPath, "{"+"buildId"+"}", utils.EscapeString(*_request.BuildId), -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Get"))
	}

	if _request.Key != nil {
		_client.AddQueryParam("key", utils.ParameterToString(_request.Key))
	}

	return _client.Fetch(result)
}

/*

   指定IDのビルドを更新する
*/
type BuildApiBuildsBuildidPatchRequest struct {
	/*
	   クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	*/
	Key *string

	/*
	   ビルドID
	*/
	BuildId *string

	/*
	   差分更新するビルド情報  値がsetされているパラメータのみを上書きする。 変更不可なパラメータに対しては何も行なわない（エラーとも扱わない）
	*/
	NewObject *BuildInfo
}

/*

   指定IDのビルドを更新する

     result: BuildInfo
*/
func (it *BuildApi) BuildsBuildidPatch(_client swagger.FetchClient, _request *BuildApiBuildsBuildidPatchRequest, result *BuildInfo) error {
	if !_client.NewValidator(_request.Key, _request.Key == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Key' when calling BuildsBuildidPatch")
	}
	if !_client.NewValidator(_request.BuildId, _request.BuildId == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'BuildId' when calling BuildsBuildidPatch")
	}
	if !_client.NewValidator(_request.NewObject, _request.NewObject == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'NewObject' when calling BuildsBuildidPatch")
	}

	// create path and map variables
	{
		localVarPath := strings.Replace("/builds/{buildId}", "{format}", "json", -1)
		localVarPath = strings.Replace(localVarPath, "{"+"buildId"+"}", utils.EscapeString(*_request.BuildId), -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Patch"))
	}

	if _request.Key != nil {
		_client.AddQueryParam("key", utils.ParameterToString(_request.Key))
	}

	if _request.NewObject != nil {
		_client.SetPayload(utils.NewJsonPayload(_request.NewObject))
	}

	return _client.Fetch(result)
}

/*

   指定条件のビルドを取得する
*/
type BuildApiBuildsGetRequest struct {
	/*
	   クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	*/
	Key *string

	/*
	   列挙するビルドステータス
	*/
	State *string
}

/*

   指定条件のビルドを取得する

     result: BuildInfoArray
*/
func (it *BuildApi) BuildsGet(_client swagger.FetchClient, _request *BuildApiBuildsGetRequest, result *BuildInfoArray) error {
	if !_client.NewValidator(_request.Key, _request.Key == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Key' when calling BuildsGet")
	}
	if !_client.NewValidator(_request.State, _request.State == nil).Valid(_client) {
		errors.New(0, "Missing the required parameter 'State' when calling BuildsGet")
	}

	// create path and map variables
	{
		localVarPath := strings.Replace("/builds", "{format}", "json", -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Get"))
	}

	if _request.Key != nil {
		_client.AddQueryParam("key", utils.ParameterToString(_request.Key))
	}
	if _request.State != nil {
		_client.AddQueryParam("state", utils.ParameterToString(_request.State))
	}

	return _client.Fetch(result)
}

/*

   新規にビルドを開始させる。
*/
type BuildApiBuildsPostRequest struct {
	/*
	   クライアントの妥当性を検証するためのAPIKey  発行されたAPIKey以外はAPIを呼び出すことはできない。
	*/
	Key *string

	/*
	   ビルド情報
	*/
	Payload *BuildRequest
}

/*

   新規にビルドを開始させる。

     result: BuildInfo
*/
func (it *BuildApi) BuildsPost(_client swagger.FetchClient, _request *BuildApiBuildsPostRequest, result *BuildInfo) error {
	if !_client.NewValidator(_request.Key, _request.Key == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Key' when calling BuildsPost")
	}
	if !_client.NewValidator(_request.Payload, _request.Payload == nil).Required(true).Valid(_client) {
		errors.New(0, "Missing the required parameter 'Payload' when calling BuildsPost")
	}

	// create path and map variables
	{
		localVarPath := strings.Replace("/builds", "{format}", "json", -1)
		_client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
		_client.SetMethod(strings.ToUpper("Post"))
	}

	if _request.Key != nil {
		_client.AddQueryParam("key", utils.ParameterToString(_request.Key))
	}

	if _request.Payload != nil {
		_client.SetPayload(utils.NewJsonPayload(_request.Payload))
	}

	return _client.Fetch(result)
}

func (it *BuildApi) this_is_call_dummy() {
	v := url.Values{}
	v.Add("Key", "Value")

	errors.New(0, "stub")
	strings.ToUpper("")
}

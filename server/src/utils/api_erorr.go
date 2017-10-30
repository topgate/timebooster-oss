package utils

import (
	"api_v1"
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/swag"
	"net/http"
)

/**
 * API内部で使用するエラー返却値
 */
type InternalErrorResponse struct {
	StatusCode int
	ApiError   *api_v1.ApiError
}

func (it *InternalErrorResponse) Code() api_v1.ApiErrorCodeEnum {
	return *it.ApiError.Code
}

func (it *InternalErrorResponse) Message() string {
	return *it.ApiError.Message
}

/**
 * エラーレスポンスを生成する
 */
func (it *InternalErrorResponse) WriteResponse(w http.ResponseWriter, p swagger.Producer) {
	w.WriteHeader(it.StatusCode)
	w.Header().Add("Content-Type", "application/json")

	buf, _ := json.Marshal(it.ApiError)
	w.Write(buf)
}

/**
 * エラーレスポンスを生成する
 * API戻り値はapi_v1.ApiError構造体のデータ互換となる
 */
func NewApiErrorResponse(code api_v1.ApiErrorCodeEnum, message string) swagger.Responder {

	statusCode := http.StatusNotImplemented
	defMessage := ""
	switch code {
	case api_v1.ApiErrorCodeEnum_AuthFailed:
		statusCode = http.StatusForbidden
		defMessage = "Required login"
	case api_v1.ApiErrorCodeEnum_ParameterError:
		statusCode = http.StatusBadRequest
		defMessage = "Parameter bind error"
	case api_v1.ApiErrorCodeEnum_DataConflict:
		statusCode = http.StatusInternalServerError
		defMessage = "Conflict Datastore?"
	case api_v1.ApiErrorCodeEnum_DataModifyFailed:
		statusCode = http.StatusInternalServerError
		defMessage = "Conflict Datastore?"
	case api_v1.ApiErrorCodeEnum_Unknown:
		statusCode = http.StatusInternalServerError
		defMessage = "Unknown!"
	}

	if len(message) == 0 {
		message = defMessage
	}

	return &InternalErrorResponse{
		StatusCode: statusCode,
		ApiError: &api_v1.ApiError{
			Code:    code.Ptr(),
			Message: swag.String(message),
		},
	}
}

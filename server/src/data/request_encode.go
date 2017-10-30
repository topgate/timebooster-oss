package data

import (
	"github.com/eaglesakura/swagger-go-core"
	swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
)

/**
 * リクエストに対してデータエンコード / デコードを提供する
 */
type RequestEncoder struct {
	jsonConsumer swagger.Consumer
	jsonProducer swagger.Producer
}

/**
 * エンコーダを生成する
 */
func NewRequestEncoder() *RequestEncoder {
	return &RequestEncoder{
		jsonConsumer: swagger_utils.NewJsonConsumer(),
		jsonProducer: swagger_utils.NewJsonProducer(),
	}
}

func (it *RequestEncoder) GetConsumer(contentType string) swagger.Consumer {
	return it.jsonConsumer
}

func (it *RequestEncoder) GetProducer(contentType string) swagger.Producer {
	return it.jsonProducer
}

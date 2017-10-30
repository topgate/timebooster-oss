package data

import (
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/swagger-go-core"
)

type Application interface {
	swagger.ContextFactory

	/**
	 * Firebase Service Account管理を取得する
	 */
	GetFirebaseServiceAccount() gaefire.ServiceAccount

	/**
	 * ビルド情報を取得する
	 */
	GetConfiguration() *Configuration
}

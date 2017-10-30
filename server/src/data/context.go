package data

import (
	"github.com/eaglesakura/gaefire"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
)

type RequestOptions struct {
	/**
	 * ユーザー認証情報（認証されていれば != nil）
	 */
	Auth *gaefire.AuthenticationInfo

	/**
	 * APIキーに割り当てられたビルドマシンID
	 */
	MachineId string

	/**
	 * リクエストごとのContext
	 */
	Context context.Context
}

func (it *RequestOptions) GetUserId() string {
	if it.Auth.User != nil {
		return swag.StringValue(it.Auth.User.Id)
	} else {
		return ""
	}
}

/**
 * セキュアアクセスチェック
 */
func (it *RequestOptions) IsSecure() bool {
	return swag.StringValue(it.Auth.ApiKey) != ""
}

type Context interface {
	/**
	 * Swagger Requestを兼ねる
	 */
	swagger.RequestContext

	/**
	 * リクエストごとのオプション情報を取得する
	 * 認証チェック等はこれで行う。
	 */
	GetOptions() *RequestOptions

	/**
	 * アプリ本体を取得する
	 */
	GetApp() Application

	/**
	 * Goon Instanceを取得する
	 */
	GetGoon() *goon.Goon

	/**
	 * アセットロードを行なう
	 */
	GetAssets() gaefire.AssetManager

	/**
	 * エラーログ出力を行なう
	 */
	LogError(fmt string, args ...interface{})

	/**
	 * デバッグログ出力を行なう
	 */
	LogDebug(fmt string, args ...interface{})

	/**
	 * インフォログ出力を行なう
	 */
	LogInfo(fmt string, args ...interface{})
}

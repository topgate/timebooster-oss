package data

import "time"

/**
 * サーバーのビルド情報など、固定値/設定値を取得する
 */
type Configuration struct {
	Build BuildConfig
}

/**
 * ビルド情報の固定値
 */
type BuildConfig struct {
	Env struct {
		/**
		 * ビルド番号
		 */
		CiVersion int

		/**
		 * ビルド日時
		 */
		BuildDate string

		/**
		 * gitリビジョン
		 */
		Revision string
	}
}

func (it *Configuration) GetBuildInfo() BuildConfig {
	return it.Build
}

func (it *BuildConfig) GetBuildDate() time.Time {
	result, err := time.Parse("2006-01-02 03:04.05", it.Env.BuildDate)
	if err != nil {
		return time.Time{}
	} else {
		return result
	}
}

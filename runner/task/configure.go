package task

/**
 * ビルド設定ファイル
 * Yamlで記述する
 */
type Configure struct {
	Env struct {
		Repository string                 `yaml:"repository"`
		Revision   string                 `yaml:"revision`
		Variable   map[string]interface{} `yaml:"variable"`
		Cache      []string               `yaml:"cache"`
	}
	Task struct {
		Exec []struct {
			/**
			 * 実行されるコンテナ定義
			 */
			Dockerfile string `yaml:"dockerfile"`

			/**
			 * 実行されるコンテナ定義
			 */
			DockerImage string `yaml:"dockerimage"`

			/**
			 * 実行されるコマンド一覧
			 */
			Cmd []string `yaml:"cmd"`
		}
	}
}

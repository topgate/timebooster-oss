package model

/**
 * ビルドマシン情報
 * APIキーごとに1Entityが生成される
 */
type BuildMachine struct {
	/**
	 * マシンごとに一意に生成される管理ID
	 */
	Id string `datastore:"-" goon:"id"`

	/**
	 * 起動スクリプト
	 * emptyの場合、デフォルトの起動スクリプトが使用される
	 */
	StartupScript string `datastore:",noindex"`

	/**
	 * 使用されているZone
	 */
	Zone string `datastore:",noindex"`
}

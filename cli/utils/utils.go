package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

//
// 文字列をMD5に変換する
//
func ToMD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

/**
 * ファイルのMD5を計算する
 */
func ToFileMD5(path string) string {
	buf, _ := ioutil.ReadFile(path)
	if buf == nil {
		return ""
	} else {
		hash := md5.New()
		hash.Write(buf)
		return hex.EncodeToString(hash.Sum(nil))
	}
}

/**
 * リポジトリURLをユーザー名とリポジトリ名に分解する
 */
func SplitGithubRepository(repository string) (user, repo string) {
	if strings.Index(repository, "git@github.com:") == 0 {
		repository = repository[len("git@github.com:"):]
		parts := strings.Split(repository, "/")
		user = parts[0]
		repo = parts[1]
	} else {
		// http
		parts := strings.Split(repository, "/")
		user = parts[len(parts)-2]
		repo = parts[len(parts)-1]
	}

	if strings.Index(repo, ".git") > 0 {
		// .gitを削除する
		repo = repo[:len(repo)-len(".git")]
	}

	return user, repo
}

/**
 * リポジトリURLとapiKeyから実際のURLを取得する
 * git@github.com:user/repo.git
 * https://github.com/user/repo.git
 * https://GITHUB_API_KEY:x-oauth-basic@github.com/user/repo.git
 */
func GetGithubRepositoryPath(repository string, apiKey string) string {
	if len(apiKey) > 0 {
		user, repo := SplitGithubRepository(repository)
		return fmt.Sprintf("https://%v:x-oauth-basic@github.com/%v/%v.git", apiKey, user, repo)
	}
	return repository
}

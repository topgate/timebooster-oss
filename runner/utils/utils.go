package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
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

func GetApiKey() string {
	return os.Getenv("TIMEBOOSTER_API_KEY")
}

func GetServerEndpoint() string {
	return os.Getenv("TIMEBOOSTER_ENDPOINT")
}

func GetTimeboosterProjectId() string {
	return os.Getenv("TIMEBOOSTER_PROJECT_ID")
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

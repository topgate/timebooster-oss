package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

//
// 文字列をMD5に変換する
//
func ToMD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

//
// 文字列をSHA1に変換する
//
func ToSHA1(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

//
// 文字列を数値に変換する
// パースに失敗した場合、デフォルト値を返す
//
func Atoi(value string, def int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return def
	} else {
		return result
	}
}

//
// 文字列を数値に変換する
// パースに失敗した場合、デフォルト値を返す
//
func Atof(value string, def float64) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return def
	} else {
		return result
	}
}

func GetGcpProjectId() string {
	return os.Getenv("GCP_PROJECT_ID")
}

func GetBuildMachineServiceAccount() string {
	return os.Getenv("TIMEBOOSTER_SERVICE_ACCOUNT")
}

/**
 * UnixTime milliseconds
 */
func Milliseconds(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

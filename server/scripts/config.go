// +build config
package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
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

func getEnv(key string, def string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	} else {
		return def
	}
}

func getEnvI(key string, def int) int {
	value, found := os.LookupEnv(key)
	if !found {
		return def
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	} else {
		return i
	}
}

func main() {
	fmt.Printf("Config Generate...\n")

	config := ServerConfig{}

	config.Env.CiVersion = getEnvI("CIRCLE_BUILD_NUM", 1)
	config.Env.BuildDate = time.Now().Format("2006-01-02 03:04.05")
	config.Env.Revision = getEnv("CIRCLE_SHA1", "master")

	out, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Printf("Error Generate Config")
		os.Exit(1)
		return
	}

	ioutil.WriteFile("gae/assets/config.yaml", out, os.ModePerm)
}

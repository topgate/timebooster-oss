package cmd_run

import (
	"github.com/urfave/cli"
	"io/ioutil"
)

func getApiKey(ctx *cli.Context) string {
	return ctx.GlobalString("api-key")
}

func getConfigFilePath(ctx *cli.Context) string {
	return ctx.GlobalString("config")
}

func loadConfigFile(ctx *cli.Context) (string, error) {
	buf, err := ioutil.ReadFile(getConfigFilePath(ctx))
	if err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}

func getServerEndpoint(ctx *cli.Context) string {
	return ctx.GlobalString("endpoint")
}

package main

import (
	"./cmd_run"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Timebooster"
	app.Version = "1.0.__BUILD_VERSION__"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api-key",
			Usage: "Timebooster api key.",
		},
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "Timebooster Server endpoint URL",
		},
		cli.StringFlag{
			Name:  "config",
			Usage: "`timebooster.[yaml|yml]` path.",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "build on Timebooster build machine.",
			Action: func(ctx *cli.Context) {
				cmd_run.NewInstance(ctx).Execute()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repository",
					Usage: "git repository path.",
				},
				cli.StringFlag{
					Name:  "revision",
					Usage: "git repository revision.",
				},
				cli.StringFlag{
					Name:  "github-access-token",
					Usage: "github repository accesstoken.",
				},
				cli.BoolFlag{
					Name:  "print-repository",
					Usage: "for Debug, print repository info",
				},
				cli.StringFlag{
					Name:  "artifact",
					Usage: "Download artifact path.",
				},
				cli.BoolFlag{
					Name:  "print-config",
					Usage: "for Debug, print config.yaml",
				},
				cli.BoolFlag{
					Name:  "env-from-circleci",
					Usage: "CircleCI Environment variables copy to Docker container.",
				},
				cli.StringSliceFlag{
					Name:  "env",
					Usage: "Environment variables copy to Docker container.",
				},
				cli.BoolFlag{
					Name:  "print-env",
					Usage: "for Debug, print environment variables.",
				},
			},
		},
	}
	app.Run(os.Args)
}

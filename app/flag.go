package app

import (
	"github.com/rajasoun/aws-hub/app/args"
	"github.com/urfave/cli/v2"
)

func (app *App) setUpFlags() {
	flags := []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Usage: "Server port",
			Value: args.DefaultPort,
		},
		&cli.IntFlag{
			Name:  "duration, d",
			Usage: "Cache expiration time",
			Value: args.DefaultDuration,
		},
		&cli.StringFlag{
			Name:  "redis, r",
			Usage: "Redis server",
		},
		&cli.StringFlag{
			Name:  "dataset, ds",
			Usage: "BigQuery Bill dataset",
		},
		&cli.BoolFlag{
			Name:  "multiple, m",
			Usage: "Enable multiple AWS accounts",
		},
	}
	app.cli.Flags = flags
}

func (app *App) SetUpFlags() {
	app.setUpFlags()
}

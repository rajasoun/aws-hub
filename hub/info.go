package hub

import (
	"time"

	"github.com/urfave/cli/v2"
)

func (app *App) setUpIdentity() {
	app.cli.Name = "AWS Hub"
	app.cli.Usage = "AWS Cost Explorer"
	app.cli.Version = "0.0.1"
	app.cli.Compiled = time.Now()
}

func (app *App) setUpAuthors() {
	authors := []*cli.Author{
		{
			Name:  "Raja Soundaramourty",
			Email: "rajasoun@cisco.com",
		},
	}
	app.cli.Authors = authors
}

func (app *App) SetUpInfo() {
	app.setUpIdentity()
	app.setUpAuthors()
}

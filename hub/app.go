package hub

import (
	"time"

	structs "github.com/rajasoun/go-ds"
	"github.com/urfave/cli/v2"
)

type App struct {
	cli *cli.App
}

func NewApp() *App {
	app := App{
		cli: &cli.App{},
	}
	return &app
}

func (app *App) setUpApp() {
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

func (app *App) structToMap() map[string]interface{} {
	s := structs.New(app.cli)
	m := s.Map()
	return m
}

package hub

import (
	"fmt"
	"log"
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

func (app *App) setUpFlags() {
	flags := []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Usage: "Server port",
			Value: DefaultPort,
		},
		&cli.IntFlag{
			Name:  "duration, d",
			Usage: "Cache expiration time",
			Value: DefaultDuration,
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

func (app *App) setUpCommands() {
	commands := []*cli.Command{
		{
			Name:  "start",
			Usage: "Start Server",
			Flags: app.cli.Flags,
			Action: func(appCtx *cli.Context) error {
				port, cacheHandler, multiple := parseArgs(appCtx)
				loggedRouter := setUp(cacheHandler, multiple)
				start(port, loggedRouter)
				return nil
			},
		},
	}
	commandNotFound := func(appCtx *cli.Context, command string) {
		_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
		if err != nil {
			log.Println(appCtx.App.Writer, "Command not found %q !", command)
		}
	}

	app.cli.Commands = commands
	app.cli.CommandNotFound = commandNotFound
}

func (app *App) structToMap() map[string]interface{} {
	s := structs.New(app.cli)
	m := s.Map()
	return m
}

func sliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}

func Execute(args []string) {
	app := NewApp()
	app.setUpIdentity()
	app.setUpAuthors()
	app.setUpFlags()
	app.setUpCommands()
	app.cli.Run(args)
}

package hub

import (
	"fmt"
	"io"
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
	app.setUpIdentity()
	app.setUpAuthors()
	app.setUpFlags()
	app.setUpCommands()
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
				err := StartCommandRunner(appCtx)
				return err
			},
		},
	}
	commandNotFound := func(appCtx *cli.Context, command string) {
		CommandNotFoundRunner(appCtx, command)
	}

	app.cli.Commands = commands
	app.cli.CommandNotFound = commandNotFound
}

func StartCommandRunner(appCtx *cli.Context) error {
	cliContext := NewCliContext(appCtx)
	server, _ := NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	err := server.start(cliContext.Port())
	return err
}

func CommandNotFoundRunner(appCtx *cli.Context, command string) {
	_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
	if err != nil {
		log.Println(appCtx.App.Writer, "Command not found %q !", command)
	}
}

func (app *App) StructToMap() map[string]interface{} {
	s := structs.New(app.cli)
	m := s.Map()
	return m
}

func SliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}

func Execute(args []string, writer io.Writer) error {
	app := NewApp()
	app.cli.Writer = writer
	err := app.cli.Run(args)
	return err
}

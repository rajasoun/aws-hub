package hub

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

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
	err := server.Start(cliContext.Port())
	return err
}

func CommandNotFoundRunner(appCtx *cli.Context, command string) {
	_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
	if err != nil {
		log.Println(appCtx.App.Writer, "Command not found %q !", command)
	}
}

func (app *App) SetUpCommands() {
	app.setUpCommands()
}

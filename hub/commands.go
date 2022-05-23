package hub

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func startCommand(appCtx *cli.Context) {
	port, cacheHandler, multiple := parseArgs(appCtx)
	loggedRouter := setUpServer(cacheHandler, multiple)
	startServer(port, loggedRouter)
}

func getCommands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:  "start",
			Usage: "Start Server",
			Flags: getFlags(),
			Action: func(appCtx *cli.Context) error {
				startCommand(appCtx)
				return nil
			},
		},
	}
	return commands
}

func getCommandNotFound(app *cli.App) func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
		if err != nil {
			log.Println(appCtx.App.Writer, "Command not found %q !", command)
		}
	}
}

func setUpCommands(app *cli.App) {
	app.Commands = getCommands()
	app.CommandNotFound = getCommandNotFound(app)
}

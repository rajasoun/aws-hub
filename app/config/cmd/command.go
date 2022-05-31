package cmd

import (
	"fmt"
	"log"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:  "start",
			Usage: "Start Server",
			Flags: flag.GetFlags(),
			Action: func(appCtx *cli.Context) error {
				err := StartCommandRunner(appCtx)
				return err
			},
		},
	}
	return commands
}

func StartCommandRunner(appCtx *cli.Context) error {
	cliContext := arg.NewCliContext(appCtx)
	server, _ := NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
	err := server.Start(cliContext.GetPort())
	return err
}

func GetErrCommand() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
		if err != nil {
			log.Println(appCtx.App.Writer, "Command not found %q !", command)
		}
	}
}

package cmd

import (
	"fmt"
	"log"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/urfave/cli/v2"
)

// Interface for Command Handler
// This will enable TDD using mocking
type CommandHandler interface {
	StartCommandHandler(appCtx *cli.Context) error
}

func CreateStartCommand(handler func(appCtx *cli.Context) error) cli.Command {
	command := cli.Command{
		Name:  "start",
		Usage: "Start Server",
		Flags: flag.GetFlags(),
		Action: func(appCtx *cli.Context) error {
			err := handler(appCtx)
			return err
		},
	}
	return command
}
func GetCommands(handler func(appCtx *cli.Context) error) []*cli.Command {
	startCommand := CreateStartCommand(handler)
	commands := []*cli.Command{&startCommand}
	return commands
}

func StartCommandHandler(appCtx *cli.Context) error {
	cliContext := arg.NewCliContext(appCtx)
	server, _ := server.NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
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

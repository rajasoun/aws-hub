package cmd

import (
	"fmt"
	"log"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/urfave/cli/v2"
)

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
func GetCommands() []*cli.Command {
	startCommand := CreateStartCommand(StartCommandHandler)
	commands := []*cli.Command{&startCommand}
	return commands
}

func StartCommandHandler(appCtx *cli.Context) error {
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

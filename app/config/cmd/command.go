package cmd

import (
	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/urfave/cli/v2"
)

type CmdHandler struct {
	EnableShutdDown bool
}

// Create Start Command with the supplied handler
func CreateCommand(handler func(appCtx *cli.Context) error) cli.Command {
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

// Get all Commands created from the supplied handler
func GetCommand(handler func(appCtx *cli.Context) error) cli.Command {
	cmd := CreateCommand(handler)
	return cmd
}

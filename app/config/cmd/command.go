package cmd

import (
	"fmt"
	"log"

	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/urfave/cli/v2"
)

const CommandNotFoundMsg = "Invalid Command. Not Found"

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

// Get Err Command
func GetErrCommand() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		errMsg := fmt.Sprintf(CommandNotFoundMsg+" '%s' ", command)
		log.Println(errMsg)
	}
}

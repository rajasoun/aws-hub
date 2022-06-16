package cmd

import (
	"fmt"
	"log"

	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/urfave/cli/v2"
)

const CommandNotFoundMsg = "Invalid Command. Not Found"

// Create New Command with the supplied handler.
func New(name, usage string, handler func(appCtx *cli.Context) error) cli.Command {
	command := cli.Command{
		Name:  name,
		Usage: usage,
		Flags: flag.GetFlags(),
		Action: func(appCtx *cli.Context) error {
			err := handler(appCtx)
			return err
		},
	}
	return command
}

// Get Err Command.
func NewErr() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		errMsg := fmt.Sprintf(CommandNotFoundMsg+" '%s' ", command)
		log.Println(errMsg)
	}
}

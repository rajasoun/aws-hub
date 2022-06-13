package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

const CommandNotFoundMsg = "Invalid Command. Not Found"

func ErrCommand() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		errMsg := fmt.Sprintf(CommandNotFoundMsg+" '%s' ", command)
		log.Println(errMsg)
	}
}

package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func GetErrCommand() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		_, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
		if err != nil {
			log.Println(appCtx.App.Writer, "Command not found %q !", command)
		}
	}
}

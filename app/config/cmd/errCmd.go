package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
)

func ErrCommand() func(appCtx *cli.Context, command string) {
	return func(appCtx *cli.Context, command string) {
		log.Println(appCtx.App.Writer, "Command Not Found !")
		// _, err := fmt.Fprintf(appCtx.App.Writer, "Command not found %q !", command)
		// if err != nil {
		// 	log.Println(appCtx.App.Writer, "Command not found %q !", command)
		// }
	}
}

package cmd

import (
	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/urfave/cli/v2"
)

func StartCommandHandler(appCtx *cli.Context) error {
	cliContext := arg.NewCliContext(appCtx)
	server, _ := server.NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
	err := server.Start(cliContext.GetPort())
	// httpServer := server.NewHTTPServer(string(rune(cliContext.GetPort())))
	// err :=httpServer.StartHTTPServer()
	return err
}

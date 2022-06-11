package cmd

import (
	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/urfave/cli/v2"
)

type CmdHandler struct {
	EnableShutdDown bool
}

func (handler *CmdHandler) StartCommand(appCtx *cli.Context) error {
	cliContext := arg.NewCliContext(appCtx)
	server, _ := server.NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
	err := server.Start(cliContext.GetPort(), handler.EnableShutdDown)
	return err
}

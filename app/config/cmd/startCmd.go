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
	cache := cliContext.GetCache()
	isMultipleProfile := cliContext.GetAwsProfileType()
	httpServer, _ := server.NewServer(cache, isMultipleProfile)
	err := httpServer.Start(cliContext.GetPort(), handler.EnableShutdDown)
	return err
}

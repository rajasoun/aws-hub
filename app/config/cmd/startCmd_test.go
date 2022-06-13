package cmd

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStartCmd(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Command With Custom Port Zero", func(t *testing.T) {
		context := NewContext()
		startCommand := GetStartCommand()
		err := startCommand.Run(context)
		assert.NoError(err, "err = %v ", err)
	})
}

func mockStartCommand(appCtx *cli.Context) error {
	log.Println("mockStartCommandHandler !!!")
	return nil
}

func GetStartCommand() cli.Command {
	cmdhandler := CmdHandler{EnableShutdDown: true}
	startCommand := GetCommand(cmdhandler.StartCommand)
	return startCommand
}

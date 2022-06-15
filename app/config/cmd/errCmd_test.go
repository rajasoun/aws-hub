package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetErrCommand(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	t.Run("Check ErrCommand With Invalid Input", func(t *testing.T) {
		mockApp := &cli.App{Writer: ioutil.Discard}
		cmdhandler := CmdHandler{}
		cmdhandler.EnableShutdDown = false
		startCommand := GetCommand(cmdhandler.StartCommand)
		commands := []*cli.Command{&startCommand}
		mockApp.Commands = commands
		mockApp.CommandNotFound = GetErrCommand()

		err := mockApp.Run([]string{"invalidCommand"})
		assert.NoError(err, "err = %v ", err)
	})
}

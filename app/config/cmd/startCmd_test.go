package cmd

import (
	"flag"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStartCmd(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Server With Empty Context", func(t *testing.T) {
		cmdhandler := CmdHandler{}
		cmdhandler.EnableShutdDown = true
		err := cmdhandler.StartCommand(&cli.Context{})
		assert.NoError(err, "err = %v ", err)
	})
	t.Run("Check Start Server With Custom Port Zero", func(t *testing.T) {
		mockApp := &cli.App{Writer: ioutil.Discard}
		set := flag.NewFlagSet("test", 0)
		_ = set.Parse([]string{"start", "--port", "0"})
		cCtx := cli.NewContext(mockApp, set, nil)

		cmdhandler := CmdHandler{}
		cmdhandler.EnableShutdDown = true
		startCommand := GetCommand(cmdhandler.StartCommand)

		err := startCommand.Run(cCtx)
		assert.NoError(err, "err = %v ", err)
	})
}

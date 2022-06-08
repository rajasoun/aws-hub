package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStartCmd(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Server with Mock Start Function", func(t *testing.T) {
		cmdhandler := CmdHandler{}
		cmdhandler.EnableShutdDown = true
		err := cmdhandler.StartCommand(&cli.Context{})
		assert.NoError(err, "err = %v ", err)
	})
}

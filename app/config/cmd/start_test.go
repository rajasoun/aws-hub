package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartCmd(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Command With Custom Port Zero", func(t *testing.T) {
		context := NewContext()
		cmdhandler := Handler{EnableShutdDown: true}
		startCommand := New("start", "Start Server", cmdhandler.StartCommand)
		err := startCommand.Run(context)
		assert.NoError(err, "err = %v ", err)
	})
}

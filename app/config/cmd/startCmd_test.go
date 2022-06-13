package cmd

import (
	"flag"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStartCmd(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Command With Custom Port Zero", func(t *testing.T) {
		mockApp := &cli.App{Writer: ioutil.Discard}
		set := flag.NewFlagSet("test", 0)
		port, _ := test.GetFreePort("localhost:0")
		portString := strconv.Itoa(port)
		_ = set.Parse([]string{"start", "--port", portString})
		context := cli.NewContext(mockApp, set, nil)
		cmdhandler := CmdHandler{EnableShutdDown: true}
		startCommand := GetCommand(cmdhandler.StartCommand)
		err := startCommand.Run(context)
		assert.NoError(err, "err = %v ", err)
	})
}

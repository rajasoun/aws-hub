package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStartCmd(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Start Server with Mock Start Function", func(t *testing.T) {
		err := StartCommand(&cli.Context{})
		assert.Error(err, "err = %v ", err)
	})
}

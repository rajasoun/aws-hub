package cmd

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetCommand(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name            string
		wantCommandsLen int
		wantCommandName string
	}{
		{"Check start command", 1, "start"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdhandler := CmdHandler{}
			cmdhandler.EnableShutdDown = true
			cmds := GetCommand(cmdhandler.StartCommand)
			gotCommandName := cmds.Name
			assert.Containsf(tt.wantCommandName, gotCommandName,
				"setUpCommands() = %v , want = %v", gotCommandName, tt.wantCommandName)
		})
	}
}

func mockStartCommand(appCtx *cli.Context) error {
	log.Println("mockStartCommandHandler !!!")
	return nil
}

func TestCreateCommand(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name string
		cmd  string
		want string
	}{
		{"Check Start with Mock handler", "start", "mockStartCommandHandler"},
		{"Check InValid Command Handler", "dummy", CommandNotFoundMsg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{}
			app.Flags = flag.GetFlags()
			mockCmd := CreateCommand(mockStartCommand)
			commands := []*cli.Command{&mockCmd}
			app.Commands = commands
			app.CommandNotFound = ErrCommand()
			var bufferWriter bytes.Buffer
			args := os.Args[0:1]
			args = append(args, tt.cmd)
			app.Writer = &bufferWriter
			log.SetOutput(&bufferWriter)
			err := app.Run(args)
			assert.NoError(err, "mock start failed error = %v ", err)
			got := bufferWriter.String()
			contains := strings.Contains(got, tt.want)
			assert.True(contains, "got = %v, want = %v ", got, tt.want)
		})
	}
}

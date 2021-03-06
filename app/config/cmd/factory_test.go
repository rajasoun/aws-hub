package cmd

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"

	hubConfig "github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func MockStartCommand(appCtx *cli.Context) error {
	log.Println("mockStartCommandHandler !!!")
	return nil
}

func NewAppWithMockCommands(writer io.Writer) *cli.App {
	app := &cli.App{}
	app.Flags = hubConfig.GetFlags()
	mockCmd := New("start", "Mock Start Server", MockStartCommand)
	commands := []*cli.Command{&mockCmd}
	app.Commands = commands
	app.CommandNotFound = NewErr()
	app.Writer = writer
	log.SetOutput(writer)
	log.SetFlags(0)
	return app
}

func NewContext() *cli.Context {
	mockApp := &cli.App{Writer: ioutil.Discard}
	set := flag.NewFlagSet("test", 0)
	port, _ := test.GetFreePort("localhost:0")
	portString := strconv.Itoa(port)
	_ = set.Parse([]string{"start", "--port", portString})
	context := cli.NewContext(mockApp, set, nil)
	return context
}

func TestGetCommand(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check start command", func(t *testing.T) {
		cmdhandler := Handler{EnableShutdDown: true}
		cmds := New("start", "Start Server", cmdhandler.StartCommand)
		got := cmds.Name
		want := "start"
		assert.Containsf(want, got, "GetCommand() = %v , want = %v", got, want)
	})
}

func TestCreateCommand(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name string
		cmd  string
		want string
	}{
		{
			name: "Check Start with Mock handler",
			cmd:  "start",
			want: "mockStartCommandHandler",
		},
		{
			name: "Check InValid Command Handler",
			cmd:  "dummy",
			want: CommandNotFoundMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outputBuffer bytes.Buffer
			outputBuffer.Reset()
			app := NewAppWithMockCommands(&outputBuffer)
			args := os.Args[0:1]
			args = append(args, tt.cmd)
			err := app.Run(args)

			assert.NoError(err, "Err app.Run() = %v with %v", err, tt.cmd)
			got := outputBuffer.String()
			assert.Contains(got, tt.want, "got = %v, want = %v , args = %v", got, tt.want, tt.cmd)
		})
	}
}

func TestNewErrCommand(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	t.Run("Check ErrCommand With Invalid Input", func(t *testing.T) {
		mockApp := &cli.App{Writer: ioutil.Discard}
		cmdhandler := Handler{EnableShutdDown: false}
		startCommand := New("start", "Start Server", cmdhandler.StartCommand)
		commands := []*cli.Command{&startCommand}
		mockApp.Commands = commands
		mockApp.CommandNotFound = NewErr()

		err := mockApp.Run([]string{"invalidCommand"})
		assert.NoError(err, "err = %v ", err)
	})
}

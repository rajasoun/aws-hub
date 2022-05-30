package hub

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_setUpApp(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	app := NewApp()
	tests := []struct {
		name string
		key  string
		want string
	}{
		{"Check Name", "Name", "AWS Hub"},
		{"Check Description", "Usage", "AWS Cost Explorer"},
		{"Check Version", "Version", "0.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.setUpIdentity()
			appMap := app.StructToMap()
			got := appMap[tt.key]
			assert.Equal(got, tt.want, "setUp() = %v , want = %v", got, tt.want)
		})
	}
}

func Test_setUpFlags(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	app := NewApp()

	tests := []struct {
		name  string
		index int
		want  string
	}{
		{"Check For Flag port", 0, "--port"},
		{"Check For Flag duration", 1, "--duration"},
		{"Check For Flag redis", 2, "--redis"},
		{"Check For Flag dataset", 3, "--dataset"},
		{"Check For Flag multiple", 4, "--multiple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.setUpFlags()
			got := SliceToStrMap(app.cli.Flags)
			assert.Containsf(got[tt.index], tt.want, "setFlags(tt.app) = %v, want = %v", got, tt.want)
		})
	}
}

func Test_setUpCommands(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	app := NewApp()

	tests := []struct {
		name            string
		wantCommandsLen int
		wantCommandName string
	}{
		{"Check start command", 1, "start"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.setUpCommands()
			commands := app.cli.Commands
			gotCommandsLen := len(commands)
			assert.Equal(tt.wantCommandsLen, gotCommandsLen,
				"len(tt.app.Commands) = %v , want = %v", gotCommandsLen, tt.wantCommandsLen)
			gotCommandName := commands[0].Name
			assert.Containsf(tt.wantCommandName, gotCommandName,
				"setUpCommands() = %v , want = %v", gotCommandName, tt.wantCommandName)
		})
	}
}

func TestExecute(t *testing.T) {
	assert := assert.New(t)
	args := os.Args[0:1]
	t.Parallel()
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Check Starting App with wrong command",
			args: append(args, "dummy"),
			want: "Command not found \"dummy\" !",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Execute(tt.args, &buf)
			got := buf.String()
			assert.Equal(tt.want, got, "got = %v, want = %v ", got, tt.want)
		})
	}
}

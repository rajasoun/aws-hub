package hub

import (
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
			appMap := app.structToMap()
			got := appMap[tt.key]
			assert.Equal(got, tt.want, "setUp() = %v , want = %v", got, tt.want)
		})
	}
}

func Test_setFlags(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	app := NewApp()

	tests := []struct {
		name  string
		index int
		want  string
	}{
		{"Check For Flag port", 0, "--port"},
		{"Check For Flag port", 1, "--duration"},
		{"Check For Flag port", 2, "--redis"},
		{"Check For Flag port", 3, "--dataset"},
		{"Check For Flag port", 4, "--multiple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.setUpFlags()
			got := sliceToStrMap(app.cli.Flags)
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

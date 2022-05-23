package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_setUpCommands(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name            string
		app             *cli.App
		wantCommandsLen int
		wantCommandName string
	}{
		{"Check start command", &cli.App{}, 1, "start"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpCommands(tt.app)
			commands := tt.app.Commands
			gotCommandsLen := len(commands)
			assert.Equal(tt.wantCommandsLen, gotCommandsLen,
				"len(tt.app.Commands) = %v , want = %v", gotCommandsLen, tt.wantCommandsLen)
			gotCommandName := commands[0].Name
			assert.Containsf(tt.wantCommandName, gotCommandName,
				"setUpCommands() = %v , want = %v", gotCommandName, tt.wantCommandName)
		})
	}
}

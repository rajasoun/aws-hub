package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetCommands(t *testing.T) {
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
			cmds := GetCommands()
			gotCommandsLen := len(cmds)
			assert.Equal(tt.wantCommandsLen, gotCommandsLen,
				"len(cmd.GetCommands() = %v , want = %v", gotCommandsLen, tt.wantCommandsLen)
			gotCommandName := cmds[0].Name
			assert.Containsf(tt.wantCommandName, gotCommandName,
				"setUpCommands() = %v , want = %v", gotCommandName, tt.wantCommandName)
		})
	}
}

func TestGetErrCommand(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		want func(appCtx *cli.Context, command string)
	}{
		{
			name: "Check ErrCommand",
			want: func(appCtx *cli.Context, command string) {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reflect.TypeOf(GetErrCommand())
			want := reflect.TypeOf(tt.want)
			assert.Equal(want, got, "reflect.TypeOf(GetErrCommand() = %v , want = %v", got, want)
		})
	}
}

package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetErrCommand(t *testing.T) {
	t.Parallel()
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
			got := reflect.TypeOf(ErrCommand())
			want := reflect.TypeOf(tt.want)
			assert.Equal(want, got, "reflect.TypeOf(GetErrCommand() = %v , want = %v", got, want)
		})
	}
}

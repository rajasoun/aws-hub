package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_setUp(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		app  *cli.App
		want string
	}{
		{"Check for start command", New(), "start"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.app.Command("start").Name
			// tt.app.Command("start").Action(&cli.Context{})
			assert.Equal(tt.want, got, "setUp() = %v, want = %v", tt.want, got)
		})
	}
}

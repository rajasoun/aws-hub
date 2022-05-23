package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_setFlags(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name  string
		app   *cli.App
		index int
		want  string
	}{
		{"Check For Flag port", &cli.App{}, 0, "--port"},
		{"Check For Flag port", &cli.App{}, 1, "--duration"},
		{"Check For Flag port", &cli.App{}, 2, "--redis"},
		{"Check For Flag port", &cli.App{}, 3, "--dataset"},
		{"Check For Flag port", &cli.App{}, 4, "--multiple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFlags(tt.app)
			got := sliceToStrMap(getFlags())
			assert.Containsf(got[tt.index], tt.want, "setFlags(tt.app) = %v, want = %v", got, tt.want)
		})
	}
}

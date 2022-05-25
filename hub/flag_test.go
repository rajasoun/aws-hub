package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			setUpFlags(app.cli)
			got := sliceToStrMap(getFlags())
			assert.Containsf(got[tt.index], tt.want, "setFlags(tt.app) = %v, want = %v", got, tt.want)
		})
	}
}

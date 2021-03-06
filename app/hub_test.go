package app

import (
	"bytes"
	"os"
	"testing"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/rajasoun/aws-hub/app/config/flag"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	app := New()
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
			app.setUpInfo()
			app.setUpAuthors()
			appMap := app.StructToMap(app.cli)
			got := appMap[tt.key]
			assert.Equal(got, tt.want, "setUp() = %v , want = %v", got, tt.want)
		})
	}
}

func TestSetUpFlags(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	app := New()

	tests := []struct {
		name  string
		index int
		want  string
	}{
		{"Check For Flag port", 0, "--port"},
		{"Check For Flag duration", 1, "--duration"},
		{"Check For Flag redis", 2, "--cache"},
		{"Check For Flag multiple", 3, "--multiple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.cli.Flags = flag.GetFlags()
			got := app.SliceToStrMap(app.cli.Flags)
			assert.Containsf(got[tt.index], tt.want, "setFlags(tt.app) = %v, want = %v", got, tt.want)
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
			want: cmd.CommandNotFoundMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := Execute(tt.args, &buf)
			assert.NoError(err, "Execute err = %v", err)
			got := buf.String()
			assert.Contains(got, tt.want, "got = %v, want = %v ", got, tt.want)
		})
	}
}

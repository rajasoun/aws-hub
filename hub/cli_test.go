package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_setUpApp(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name string
		app  *cli.App
		key  string
		want string
	}{
		{"Check Name", &cli.App{}, "Name", "AWS Hub"},
		{"Check Description", &cli.App{}, "Usage", "AWS Cost Explorer"},
		{"Check Version", &cli.App{}, "Version", "0.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUpApp(tt.app)
			appMap := StructToMap(tt.app)
			got := appMap[tt.key]
			assert.Equal(got, tt.want, "setUp() = %v , want = %v", got, tt.want)
		})
	}
}

func Test_parseArgs(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name         string
		appCtx       *cli.Context
		wantPort     int
		wantMultiple bool
	}{
		{"Check Args port and multiple", &cli.Context{}, 3000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPort, _, gotMultiple := parseArgs(tt.appCtx)
			assert.Equal(gotPort, tt.wantPort, "parseArgs() = %v , want = %v", gotPort, tt.wantPort)
			assert.Equal(gotMultiple, tt.wantMultiple, "parseArgs() = %v , want = %v", gotMultiple, tt.wantMultiple)
		})
	}
}

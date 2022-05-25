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
			app.setUpApp()
			appMap := app.structToMap()
			got := appMap[tt.key]
			assert.Equal(got, tt.want, "setUp() = %v , want = %v", got, tt.want)
		})
	}
}

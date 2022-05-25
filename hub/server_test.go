package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewServer(t *testing.T) {

	assert := assert.New(t)
	t.Parallel()
	cliContext := NewCliContext(&cli.Context{})
	server := NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	tests := []struct {
		name string
		want string
	}{
		{"Check NewServer ", "Mux Server 0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := server.name
			assert.Equal(got, tt.want, "server.name = %v, want %v", got, tt.want)
		})
	}
}

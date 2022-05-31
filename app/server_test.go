package app

import (
	"testing"

	"github.com/rajasoun/aws-hub/app/args"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewServer(t *testing.T) {

	assert := assert.New(t)
	t.Parallel()
	cliContext := args.NewCliContext(&cli.Context{})
	server, _ := NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	tests := []struct {
		name string
		want string
	}{
		{"Check NewServer Name", "Mux Server 0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := server.name
			assert.Equal(got, tt.want, "server.name = %v, want %v", got, tt.want)
			gotHandler := server.GetAWSHandler()
			assert.NotNil(gotHandler, "awsHandler = %v", gotHandler)
			err := server.Start(999999999)
			assert.Error(err, "Invalid Port err = %v ", err)
		})
	}
}

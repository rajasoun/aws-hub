package server

import (
	"context"
	"testing"
	"time"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewServer(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	cliContext := arg.NewCliContext(&cli.Context{})
	server, _ := NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
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
			err := server.Start(999999999, false)
			assert.Error(err, "Invalid Port err = %v ", err)
		})
	}
}

func TestHTTPServerStart(t *testing.T) {
	cliContext := arg.NewCliContext(&cli.Context{})
	server, _ := NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
	srv := server.NewHTTPServer(":45566")
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()
	err := srv.StartHTTPServer()
	if err != nil {
		t.Error("unexpected error:", err)
	}
}

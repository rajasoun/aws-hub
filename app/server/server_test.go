package server

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	cliContext := arg.NewDefaultCliContext()
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
			server.shutdownDuration = 0
			serverErr := server.Start(0, true)
			assert.NoError(serverErr, "server start on port 0 = %v ", err)
		})
	}
}

func TestHTTPServerStart(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	cliContext := arg.NewDefaultCliContext()
	server, _ := NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
	srv := server.NewHTTPServer(":45566")
	go func() {
		time.Sleep(0 * time.Second)
		err := srv.Shutdown(context.Background())
		assert.NoError(err, "srv.Shutdown() = %v", err)
	}()
	err := srv.StartHTTPServer()
	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestHandleShutdown(t *testing.T) {
	//assert := assert.New(t)
	t.Parallel()
	t.Run("Check Handle Server Shutdown", func(t *testing.T) {
		cliContext := arg.NewDefaultCliContext()
		server, _ := NewServer(cliContext.GetCache(), cliContext.GetAwsProfileType())
		httpServer := server.NewHTTPServer(":0")
		server.RegisterShutdown(httpServer)
		err := httpServer.StartHTTPServer()
		if err != nil {
			log.Printf("err starting http server = %v", err)
		}
		defer httpServer.Close()
	})
}

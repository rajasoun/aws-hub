package test

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTestRun(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check Is Test Run", func(t *testing.T) {
		want := true
		got := IsTestRun()
		assert.Equal(want, got, "IsTestRun() = %v, want %v", got, want)

	})
	t.Run("Check GetFreePort with Valid Address", func(t *testing.T) {
		_, err := GetFreePort("localhost:0")
		assert.NoError(err, "Err = %v", err)
	})
	t.Run("Check GetFreePort with InValid Address", func(t *testing.T) {
		_, err := GetFreePort("Invalid:Invalid")
		assert.Error(err, "Err = %v", err)
	})
}

func PingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode("{Ok}")
}

func TestDoSimulation(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := MockServer{}
	t.Run("Check DoSimulation For Ping Handler", func(t *testing.T) {
		responseRecorder := mock.DoSimulation(PingHandler, nil)
		got := responseRecorder.Code
		assert.Equal(http.StatusOK, got, "PingHandler() = %v, want = %v", got, http.StatusOK)
	})
	t.Run("Check DoSimulation For Success Handler", func(t *testing.T) {
		responseRecorder := mock.DoSimulation(MockSuccessHandler, nil)
		got := responseRecorder.Code
		want := http.StatusOK
		assert.Equal(want, got, "MockSuccessHandler() = %v, want = %v", got, want)
	})
	t.Run("Check DoSimulation For Failure Handler", func(t *testing.T) {
		responseRecorder := mock.DoSimulation(MockFailureHandler, nil)
		got := responseRecorder.Code
		want := http.StatusInternalServerError
		assert.Equal(want, got, "MockFailureHandler() = %v, want = %v", got, want)
	})
}

func TestHandleErr(t *testing.T) {
	t.Parallel()
	type args struct {
		err    error
		errMsg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Check HandleErr with nil",
			args: args{
				err:    nil,
				errMsg: "",
			},
		},
		{
			name: "Check HandleErr with err",
			args: args{
				err:    errors.New("test error"),
				errMsg: "TestErr",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleErr(tt.args.err, tt.args.errMsg)
		})
	}
}

func TestCheckAddressAvailable(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		addr *net.TCPAddr
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Check Address Available",
			args: args{
				addr: &net.TCPAddr{
					Port: 44519,
				},
			},
			wantErr: false,
		},
		{
			name: "Check Address Available",
			args: args{
				addr: &net.TCPAddr{
					Port: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckAddressAvailable(tt.args.addr)
			assert.NoError(err, "CheckAddressAvailable() = %v", err)
			if (err != nil) != tt.wantErr {
				assert.Error(err, "CheckAddressAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package test

import (
	"encoding/json"
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
	t.Run("Check ExecuteHandler", func(t *testing.T) {
		mock := MockServer{}
		responseRecorder := mock.DoSimulation(PingHandler, nil)
		got := responseRecorder.Code
		assert.Equal(http.StatusOK, got, "got = %v, want = %v", got, http.StatusOK)
	})
}

func PingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode("{Ok}")
}

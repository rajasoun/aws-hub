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

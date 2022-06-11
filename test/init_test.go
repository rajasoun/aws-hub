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
	t.Run("Check GetFreePort", func(t *testing.T) {
		_, err := GetFreePort()
		assert.NoError(err, "Err = %v", err)
	})
	t.Run("Check GetFreePort", func(t *testing.T) {
		responseRecorder := ExecuteHandler(PingHandler, map[string]string{})
		got := responseRecorder.Code
		assert.Equal(http.StatusOK, got, "got = %v, want = %v", got, http.StatusOK)
	})
}

func PingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode("{Ok}")
}

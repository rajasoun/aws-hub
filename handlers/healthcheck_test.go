package handlers

import (
	"net/http"
	"testing"

	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	handler := NewDefaultAWSHandler(false)
	mock := test.MockServer{}

	t.Run("Check ", func(t *testing.T) {
		responseWriter := mock.DoSimulation(handler.HealthCheckHandler, nil)
		handler.ListProfilesHandler(responseWriter, nil)
		got := responseWriter.Code
		assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
	})
}

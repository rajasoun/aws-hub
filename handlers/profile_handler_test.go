package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
)

func TestAWSHandlerGetSections(t *testing.T) {
	t.Run("Check Get Sections for Multiple Profile", func(t *testing.T) {
		assert := assert.New(t)
		handler := NewDefaultAWSHandler(true)
		responseWriter := test.ExecuteHandler(PingHandler, nil)
		sections := handler.GetSections(responseWriter)
		assert.GreaterOrEqual(len(sections), 0, "GetSections() = %v, want >= %v", len(sections), 0)
		got := responseWriter.Code
		assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
	})
}

func PingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	payload := map[string]string{"Status": "Ok"}
	jsonPayLoad, _ := json.Marshal(payload)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(jsonPayLoad)
}

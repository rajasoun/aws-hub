package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
)

func TestAWSHandlerGetSections(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := test.MockServer{}
	tests := []struct {
		name           string
		credentialFile string
		isMultiple     bool
	}{
		{
			name:           "Check Get Sections for Multiple Profile",
			credentialFile: config.DefaultSharedCredentialsFilename(),
			isMultiple:     true,
		},
		{
			name:           "Check Get Sections for Multiple Profile with Invalid File",
			credentialFile: "InvalidFile",
			isMultiple:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(tt.isMultiple)
			responseWriter := mock.DoSimulation(PingHandler, nil)
			credentialFile := tt.credentialFile
			sections := handler.GetSections(responseWriter, credentialFile)
			assert.GreaterOrEqual(len(sections), 0, "GetSections() = %v, want >= %v", len(sections), 0)
			got := responseWriter.Code
			assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
			handler.ListProfilesHandler(responseWriter, nil)
			got = responseWriter.Code
			assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
		})
	}

}

func PingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	payload := map[string]string{"Status": "Ok"}
	jsonPayLoad, _ := json.Marshal(payload)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(jsonPayLoad)
}

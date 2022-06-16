package handlers

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
)

func TestAWSListProfilesHandler(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := test.MockServer{}
	tests := []struct {
		name           string
		credentialFile string
		isMultiple     bool
		sections       int
		wantErr        bool
	}{
		{
			name:           "Check Get Sections for Multiple Profile",
			credentialFile: config.DefaultSharedCredentialsFilename(),
			isMultiple:     true,
			sections:       0,
			wantErr:        false,
		},
		{
			name:           "Check Get Sections for Multiple Profile with Invalid File",
			credentialFile: "InvalidFile",
			isMultiple:     false,
			sections:       0,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(tt.isMultiple)
			responseWriter := mock.DoSimulation(handler.HealthCheckHandler, nil)
			handler.ListProfilesHandler(responseWriter, nil)
			got := responseWriter.Code
			assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
			if tt.wantErr {
				credentialFile := tt.credentialFile
				sections := praseSections(responseWriter, credentialFile)
				assert.GreaterOrEqual(len(sections), tt.sections, "GetSections() = %v, want >= %v", len(sections), 0)
				got := responseWriter.Code
				assert.Equal(http.StatusOK, got, "Status = %v , want = %v", got, http.StatusOK)
			}
		})
	}

}

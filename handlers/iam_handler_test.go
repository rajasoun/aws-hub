package handlers

import (
	"net/http"
	"testing"

	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIAMHandlerWithPing(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mockServer := test.MockServer{}
	handler := NewDefaultAWSHandler(false)

	tests := []struct {
		name        string
		muxVars     map[string]string
		handlerFunc func(w http.ResponseWriter, r *http.Request)
		want        int
	}{
		{
			name:        "Check handler.IAMGetUserCountHandler",
			muxVars:     map[string]string{"ApiName": "DoPing"},
			handlerFunc: handler.IAMGetUserCountHandler,
			want:        http.StatusOK,
		},
		{
			name:        "Check  handler.IAMGetUserIdentityHandler",
			muxVars:     map[string]string{"ApiName": "DoPing"},
			handlerFunc: handler.IAMGetUserIdentityHandler,
			want:        http.StatusOK,
		},
		{
			name:        "Check  handler.IAMGetAliasesHandler",
			muxVars:     map[string]string{"ApiName": "DoPing"},
			handlerFunc: handler.IAMGetAliasesHandler,
			want:        http.StatusOK,
		},
		{
			name:        "Check  handler.IAMGetAliasesHandler",
			muxVars:     map[string]string{"None": "None"},
			handlerFunc: handler.IAMGetAliasesHandler,
			want:        http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := mockServer.DoSimulation(tt.handlerFunc, tt.muxVars)
			got := responseRecorder.Code
			assert.Equal(tt.want, got, "got = %v, want = %v", got, tt.want)

			awsWrapper := MockAWSWrapper(test.MockSuccessHandler)
			expectedOutput := &MockOutput{Message: "Test with Success"}
			client := new(MockAwsAPI)
			client.On("Execute", mock.Anything).Return(expectedOutput, nil)
			awsWrapper.InvokeAPI(client, "dummy", "dummy")
			// Check Cache
			awsWrapper.InvokeAPI(client, "dummy", "dummy")
		})
	}
}

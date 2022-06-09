package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIAMHandler(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	assert := assert.New(t)
	t.Parallel()
	var request *http.Request
	var response *httptest.ResponseRecorder

	handler := NewDefaultAWSHandler(false)
	router := handler.SetUpRoutes()

	tests := []struct {
		name     string
		endPoint string
		wantKey  string
	}{
		{
			name:     "UserCount API /aws/iam/users",
			endPoint: "/aws/iam/users",
		},
		{
			name:     "User Identity API /aws/iam/account",
			endPoint: "/aws/iam/account",
		},
		{
			name:     "Account Alias API /aws/iam/alias",
			endPoint: "/aws/iam/alias",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ = http.NewRequest("GET", tt.endPoint, nil)
			response = executeRequest(router, request)
			assert.Equal(http.StatusOK, response.Code, "OK response is expected")
		})
	}
}

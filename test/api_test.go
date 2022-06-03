package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/handlers"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/rajasoun/aws-hub/service/cache"
)

func TestAPI(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()

	_, router := server.NewServer(&cache.Memory{}, false)
	server := httptest.NewServer(handlers.LoggingHandler(os.Stdout, router))
	defer server.Close()
	expect := httpexpect.New(t, server.URL)

	tests := []struct {
		name     string
		endPoint string
		wantKey  string
	}{
		{
			name:     "HealthCheck API /health",
			endPoint: "/health",
			wantKey:  "http-server-alive",
		},
		{
			name:     "Profiles API /aws/profiles",
			endPoint: "/aws/profiles",
			wantKey:  "multiple",
		},
		{
			name:     "UserCount API /aws/iam/users",
			endPoint: "/aws/iam/users",
			wantKey:  "usercount",
		},
		{
			name:     "User Identity API /aws/iam/account",
			endPoint: "/aws/iam/account",
			wantKey:  "username",
		},
		{
			name:     "Account Alias API /aws/iam/alias",
			endPoint: "/aws/iam/alias",
			wantKey:  "list",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect.GET(tt.endPoint).
				Expect().
				Status(http.StatusOK).
				JSON().Object().ContainsKey(tt.wantKey).
				NotEmpty()
		})
	}
}

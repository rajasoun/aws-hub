package aws

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/handlers"
)

func TestIAMHandlers(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	tests := []struct {
		name               string
		endPoint           string
		hasMultipleProfile bool
		wantKey            string
	}{
		{"Check IAMGetAliasesHandler", "/aws/iam/alias", false, "list"},
		{"Check IAMGetUserIdentityHandler", "/aws/iam/account", false, "username"},
		{"Check IAMGetUserCountHandler", "/aws/iam/users", false, "usercount"},
		{"Check ConfigProfilesHandler", "/aws/profiles", false, "multiple"},
		{"Check ConfigProfilesHandler", "/aws/profiles", true, "multiple"},
		{"Check HealthCheckHandler", "/health", false, "http-server-alive"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(tt.hasMultipleProfile)
			router := handler.SetUpRoutes()
			server := httptest.NewServer(handlers.LoggingHandler(os.Stdout, router))
			defer server.Close()

			expect := httpexpect.New(t, server.URL)

			expect.GET(tt.endPoint).
				Expect().
				Status(http.StatusOK).
				JSON().Object().ContainsKey(tt.wantKey)
		})
	}
}

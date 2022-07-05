package e2e_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/handlers"
	"github.com/rajasoun/aws-hub/app/server"
	routes "github.com/rajasoun/aws-hub/handlers"
	"github.com/rajasoun/aws-hub/service/cache"
)

func TestAPI(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()

	_, router := server.NewServer(&cache.Memory{}, false)
	server := httptest.NewServer(handlers.LoggingHandler(os.Stdout, router))
	expect := httpexpect.New(t, server.URL)
	defer server.Close()

	tests := []struct {
		name     string
		endPoint string
		wantKey  string
	}{
		{
			name:     "HealthCheck API /health",
			endPoint: routes.HealthEndPoint,
			wantKey:  "http-server-alive",
		},
		{
			name:     "Profiles API /aws/profiles",
			endPoint: routes.LocalProfilesEndPoint,
			wantKey:  "multiple",
		},
		{
			name:     "UserCount API /aws/iam/users",
			endPoint: routes.UsersCountEndPoint,
			wantKey:  "usercount",
		},
		{
			name:     "User Identity API /aws/iam/account",
			endPoint: routes.UserIdentityEndPoint,
			wantKey:  "username",
		},
		{
			name:     "Account Alias API /aws/iam/alias",
			endPoint: routes.AccountAliasEndPoint,
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

// func TestAWSAPI(t *testing.T) {
// 	assert := assert.New(t)
// 	cfg, _ := credential.New().LoadDefaultConfig()
// 	client := iam.NewFromConfig(cfg)

// 	response, apiErr := service.GetAliases(client)
// 	assert.NoError(apiErr)
// 	assert.Equal("secops-experiments", response.List[0])
// }

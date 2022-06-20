package e2e_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/rajasoun/aws-hub/service/cache"
	"github.com/steinfletcher/apitest"

	routes "github.com/rajasoun/aws-hub/handlers"
)

func TestGenerateSequenceDiagram(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()

	_, router := server.NewServer(&cache.Memory{}, false)

	tests := []struct {
		name     string
		endPoint string
	}{
		{"HealthCheck API", routes.HealthEndPoint},
		{"AWS Profiles API", routes.LocalProfilesEndPoint},
		{"User Identity API", routes.UserIdentityEndPoint},
		{"Account Alias API", routes.AccountAliasEndPoint},
		{"User Count API", routes.UsersCountEndPoint},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateSequenceDiagram(t, router, tt.endPoint)
		})
	}

}

func GenerateSequenceDiagram(t *testing.T, router *mux.Router, endPoint string) {
	apitest.New("Check End Point -> " + endPoint).
		Report(apitest.SequenceDiagram()).
		Handler(router).
		Get(endPoint).
		Expect(t).
		Status(http.StatusOK).
		End()
}

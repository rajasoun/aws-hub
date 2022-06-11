package api_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rajasoun/aws-hub/app/server"
	"github.com/rajasoun/aws-hub/service/cache"
	"github.com/steinfletcher/apitest"
)

func TestGenerateSequenceDiagram(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()

	_, router := server.NewServer(&cache.Memory{}, false)

	t.Run("HealthCheck API /health", func(t *testing.T) {
		GenerateSequenceDiagram(t, router, "/health")
	})

	t.Run("AWS Profiles API /aws/profile", func(t *testing.T) {
		GenerateSequenceDiagram(t, router, "/aws/profiles")
	})

	t.Run("User Identity API /aws/iam/account", func(t *testing.T) {
		GenerateSequenceDiagram(t, router, "/aws/iam/account")
	})

	t.Run("User Identity API /aws/iam/users", func(t *testing.T) {
		GenerateSequenceDiagram(t, router, "/aws/iam/users")
	})

	t.Run("Account Alias API /aws/iam/alias", func(t *testing.T) {
		GenerateSequenceDiagram(t, router, "/aws/iam/alias")
	})
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

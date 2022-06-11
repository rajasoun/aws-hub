package api_test

import (
	"net/http"
	"os"
	"testing"

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
		apitest.New("Health Check API  /health").
			Report(apitest.SequenceDiagram()).
			Meta(map[string]interface{}{"host": "HealthCheckHandler"}).
			Handler(router).
			Get("/health").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
	t.Run("User Identity API /aws/iam/account", func(t *testing.T) {
		apitest.New("User Identity API /aws/iam/account").
			Report(apitest.SequenceDiagram()).
			Meta(map[string]interface{}{"host": "IAMUserHandler"}).
			Handler(router).
			Get("/aws/iam/account").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
	t.Run("User Identity API /aws/iam/users", func(t *testing.T) {
		apitest.New("User Count API /aws/iam/users").
			Report(apitest.SequenceDiagram()).
			Meta(map[string]interface{}{"host": "IAMUserCountHandler"}).
			Handler(router).
			Get("/aws/iam/users").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
	t.Run("Account Alias API /aws/iam/alias", func(t *testing.T) {
		apitest.New("User Count API /aws/iam/alias").
			Report(apitest.SequenceDiagram()).
			Meta(map[string]interface{}{"host": "IAMAccountAliasHandler"}).
			Handler(router).
			Get("/aws/iam/alias").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

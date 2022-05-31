package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/rajasoun/aws-hub/app"
	"github.com/rajasoun/aws-hub/services/cache"
	"github.com/steinfletcher/apitest"
)

func Test_Flow(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()

	_, router := app.NewServer(&cache.Memory{}, false)

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

}

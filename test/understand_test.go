package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/rajasoun/aws-hub/hub"
	"github.com/steinfletcher/apitest"
	"github.com/urfave/cli/v2"
)

func Test_Flow(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()
	cliContext := hub.NewCliContext(&cli.Context{})
	server := hub.NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	awsHandler := server.GetAWSHandler()
	router := awsHandler.SetUpRoutes()

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

package test

import (
	"net/http"
	"testing"

	"github.com/rajasoun/aws-hub/handlers/aws"
	"github.com/rajasoun/aws-hub/hub"
	"github.com/steinfletcher/apitest"
	"github.com/urfave/cli/v2"

	"github.com/gorilla/mux"
)

func HealthCheckInjector(awsHandler *aws.AWSHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		awsHandler.HealthCheckHandler(w, r)
	}
}

func IAMUserInjector(awsHandler *aws.AWSHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		awsHandler.IAMUserHandler(w, r)
	}
}

func newRouter() *mux.Router {
	cliContext := hub.NewCliContext(&cli.Context{})
	server := hub.NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	router := mux.NewRouter()
	awsHandler := server.GetAWSHandler()
	router.HandleFunc("/health", HealthCheckInjector(awsHandler)).Methods("GET")
	router.HandleFunc("/aws/iam/account", IAMUserInjector(awsHandler)).Methods("GET")
	return router
}

func Test_Flow(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	t.Skip("Skipping INTEGRATION Tests")
	t.Parallel()
	router := newRouter()
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

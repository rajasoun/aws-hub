package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/rajasoun/aws-hub/hub"
	"github.com/urfave/cli/v2"
)

func TestAPI_All(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	t.Skip("Skipping INTEGRATION Tests")
	t.Parallel()
	cliContext := hub.NewCliContext(&cli.Context{})
	server := hub.NewServer(cliContext.Cache(), cliContext.IsMultipleAwsProfiles())
	awsHandler := server.GetAWSHandler()
	mux := http.NewServeMux()

	t.Run("HealthCheck API /health", func(t *testing.T) {
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.HealthCheckHandler(w, r)
		})
		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/health").
			Expect().
			Status(http.StatusOK).
			JSON().Object().ContainsKey("http-server-alive").
			ValueEqual("http-server-alive", "Ok")
	})

	t.Run("Profiles API /aws/profiles", func(t *testing.T) {
		mux.HandleFunc("/aws/profiles", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.ConfigProfilesHandler(w, r)
		})
		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/aws/profiles").
			Expect().
			Status(http.StatusOK).
			JSON().Array().Empty()
	})

	t.Run("UserCount API /aws/iam/users", func(t *testing.T) {
		mux.HandleFunc("/aws/iam/users", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.IAMListUsersHandler(w, r)
		})
		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/aws/iam/users").
			Expect().
			Status(http.StatusOK).
			JSON().Object().ContainsKey("usercount").
			NotEmpty()
	})
	t.Run("User Identity API /aws/iam/account", func(t *testing.T) {
		mux.HandleFunc("/aws/iam/account", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.IAMUserHandler(w, r)
		})
		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/aws/iam/account").
			Expect().
			Status(http.StatusOK).
			JSON().Object().ContainsKey("username").
			NotEmpty()
	})
}

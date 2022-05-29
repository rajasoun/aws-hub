package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/rajasoun/aws-hub/hub"
	"github.com/urfave/cli/v2"
)

func Test_API(t *testing.T) {
	//ToDo: Secure way of Passing Credentials
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
	t.Parallel()
	hubCliCtx := hub.NewCliContext(&cli.Context{})
	hub := hub.NewServer(hubCliCtx.Cache(), hubCliCtx.IsMultipleAwsProfiles())
	awsHandler := hub.GetAWSHandler()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("HealthCheck API /health", func(t *testing.T) {
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.HealthCheckHandler(w, r)
		})
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
			awsHandler.IAMGetUserCountHandler(w, r)
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
			awsHandler.IAMGetUserIdentityHandler(w, r)
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
	t.Run("Account Alias API /aws/iam/alias", func(t *testing.T) {
		mux.HandleFunc("/aws/iam/alias", func(w http.ResponseWriter, r *http.Request) {
			awsHandler.IAMGetAliasesHandler(w, r)
		})
		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/aws/iam/alias").
			Expect().
			Status(http.StatusOK).
			JSON().Object().ContainsKey("list").
			NotEmpty()
	})
}

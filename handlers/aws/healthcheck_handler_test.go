package aws

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAWSHandler_HealthCheckHandler(t *testing.T) {
	t.Parallel()
	t.Run("Check Health API /health", func(t *testing.T) {
		mux := http.NewServeMux()
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler := http.HandlerFunc(awsHandler.HealthCheckHandler)
		// mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// 	awsHandler := NewHandler()
		// 	awsHandler.HealthCheckHandler(w, r)
		// }

		server := httptest.NewServer(mux)
		defer server.Close()
		expect := httpexpect.New(t, server.URL)

		expect.GET("/health").
			Expect().
			Status(http.StatusOK).
			JSON().Object().ContainsKey("http-server-alive").
			ValueEqual("http-server-alive", "Ok")
	})
}

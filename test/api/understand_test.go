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

	tests := []struct {
		name     string
		endPoint string
	}{
		{"HealthCheck API", "/health"},
		{"AWS Profiles API", "/aws/profiles"},
		{"User Identity API", "/aws/iam/account"},
		{"Account Alias API", "/aws/iam/alias"},
		{"User Count API", "/aws/iam/users"},
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

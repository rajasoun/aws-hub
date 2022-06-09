package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func GetRouteMap(r *mux.Router) map[string]string {
	routes := make(map[string]string)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			log.Printf("err = %v", err)
			return nil
		}
		routes[path] = path
		return nil
	})
	log.Printf("Routes%s\n", routes)
	return routes
}

func TestAPIRoutes(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name     string
		endPoint string
	}{
		{"Check IAMGetAliasesHandler", "/aws/iam/alias"},
		{"Check IAMGetUserIdentityHandler", "/aws/iam/account"},
		{"Check IAMGetUserCountHandler", "/aws/iam/users"},
		{"Check ConfigProfilesHandler", "/aws/profiles"},
		{"Check ConfigProfilesHandler", "/aws/profiles"},
		{"Check HealthCheckHandler", "/health"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(false)
			router := handler.SetUpRoutes()
			got := GetRouteMap(router)[tt.endPoint]
			assert.Equal(tt.endPoint, got, "got = %v, want = %v", got, tt.endPoint)
		})
	}
}

func TestRoutes(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	var request *http.Request
	var response *httptest.ResponseRecorder

	handler := NewDefaultAWSHandler(false)
	router := handler.SetUpRoutes()

	tests := []struct {
		name     string
		endPoint string
	}{
		{"Check Health Check Handler", "/health"},
		{"Check Profile Handler", "/aws/profiles"},
	}
	//Note Roues Making AWS Calls are not included as they are not unit Tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ = http.NewRequest("GET", tt.endPoint, nil)
			response = executeRequest(router, request)
			assert.Equal(http.StatusOK, response.Code, "OK response is expected")
		})
	}
}

// IssueTestRequest executes an HTTP request described by rt, to a
// specified REST router.  It returns the HTTP response to the request.
func executeRequest(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

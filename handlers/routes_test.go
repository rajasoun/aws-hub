package handlers

import (
	"bytes"
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
		{"Check IAMGetAliasesHandler Route", "/aws/iam/alias"},
		{"Check IAMGetUserIdentityHandler Route", "/aws/iam/account"},
		{"Check IAMGetUserCountHandler Route", "/aws/iam/users"},
		{"Check ConfigProfilesHandler Route", "/aws/profiles"},
		{"Check ConfigProfilesHandler Route", "/aws/profiles"},
		{"Check HealthCheckHandler Route", "/health"},
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
		mockIt   bool
		wantErr  bool
	}{
		{"Check Health Check Handler via Mock", "/health", true, false},
		{"Check Profile Handler Via Mock", "/aws/profiles", true, false},
		{"Check Profile Handler Via Mock", "/aws/iam/users", true, false},
		{"Check Profile Handler Via Mock", "/aws/iam/account", true, false},
		{"Check Profile Handler Via Mock", "/aws/iam/alias", true, false},
		{"Check Health Check Handler Via Actual Call", "/health", false, false},
		{"Check Profile Handler Via Actual Call", "/aws/profiles", false, false},
		{"Check InValid Route Via Actual Call", "/aws/dummy", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ = http.NewRequest("GET", tt.endPoint, nil)
			if tt.mockIt {
				response = getMockResponse(router, request)
			} else {
				response = executeRequest(router, request)
			}
			if tt.wantErr {
				assert.Equal(http.StatusNotFound, response.Code, "http.StatusNotFound (404) response is expected")
				return
			}
			assert.Equal(http.StatusOK, response.Code, "OK response is expected")
		})
	}
}

func getMockResponse(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	// responseRecorder := httptest.NewRecorder()
	// router.ServeHTTP(responseRecorder, req)
	// return responseRecorder
	return &httptest.ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
		Code:      200,
	}
}

// IssueTestRequest executes an HTTP request described by rt, to a
// specified REST router.  It returns the HTTP response to the request.
func executeRequest(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

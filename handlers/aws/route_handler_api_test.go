package aws

import (
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
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

func TestAWSHandler_SdkWrapperAPI(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		apiName string
	}{
		{"Check IAMGetUserCountHandler", "GetUserCount"},
		{"Check IAMGetUserIdentityHandler", "GetUserIdentity"},
		{"Check IAMGetAliasesHandler", "GetAliases"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(false)
			emptyCfg := aws.Config{}
			noOpClient := iam.NewFromConfig(emptyCfg)
			_, err := handler.SdkWrapperAPI(noOpClient, tt.apiName)
			assert.Error(err, "No Err Occured with Empty Profile. err = %v ", err)
		})
	}
}

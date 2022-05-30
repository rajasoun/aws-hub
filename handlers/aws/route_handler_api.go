package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/gorilla/mux"

	hubIAM "github.com/rajasoun/aws-hub/services/aws/iam"
)

func (handler *AWSHandler) SetUpRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/aws/profiles", handler.ConfigProfilesHandler)
	router.HandleFunc("/aws/iam/users", handler.IAMGetUserCountHandler)
	router.HandleFunc("/aws/iam/account", handler.IAMGetUserIdentityHandler)
	router.HandleFunc("/aws/iam/alias", handler.IAMGetAliasesHandler)
	router.HandleFunc("/aws/cost/current", handler.CurrentCostHandler)
	router.HandleFunc("/aws/cost/history", handler.CostAndUsageHandler)
	router.HandleFunc("/aws/cost/forecast", handler.DescribeForecastPriceHandler)
	router.HandleFunc("/aws/cost/instance_type", handler.CostAndUsagePerInstanceTypeHandler)
	router.HandleFunc("/health", handler.HealthCheckHandler)
	return router
}

// ToDo Technical Debt - Use Interface to call the right method
// Use Dependency Injection
func (handler *AWSHandler) ExternalServiceGateway(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	var client *iam.Client = iam.NewFromConfig(cfg)
	switch {
	case apiName == "GetUserCount":
		response, err = hubIAM.GetUserCount(client)
	case apiName == "GetUserIdentity":
		response, err = hubIAM.GetUserIdentity(client)
	case apiName == "GetAliases":
		response, err = hubIAM.GetAliases(client)
	}
	return response, err
}

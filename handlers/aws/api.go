package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gorilla/mux"
)

func (handler *AWSHandler) AwsAPI(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "IAMListUsers":
		response, err = handler.aws.IAMListUsers(cfg)
	case apiName == "IAMUser":
		response, err = handler.aws.IAMUser(cfg)
		// case apiName == "DescribeCostAndUsage":
		// 	response, err = handler.aws.DescribeCostAndUsage(cfg)
		// case apiName == "DescribeCostAndUsagePerInstanceType":
		// 	response, err = handler.aws.DescribeCostAndUsagePerInstanceType(cfg)
		// case apiName == "DescribeForecastPrice":
		// 	response, err = handler.aws.DescribeForecastPrice(cfg)
	}
	return response, err
}

func (handler *AWSHandler) SetUpRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/aws/profiles", handler.ConfigProfilesHandler)
	router.HandleFunc("/aws/iam/users", handler.IAMListUsersHandler)
	router.HandleFunc("/aws/iam/account", handler.IAMUserHandler)
	router.HandleFunc("/aws/cost/current", handler.CurrentCostHandler)
	router.HandleFunc("/aws/cost/history", handler.CostAndUsageHandler)
	router.HandleFunc("/aws/cost/forecast", handler.DescribeForecastPriceHandler)
	router.HandleFunc("/aws/cost/instance_type", handler.CostAndUsagePerInstanceTypeHandler)
	router.HandleFunc("/health", handler.HealthCheckHandler)
	return router
}

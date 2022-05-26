package aws

import (
	"github.com/gorilla/mux"
)

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

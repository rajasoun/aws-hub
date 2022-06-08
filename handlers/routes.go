package handlers

import (
	"github.com/gorilla/mux"
)

func (handler *AWSHandler) SetUpRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.HealthCheckHandler)
	router.HandleFunc("/aws/profiles", handler.ListProfilesHandler)
	router.HandleFunc("/aws/iam/users", handler.IAMGetUserCountHandler)
	router.HandleFunc("/aws/iam/account", handler.IAMGetUserIdentityHandler)
	router.HandleFunc("/aws/iam/alias", handler.IAMGetAliasesHandler)
	// router.HandleFunc("/aws/cost/current", handler.CurrentCostHandler)
	// router.HandleFunc("/aws/cost/history", handler.CostAndUsageHandler)
	// router.HandleFunc("/aws/cost/forecast", handler.DescribeForecastPriceHandler)
	// router.HandleFunc("/aws/cost/instance_type", handler.CostAndUsagePerInstanceTypeHandler)
	return router
}

package aws

import (
	"net/http"
)

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.ce.total"
	apiToBeInvoked := "DescribeCostAndUsage"
	onErrMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.ce.history"
	apiToBeInvoked := "DescribeCostAndUsage"
	onErrMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) CostAndUsagePerInstanceTypeHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.ce.instance_type"
	apiToBeInvoked := "DescribeCostAndUsagePerInstanceType"
	onErrMsg := "ce:CostAndUsagePerInstance - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) DescribeForecastPriceHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.ce.forecast"
	apiToBeInvoked := "DescribeForecastPrice"
	onErrMsg := "ce:DescribeForecastPriceHandler - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

package aws

import (
	"net/http"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.ce.history"
	errMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
	apiName := "DescribeCostAndUsage"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.ce.total"
	errMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
	apiName := "DescribeCostAndUsage"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) CostAndUsagePerInstanceTypeHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.ce.instance_type"
	errMsg := "ce:CostAndUsagePerInstance - Credential Missing or Policy Denied"
	apiName := "DescribeCostAndUsagePerInstanceType"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) DescribeForecastPriceHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.ce.forecast"
	errMsg := "ce:DescribeForecastPriceHandler - Credential Missing or Policy Denied"
	apiName := "DescribeForecastPrice"
	handler.API(r, w, apiName, keyCode, errMsg)
}

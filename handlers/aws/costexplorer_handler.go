package aws

import (
	"fmt"
	"net/http"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg := handler.LoadProfileConfigFor(profile, r, w)
	key := fmt.Sprintf("aws.%s.ce.history", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(cfg)
		errMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg := handler.LoadProfileConfigFor(profile, r, w)
	key := fmt.Sprintf("aws.%s.ce.total", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(cfg)
		errMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}

func (handler *AWSHandler) CostAndUsagePerInstanceTypeHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg := handler.LoadProfileConfigFor(profile, r, w)
	key := fmt.Sprintf("aws.%s.ce.instance_type", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsagePerInstanceType(cfg)
		errMsg := "ce:GetCostAndUsage - Credential Missing or Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}

func (handler *AWSHandler) DescribeForecastPriceHandler(w http.ResponseWriter, r *http.Request) {

	profile := r.Header.Get("profile")
	cfg := handler.LoadProfileConfigFor(profile, r, w)
	key := fmt.Sprintf("aws.%s.ce.forecast", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeForecastPrice(cfg)
		errMsg := "ce:GetCostForecast - Credential Missing or Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}

}

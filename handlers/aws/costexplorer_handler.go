package aws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default profile loaded successfuly")
	}

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.ce.history", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set(key, response.History)
			respondWithJSON(w, http.StatusOK, response.History)
		}
	}
}

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}
	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.ce.total", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set(key, response.Total)
			respondWithJSON(w, http.StatusOK, response.Total)
		}
	}
}

func (handler *AWSHandler) CostAndUsagePerInstanceTypeHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("AWSConfig loaded successfuly")
	}
	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.ce.instance_type", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsagePerInstanceType(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, http.StatusOK, response)
		}
	}
}

func (handler *AWSHandler) DescribeForecastPriceHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}
	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.ce.forecast", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeForecastPrice(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostForecast is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, http.StatusOK, response)
		}
	}
}

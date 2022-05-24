package aws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	awsCredential "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/rajasoun/aws-hub/services/aws"
	"github.com/rajasoun/aws-hub/services/cache"

	ini "github.com/rajasoun/go-config-parsers/aws_credentials"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
	aws      aws.AWS
}

func NewAWSHandler(cache cache.Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws:      aws.AWS{},
	}
	return &awsHandler
}

func (handler *AWSHandler) GetAWSHandler() aws.AWS {
	return handler.aws
}

func (handler *AWSHandler) HasMultipleEnvs() bool {
	return handler.multiple
}

func (handler *AWSHandler) API(r *http.Request, w http.ResponseWriter,
	apiName string, keyCode string, errMsg string) {
	profile := r.Header.Get("profile")
	cfg := handler.LoadProfileConfigFor(profile, r, w)
	key := fmt.Sprintf(keyCode, profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.awsAPI(cfg, apiName)
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}

func (handler *AWSHandler) awsAPI(cfg awsCredential.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "DescribeIAMUsers":
		response, err = handler.aws.DescribeIAMUsers(cfg)
	case apiName == "DescribeIAMUser":
		response, err = handler.aws.DescribeIAMUser(cfg)
	case apiName == "DescribeCostAndUsage":
		response, err = handler.aws.DescribeCostAndUsage(cfg)
	case apiName == "DescribeCostAndUsagePerInstanceType":
		response, err = handler.aws.DescribeCostAndUsagePerInstanceType(cfg)
	case apiName == "DescribeForecastPrice":
		response, err = handler.aws.DescribeForecastPrice(cfg)
	}
	return response, err
}

func (handler *AWSHandler) respondWithJSONandSetCache(response interface{}, err error,
	w http.ResponseWriter, errMsg string, key string) {
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errMsg)
	} else {
		handler.cache.Set(key, response)
		respondWithJSON(w, http.StatusOK, response)
	}
}

func (handler *AWSHandler) LoadProfileConfigFor(profile string, r *http.Request,
	w http.ResponseWriter) awsCredential.Config {
	cfg, err := loadAwsConfig(handler.multiple, profile)
	respondOnError(err, w, "Couldn't read "+profile+" profile")
	return cfg
}

func loadAwsConfig(multiple bool, profile string) (awsCredential.Config, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}
	if multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
	}
	return cfg, err
}

func respondOnError(err error, w http.ResponseWriter, errMsg string) {
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errMsg)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func readCredentials(w http.ResponseWriter) ini.Sections {
	sections, err := ini.OpenFile(external.DefaultSharedCredentialsFilename())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
	}
	return sections
}

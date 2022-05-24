package aws

import (
	"encoding/json"
	"log"
	"net/http"

	awsCredential "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"

	"github.com/rajasoun/aws-hub/services/aws"
	"github.com/rajasoun/aws-hub/services/cache"
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

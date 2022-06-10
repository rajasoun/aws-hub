package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws"
	"github.com/rajasoun/aws-hub/service/cache"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
	aws      aws.AWS
}

func NewDefaultAWSHandler(multiple bool) *AWSHandler {
	cache := &cache.Memory{Expiration: time.Duration(30)}
	cache.Connect()
	return NewAWSHandler(cache, multiple)
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
	cfg := handler.LoadCredentialForProfile(profile, r, w)
	key := fmt.Sprintf(keyCode, profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		client := iam.NewFromConfig(cfg)
		response, err := handler.SdkWrapperAPI(client, apiName)
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
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
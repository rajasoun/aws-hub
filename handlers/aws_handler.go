package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/handlers/api"
	"github.com/rajasoun/aws-hub/service/aws"
	"github.com/rajasoun/aws-hub/service/cache"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
	aws      aws.AWS
}

func NewDefaultAWSHandler(multiple bool) *AWSHandler {
	defaultCacheDuration := 30
	cacheHandler := &cache.Memory{Expiration: time.Duration(defaultCacheDuration)}
	cacheHandler.Connect()
	return NewAWSHandler(cacheHandler, multiple)
}

func NewAWSHandler(cacheHandler cache.Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cacheHandler,
		multiple: multiple,
		aws:      aws.AWS{},
	}
	return &awsHandler
}

func (handler *AWSHandler) GetAWSHandler() aws.AWS {
	return handler.aws
}

type AWSWrapper struct {
	request  *http.Request
	writer   http.ResponseWriter
	cache    cache.Cache
	multiple bool
}

func (awsWrapper *AWSWrapper) RespondWithJSON(code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	responseWritter := awsWrapper.writer
	responseWritter.Header().Set("Content-Type", "application/json")
	responseWritter.WriteHeader(code)
	responseWritter.Write(response)
}

func (awsWrapper *AWSWrapper) RespondWithErrorJSON(err error, errMsg string) {
	var errReasons = "Possible Reasons: Connectivity Failed or Credential Missing or Policy Denied"
	errMessage := errMsg + " : " + errReasons
	if err != nil {
		code := http.StatusInternalServerError
		awsWrapper.RespondWithJSON(code, map[string]string{"error": errMessage})
	}
}

func (awsWrapper *AWSWrapper) InvokeAPI(awsAPI api.AwsAPI, cacheKeyCode, errMsg string) {
	profile := awsWrapper.request.Header.Get("profile")
	cacheKey := fmt.Sprintf(cacheKeyCode, profile)
	response, foundInCache := awsWrapper.cache.Get(cacheKey)
	if foundInCache {
		awsWrapper.RespondWithJSON(http.StatusOK, response)
		return
	} else {
		cfg, _ := api.GetConfigFromFileSystem(profile, awsWrapper.multiple)
		client := iam.NewFromConfig(cfg)
		response, err := awsAPI.Execute(client)
		if err != nil {
			awsWrapper.RespondWithErrorJSON(err, errMsg)
		} else {
			awsWrapper.cache.Set(cacheKey, response)
			awsWrapper.RespondWithJSON(http.StatusOK, response)
		}
	}
}

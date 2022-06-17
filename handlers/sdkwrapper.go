package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/provider/credential"
	"github.com/rajasoun/aws-hub/service/cache"
	hubAWS "github.com/rajasoun/aws-hub/service/external/aws"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
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
	}
	return &awsHandler
}

type AWSWrapper struct {
	request  *http.Request
	writer   http.ResponseWriter
	cache    cache.Cache
	multiple bool
}

func NewAWSWrapper(r *http.Request, w http.ResponseWriter, handler *AWSHandler) AWSWrapper {
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	return awsWrapper
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

func (awsWrapper *AWSWrapper) InvokeAPI(awsAPI hubAWS.API, cacheKeyCode, errMsg string) {
	profile := awsWrapper.request.Header.Get("profile")
	cacheKey := fmt.Sprintf(cacheKeyCode, profile)
	response, foundInCache := awsWrapper.cache.Get(cacheKey)
	if foundInCache {
		awsWrapper.RespondWithJSON(http.StatusOK, response)
		return
	}
	// Not In Cache
	cfg, _ := loadConfigFromFileSystem(profile, awsWrapper.multiple)
	client := iam.NewFromConfig(cfg)
	response, err := awsAPI.Execute(client)
	if err != nil {
		awsWrapper.RespondWithErrorJSON(err, errMsg)
		return
	}
	awsWrapper.cache.Set(cacheKey, response)
	awsWrapper.RespondWithJSON(http.StatusOK, response)
}

func loadConfigFromFileSystem(profile string, isMultipleProfile bool) (aws.Config, error) {
	var cfg aws.Config
	var err error
	credentialLoader := credential.New()
	if isMultipleProfile {
		cfg, err = credentialLoader.LoadDefaultConfigForProfile(profile)
		handleErr(err, "AWSConfig For multiple Profile ")
	} else {
		cfg, err = credentialLoader.LoadDefaultConfig()
		handleErr(err, "Default AWSConfig")
	}
	return cfg, err
}

func handleErr(err error, msg string) {
	if err != nil {
		log.Println(msg+" Load Failed err = %v", err)
	} else {
		log.Println(msg + "loaded successfully")
	}
}

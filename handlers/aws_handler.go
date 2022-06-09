package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/handlers/api"
	"github.com/rajasoun/aws-hub/service/cache"
)

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
	var errReasons string = "Possible Reasons: Connectivity Failed or Credential Missing or Policy Denied"
	errMessage := errMsg + " : " + errReasons
	if err != nil {
		code := http.StatusInternalServerError
		awsWrapper.RespondWithJSON(code, map[string]string{"error": errMessage})
	}
}

func (awsWrapper *AWSWrapper) InvokeAPI(awsApi api.AwsAPI, keyCode string, errMsg string) {
	profile := awsWrapper.request.Header.Get("profile")
	key := fmt.Sprintf(keyCode, profile)
	response, foundInCache := awsWrapper.cache.Get(key)
	if foundInCache {
		awsWrapper.RespondWithJSON(http.StatusOK, response)
		return
	} else {
		cfg, _ := api.GetConfig(profile, awsWrapper.multiple)
		client := iam.NewFromConfig(cfg)
		response, err := awsApi.Execute(client)
		if err != nil {
			awsWrapper.RespondWithErrorJSON(err, errMsg)
		} else {
			awsWrapper.cache.Set(key, response)
			awsWrapper.RespondWithJSON(http.StatusOK, response)
		}
	}
}

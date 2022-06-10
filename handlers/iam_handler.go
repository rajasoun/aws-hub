package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajasoun/aws-hub/handlers/api"
)

var cacheKeyTemplate = "aws.%s.iam."

func (handler *AWSHandler) IAMGetUserCountHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "users"
	apiToBeInvoked := api.IAMGetUserCountAPI
	onErrMsg := "iam:GetUserCount - Failed."
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	api := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(api, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMGetUserIdentityHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "useraccount"
	apiToBeInvoked := api.IAMGetUserIdentityAPI
	onErrMsg := "iam:GetUserIdentity - Failed."
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	api := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(api, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMGetAliasesHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "aliases"
	apiToBeInvoked := api.IAMGetAliasesAPI
	onErrMsg := "iam:GetAliases - Failed."
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	api := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(api, cacheKey, onErrMsg)
}

func GetAPI(r *http.Request, apiToBeInvoked string) api.AwsAPI {
	params := mux.Vars(r)
	apiName := params["ApiName"]
	if apiName == api.IAMPing {
		log.Println(" IAM Ping API ")
		return api.NewAwsAPI(api.IAMPing)
	}
	return api.NewAwsAPI(apiToBeInvoked)
}

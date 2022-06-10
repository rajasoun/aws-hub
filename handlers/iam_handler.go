package handlers

import (
	"net/http"

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
	api := api.NewAwsAPI(apiToBeInvoked)
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
	api := api.NewAwsAPI(apiToBeInvoked)
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
	api := api.NewAwsAPI(apiToBeInvoked)
	awsWrapper.InvokeAPI(api, cacheKey, onErrMsg)
}

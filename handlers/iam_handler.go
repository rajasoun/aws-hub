package handlers

import (
	"net/http"
)

var errReasons string = "Possible Reasons: Connectivity Failed or Credential Missing or Policy Denied"

func (handler *AWSHandler) IAMGetUserCountHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.iam.users"
	apiToBeInvoked := "GetUserCount"
	onErrMsg := "iam:GetUserCount - Failed." + errReasons
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	awsWrapper.InvokeAPI(apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMGetUserIdentityHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.iam.user"
	apiToBeInvoked := "GetUserIdentity"
	onErrMsg := "iam:GetUserIdentity - Failed." + errReasons
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	awsWrapper.InvokeAPI(apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMGetAliasesHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.iam.aliases"
	apiToBeInvoked := "GetAliases"
	onErrMsg := "iam:GetAliases - Failed." + errReasons
	awsWrapper := AWSWrapper{
		request:  r,
		writer:   w,
		cache:    handler.cache,
		multiple: handler.multiple,
	}
	awsWrapper.InvokeAPI(apiToBeInvoked, cacheKey, onErrMsg)
}

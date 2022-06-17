package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/rajasoun/aws-hub/api/external/aws"
)

const cacheKeyTemplate = "aws.%s.iam."

// Get Number of Users associated to a AWS Account.
func (handler *AWSHandler) IAMGetUserCountHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "users"
	apiToBeInvoked := api.IAMGetUserCountAPI
	onErrMsg := "iam:GetUserCount - Failed."
	awsWrapper := NewAWSWrapper(r, w, handler)
	awsAPI := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(awsAPI, cacheKey, onErrMsg)
}

// Get User Identity Details for the user associated to a AWS Account.
func (handler *AWSHandler) IAMGetUserIdentityHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "useraccount"
	apiToBeInvoked := api.IAMGetUserIdentityAPI
	onErrMsg := "iam:GetUserIdentity - Failed."
	awsWrapper := NewAWSWrapper(r, w, handler)
	awsAPI := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(awsAPI, cacheKey, onErrMsg)
}

// Get Aliases for the  AWS Account.
func (handler *AWSHandler) IAMGetAliasesHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := cacheKeyTemplate + "aliases"
	apiToBeInvoked := api.IAMGetAliasesAPI
	onErrMsg := "iam:GetAliases - Failed."
	awsWrapper := NewAWSWrapper(r, w, handler)
	awsAPI := GetAPI(r, apiToBeInvoked)
	awsWrapper.InvokeAPI(awsAPI, cacheKey, onErrMsg)
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

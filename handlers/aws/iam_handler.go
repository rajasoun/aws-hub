package aws

import (
	"net/http"
)

func (handler *AWSHandler) IAMListUsersHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.iam.users"
	apiToBeInvoked := "IAMListUsers"
	onErrMsg := "iam:ListUsers - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	cacheKey := "aws.%s.iam.user"
	apiToBeInvoked := "IAMUser"
	onErrMsg := "iam:GetUser - Credential Missing or Policy Denied"
	handler.API(r, w, apiToBeInvoked, cacheKey, onErrMsg)
}

func (handler *AWSHandler) IAMAliasHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.aliases"
	errMsg := "iam:GetAliases - Credential Missing or Policy Denied"
	apiName := "IAMAliases"
	handler.API(r, w, apiName, keyCode, errMsg)
}

package aws

import (
	"net/http"
)

func (handler *AWSHandler) IAMListUsersHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.users"
	errMsg := "iam:ListUsers - Credential Missing or Policy Denied"
	apiName := "IAMListUsers"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.user"
	errMsg := "iam:GetUser - Credential Missing or Policy Denied"
	apiName := "IAMUser"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) IAMAliasHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.aliases"
	errMsg := "iam:GetAliases - Credential Missing or Policy Denied"
	apiName := "IAMAliases"
	handler.API(r, w, apiName, keyCode, errMsg)
}

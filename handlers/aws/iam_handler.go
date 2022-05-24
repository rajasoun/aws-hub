package aws

import (
	"net/http"
)

func (handler *AWSHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.users"
	errMsg := "iam:ListUsers - Credential Missing or Policy Denied"
	apiName := "DescribeIAMUsers"
	handler.API(r, w, apiName, keyCode, errMsg)
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	keyCode := "aws.%s.iam.user"
	errMsg := "iam:GetUser - Credential Missing or Policy Denied"
	apiName := "DescribeIAMUser"
	handler.API(r, w, apiName, keyCode, errMsg)
}

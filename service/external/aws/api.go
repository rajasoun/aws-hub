package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const IAMGetUserCountAPI = "GetUserCountAPI"
const IAMGetUserIdentityAPI = "GetUserIdentityAPI"
const IAMGetAliasesAPI = "GetAliasesAPI"
const IAMPing = "DoPing"

type API interface {
	Execute(client *iam.Client) (interface{}, error)
}

func New(api string) API {
	var awsAPI API
	switch api {
	case IAMGetUserCountAPI:
		awsAPI = GetUserCountAPI{}
	case IAMGetUserIdentityAPI:
		awsAPI = GetUserIdentityAPI{}
	case IAMGetAliasesAPI:
		awsAPI = GetAliasesAPI{}
	case IAMPing:
		awsAPI = DoPing{}
	}
	return awsAPI
}

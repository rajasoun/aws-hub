package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	hubIAM "github.com/rajasoun/aws-hub/service/external/aws/iam"
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
		awsAPI = hubIAM.GetUserCountAPI{}
	case IAMGetUserIdentityAPI:
		awsAPI = hubIAM.GetUserIdentityAPI{}
	case IAMGetAliasesAPI:
		awsAPI = hubIAM.GetAliasesAPI{}
	case IAMPing:
		awsAPI = hubIAM.DoPing{}
	}
	return awsAPI
}

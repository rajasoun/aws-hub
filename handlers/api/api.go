package api

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const IAMGetUserCountAPI = "GetUserCountAPI"
const IAMGetUserIdentityAPI = "GetUserIdentityAPI"
const IAMGetAliasesAPI = "GetAliasesAPI"

type AwsAPI interface {
	Execute(client *iam.Client) (interface{}, error)
}

func NewAwsAPI(api string) AwsAPI {
	var awsAPI AwsAPI
	switch api {
	case IAMGetUserCountAPI:
		awsAPI = GetUserCountAPI{}
	case IAMGetUserIdentityAPI:
		awsAPI = GetUserIdentityAPI{}
	case IAMGetAliasesAPI:
		awsAPI = GetAliasesAPI{}
	}
	return awsAPI
}

func GetConfig(profile string, isMultipleProfile bool) (aws.Config, error) {
	credentialLoader := CredentialLoader{}
	cfg, err := credentialLoader.GetConfig(profile, isMultipleProfile)
	return cfg, err
}

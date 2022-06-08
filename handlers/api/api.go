package api

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AwsAPI interface {
	Execute(config aws.Config) (interface{}, error)
}

func NewAwsAPI(api string) AwsAPI {
	var awsAPI AwsAPI
	switch api {
	case "GetUserCount":
		awsAPI = GetUserCountAPI{}
	case "GetUserIdentity":
		awsAPI = GetUserIdentityAPI{}
	case "GetAliases":
		awsAPI = GetAliasesAPI{}
	}
	return awsAPI
}

func GetConfig(profile string, isMultipleProfile bool) (aws.Config, error) {
	credentialHandler := CredentialHandler{}
	cfg, err := credentialHandler.GetConfig(profile, isMultipleProfile)
	return cfg, err
}

package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/provider/credential"
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

func GetConfigFromFileSystem(profile string, isMultipleProfile bool) (aws.Config, error) {
	var cfg aws.Config
	var err error
	credentialLoader := credential.New()
	if isMultipleProfile {
		cfg, err = credentialLoader.LoadDefaultConfigForProfile(profile)
		handleErr(err, "AWSConfig For multiple Profile ")
	} else {
		cfg, err = credentialLoader.LoadDefaultConfig()
		handleErr(err, "Default AWSConfig")
	}
	return cfg, err
}

func handleErr(err error, msg string) {
	if err != nil {
		log.Println(msg+" Load Failed err = %v", err)
	} else {
		log.Println(msg + "loaded successfully")
	}
}

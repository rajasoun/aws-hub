package api

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/rajasoun/aws-hub/provider/credential"
)

type CredentialLoader struct{}

func (credHandler *CredentialLoader) GetConfig(profile string, isMultipleProfile bool) (aws.Config, error) {
	var cfg aws.Config
	var err error
	credentialLoader := credential.CredentialLoader{}
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
		log.Println(msg + "loaded successfuly")
	}
}

package aws

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/rajasoun/aws-hub/provider/credential"
)

func (handler *AWSHandler) LoadConfigForProfile(profile string, r *http.Request,
	w http.ResponseWriter) aws.Config {
	var cfg aws.Config
	var err error
	credentialLoader := credential.CredentialLoader{}
	if handler.multiple {
		cfg, err = credentialLoader.LoadDefaultConfigForProfile(profile)
		handleErr(err, "AWSConfig For multiple Profile ")
	} else {
		cfg, err = credentialLoader.LoadDefaultConfig()
		handleErr(err, "Default AWSConfig")
	}
	respondOnError(err, w, "Couldn't read "+profile+" profile")
	return cfg
}

func handleErr(err error, msg string) {
	if err != nil {
		log.Println(msg+" Load Failed err = %v", err)
	} else {
		log.Println(msg + "loaded successfuly")
	}
}

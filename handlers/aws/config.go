package aws

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	ini "github.com/rajasoun/go-config-parsers/aws_credentials"
)

func loadLocalAwsConfig(multiple bool, profile string) (aws.Config, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}
	if multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
	}
	return cfg, err
}

func readLocalCredentials(w http.ResponseWriter) ini.Sections {
	sections, err := ini.OpenFile(external.DefaultSharedCredentialsFilename())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
	}
	return sections
}

func (handler *AWSHandler) loadProfileConfigFor(profile string, r *http.Request,
	w http.ResponseWriter) aws.Config {
	cfg, err := loadLocalAwsConfig(handler.multiple, profile)
	respondOnError(err, w, "Couldn't read "+profile+" profile")
	return cfg
}

package aws

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	ini "github.com/rajasoun/go-parsers/aws_credentials"
)

func loadLocalAwsConfig(multiple bool, profile string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}

	if multiple {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-east-1"),
			config.WithSharedConfigFiles(
				config.DefaultSharedConfigFiles,
			),
			config.WithSharedConfigProfile(profile),
		)
	}
	return cfg, err
}

func readLocalCredentials(w http.ResponseWriter) ini.Sections {
	sections, err := ini.OpenFile(config.DefaultSharedCredentialsFilename())
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

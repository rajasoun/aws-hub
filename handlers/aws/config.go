package aws

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	ini "github.com/rajasoun/go-parsers/aws_credentials"
)

func handleErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg+" Load Failed err = %v", err)
	} else {
		log.Println(msg + "loaded successfuly")
	}
}

func loadLocalAwsConfig(multiple bool, profile string) (aws.Config, error) {
	var cfg aws.Config
	var err error
	log.Println("Profile -> ", profile)
	if multiple {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-east-1"),
			config.WithSharedConfigFiles(config.DefaultSharedConfigFiles),
			config.WithSharedConfigProfile(profile),
		)
		handleErr(err, "AWSConfig For multiple Profile ")
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		handleErr(err, "Default AWSConfig")
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

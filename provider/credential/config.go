package credential

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	ini "github.com/rajasoun/go-parsers/aws_credentials"
)

var defaultRegion string = "us-east-1"

type ConfigLoader interface {
	LoadDefaultConfig() (cfg aws.Config, err error)
	LoadDefaultConfigForProfile(profile string) (cfg aws.Config, err error)
}

type CredentialLoader struct{}

// Loads Default Configuration from files ~/.aws/config and ~/.aws/credential
// Output:
//     If successful, aws.Config struct & nil
//     Otherwise, empty aws.Config and an error.
//	   Ref: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
func (credLoader *CredentialLoader) LoadDefaultConfig() (aws.Config, error) {
	context := context.TODO()
	region := config.WithRegion(defaultRegion)
	cfg, err := config.LoadDefaultConfig(context, region)
	if err != nil {
		log.Printf("failed to load default configuration, %v", err)
	}
	return cfg, err
}

// Loads Default Configuration from files ~/.aws/config and ~/.aws/credential
// Input:
//	   Profile Name in the ~/.aws/config file
// Output:
//     If successful, aws.Config struct & nil
//     Otherwise, empty aws.Config and an error.
//	   Ref: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
func (credLoader *CredentialLoader) LoadDefaultConfigForProfile(profile string) (aws.Config, error) {
	context := context.TODO()
	region := config.WithRegion(defaultRegion)
	sharedConfigFiles := config.WithSharedConfigFiles(config.DefaultSharedConfigFiles)
	profileOpt := config.WithSharedConfigProfile(profile)
	cfg, err := config.LoadDefaultConfig(context, region, sharedConfigFiles, profileOpt)
	if err != nil {
		log.Printf("failed to load configuration for profile - %v, err = %v", profile, err)
	}
	return cfg, err
}

func (credLoader *CredentialLoader) GetSections() (ini.Sections, error) {
	sections, err := ini.OpenFile(config.DefaultSharedCredentialsFilename())
	return sections, err
}

package credential

import (
	"context"
	"errors"
	"log"
	"os"

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

// Loads Credentials from files ~/.aws/credential
// Output:
//     If successful, Returns sections within the ~/.aws/credential file
//     Otherwise, empty sections and an error.
func (credLoader *CredentialLoader) GetSections(credentialFile string) (ini.Sections, error) {
	//credentialFile := config.DefaultSharedCredentialsFilename()
	if fileExists(credentialFile) {
		return ini.OpenFile(credentialFile)
	} else {
		return ini.Sections{}, errors.New("credential file is not available")
	}

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

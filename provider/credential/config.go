package credential

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var defaultRegion string = "us-east-1"

type ConfigManager interface {
	LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)
}

// Loads Default Configuration from files ~/.aws/config and ~/.aws/credential
// Output:
//     If successful, aws.Config struct & nil
//     Otherwise, empty aws.Config and an error.
//	   Ref: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
func LoadDefaultProfile() (aws.Config, error) {
	context := context.TODO()
	region := config.WithRegion(defaultRegion)
	cfg, err := config.LoadDefaultConfig(context, region)
	if err != nil {
		log.Printf("failed to load default configuration, %v", err)
	}
	return cfg, err
}

package credential

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var defaultRegion string = "us-east-1"

func LoadDefaultProfile() (aws.Config, error) {
	context := context.TODO()
	region := config.WithRegion(defaultRegion)
	cfg, err := config.LoadDefaultConfig(context, region)
	return cfg, err
}

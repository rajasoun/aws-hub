package credential

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

var DefaultRegion string = "us-east-1"

func LoadConfig(profile string) (aws.Config, error) {
	var cfg aws.Config
	var err error
	if profile != "default" {
		cfg, err = LoadProfile(profile)
	} else {
		cfg, err = LoadDefaultProfile()
	}

	logMsg := fmt.Sprintf("credential loading from local file for profele %s", profile)
	if err != nil {
		log.Printf(logMsg + " : FAIL")
	} else {
		log.Printf(logMsg + " : PASS")
	}
	return cfg, err
}

func LoadProfile(profile string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(DefaultRegion),
		config.WithSharedConfigFiles(config.DefaultSharedConfigFiles),
		config.WithSharedConfigProfile(profile),
	)
	return cfg, err
}

func LoadDefaultProfile() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(DefaultRegion))
	return cfg, err
}

func CheckConfig(cfg aws.Config) bool {
	var ctx context.Context = context.TODO()
	input := &iam.GetUserInput{}
	client := iam.NewFromConfig(cfg)
	_, err := client.GetUser(ctx, input)
	if err != nil {
		log.Printf("Get User API call failed: %s", err)
		return false
	}
	return true
}

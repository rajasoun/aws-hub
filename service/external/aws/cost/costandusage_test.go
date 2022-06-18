package cost

import (
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/rajasoun/aws-hub/provider/credential"
)

func TestGetCost(t *testing.T) {
	credentialLoader := credential.New()
	cfg, _ := credentialLoader.LoadDefaultConfig()
	client := costexplorer.NewFromConfig(cfg)
	result, err := GetCost(client)
	if err != nil {
		log.Printf("Err = %v", err)
	}
	log.Printf("Result = %v", result.Total)
	t.Fail()
}

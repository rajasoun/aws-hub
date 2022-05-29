package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	hubIAM "github.com/rajasoun/aws-hub/services/aws/iam"
)

// ToDo Technical Debt - Use Interface to call the right method
// Use Dependency Injection
func (aws AWS) ExternalServiceGateway(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	var client *iam.Client = iam.NewFromConfig(cfg)
	switch {
	case apiName == "IAMListUsers":
		response, err = hubIAM.GetUserCount(client)
	case apiName == "IAMUser":
		response, err = hubIAM.GetUserIdentity(client)
	case apiName == "IAMAliases":
		response, err = hubIAM.GetAliases(client)
	}
	return response, err
}

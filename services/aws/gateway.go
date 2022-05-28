package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rajasoun/aws-hub/services/aws/iam"
)

// ToDo Technical Debt - Use Interface to call the right method
// Use Dependency Injection
func (aws AWS) ExternalServiceGateway(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "IAMListUsers":
		response, err = iam.GetUserCountForAccount(cfg)
	case apiName == "IAMUser":
		response, err = iam.GetUserIdentity(cfg)
	case apiName == "IAMAliases":
		response, err = iam.GetAliases(cfg)
	}
	return response, err
}

package handlers

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"

	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
)

// ToDo Technical Debt - Use Interface to call the right method
// Use Dependency Injection
func (handler *AWSHandler) SdkWrapperAPI(client *iam.Client, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "GetUserCount":
		response, err = hubIAM.GetUserCount(client)
	case apiName == "GetUserIdentity":
		response, err = hubIAM.GetUserIdentity(client)
	case apiName == "GetAliases":
		response, err = hubIAM.GetAliases(client)
	}
	return response, err
}

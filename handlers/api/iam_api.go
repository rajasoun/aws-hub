package api

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	service "github.com/rajasoun/aws-hub/service/aws/iam"
)

type GetUserCountAPI struct{}

func (api GetUserCountAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := service.GetUserCount(client)
	return response, err
}

type GetUserIdentityAPI struct{}

func (api GetUserIdentityAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := service.GetUserIdentity(client)
	return response, err
}

type GetAliasesAPI struct{}

func (api GetAliasesAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := service.GetAliases(client)
	return response, err
}

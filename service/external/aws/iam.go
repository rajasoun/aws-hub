package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamAPI "github.com/rajasoun/aws-hub/service/external/aws/iam"
)

type GetUserCountAPI struct{}

func (api GetUserCountAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := iamAPI.GetUserCount(client)
	return response, err
}

type GetUserIdentityAPI struct{}

func (api GetUserIdentityAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := iamAPI.GetUserIdentity(client)
	return response, err
}

type GetAliasesAPI struct{}

func (api GetAliasesAPI) Execute(client *iam.Client) (interface{}, error) {
	response, err := iamAPI.GetAliases(client)
	return response, err
}

type DoPing struct{}

type Ping struct {
	Status string `json:"status"`
}

func (api DoPing) Execute(client *iam.Client) (interface{}, error) {
	response := Ping{Status: "Ok"}
	return response, nil
}

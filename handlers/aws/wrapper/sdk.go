package wrapper

import (
	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
)

type SDKV2 interface {
	Execute() error
}

type SDK struct{}

func (sdk SDK) ExecuteAPI(sdkv2 SDKV2) error {
	return sdkv2.Execute()
}

type Account struct {
	client apiclient.IAMListAccountAliasesAPIClient
	alias  hubIAM.Aliases
}

func (account *Account) Execute() error {
	response, err := hubIAM.GetAliases(account.client)
	account.alias = response
	return err
}

type UserCount struct {
	client   apiclient.IAMListUsersAPIClient
	userList hubIAM.UserList
}

func (userCount *UserCount) Execute() error {
	response, err := hubIAM.GetUserCount(userCount.client)
	userCount.userList = response
	return err
}

type UserIdentity struct {
	client   apiclient.IAMGetUserAPIClient
	identity hubIAM.User
}

func (userIdentity *UserIdentity) Execute() error {
	response, err := hubIAM.GetUserIdentity(userIdentity.client)
	userIdentity.identity = response
	return err
}

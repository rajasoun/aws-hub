package iammock

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/mock"
)

type iamClient interface {
	ListAccountAliases(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)

	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}

func TestListAccountAliases(t *testing.T) {
	client := new(Client)
	methodName := "ListAccountAliases"
	output := iam.ListAccountAliasesOutput{}
	client.
		On(methodName, mock.Anything, mock.Anything, mock.Anything).
		Return(&output, nil)
	fn := func(c iamClient) {
		_, _ = c.ListAccountAliases(context.Background(), nil, nil)
	}
	fn(client)
	client.AssertExpectations(t)
}

func TestListUsers(t *testing.T) {
	client := new(Client)
	methodName := "ListUsers"
	output := iam.ListUsersOutput{}
	client.
		On(methodName, mock.Anything, mock.Anything, mock.Anything).
		Return(&output, nil)
	fn := func(c iamClient) {
		_, _ = c.ListUsers(context.Background(), nil, nil)
	}
	fn(client)
	client.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	client := new(Client)
	methodName := "GetUser"
	output := iam.GetUserOutput{}
	client.
		On(methodName, mock.Anything, mock.Anything, mock.Anything).
		Return(&output, nil)
	fn := func(c iamClient) {
		_, _ = c.GetUser(context.Background(), nil, nil)
	}
	fn(client)
	client.AssertExpectations(t)
}

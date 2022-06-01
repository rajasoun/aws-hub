package iammock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/mock"
)

// Client provides the API client to mock AWS operations
type MockClient struct {
	mock.Mock
}

func (c *MockClient) InjectFunctionMock(client *MockClient, methodName string, result interface{}) {
	client.
		On(methodName, mock.Anything, mock.Anything, mock.Anything).
		Return(result, nil)
}

// List Account Aliases Mock
func (c *MockClient) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*iam.ListAccountAliasesOutput), args.Error(1)
}

// List Users Mock
func (c *MockClient) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*iam.ListUsersOutput), args.Error(1)
}

// Get User Mock
func (c *MockClient) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*iam.GetUserOutput), args.Error(1)
}

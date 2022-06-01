package iammock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/mock"
)

var testAlias string = "aws-test-account-alias"

// Client provides the API client to mock operations call for Amazon Simple Queue Service.

type MockClient struct {
	mock.Mock
}

// Interface for Amazon IAM ListAccountAliases API
// This will enable TDD using mocking
type IAMListAccountAliasesAPIClient interface {
	iam.ListAccountAliasesAPIClient // Only for Refernce to Actual Client
	ListAccountAliases(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)
}

// Mock Receiver
type MockAccountAliases struct{}

//Mock Function
type MockListAccountAliasesAPIClient func(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)

// Implement AWS IAM ListAccountAliases Interface with mock reciever
func (mock MockListAccountAliasesAPIClient) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockAccountAliases) NewMockClient() IAMListAccountAliasesAPIClient {
	client := MockListAccountAliasesAPIClient(func(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
		aliases := []string{testAlias}
		result := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		return result, nil
	})
	return client
}

// // List Account Aliases Mock
// func (c *MockClient) ListAccountAliases(ctx context.Context,
// 	params *iam.ListAccountAliasesInput,
// 	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
// 	args := c.Called(ctx, params, optFns)
// 	// if args.Get(0) == nil {
// 	// 	return nil, args.Error(1)
// 	// }
// 	return args.Get(0).(*iam.ListAccountAliasesOutput), args.Error(1)
// }

// // List Users Count
// func (c *MockClient) ListUsers(ctx context.Context,
// 	params *iam.ListUsersInput,
// 	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
// 	args := c.Called(ctx, params, optFns)
// 	// if args.Get(0) == nil {
// 	// 	return nil, args.Error(1)
// 	// }
// 	return args.Get(0).(*iam.ListUsersOutput), args.Error(1)
// }

// // List User Identity
// func (c *MockClient) GetUser(ctx context.Context,
// 	params *iam.GetUserInput,
// 	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
// 	args := c.Called(ctx, params, optFns)
// 	// if args.Get(0) == nil {
// 	// 	return nil, args.Error(1)
// 	// }
// 	return args.Get(0).(*iam.GetUserOutput), args.Error(1)
// }

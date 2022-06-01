package iammock

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
)

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

func (mock MockAccountAliases) NewClient() apiclient.IAMListAccountAliasesAPIClient {
	fn := func(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
		var testAlias string = "aws-test-account-alias"
		aliases := []string{testAlias}
		result := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		return result, nil
	}
	client := MockListAccountAliasesAPIClient(fn)
	return client
}

// Mock Receiver
type MockUser struct{}

//Mock Function
type MockIAMListUsersAPIClient func(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

// Implement AWS IAM ListUsers Method with Mock Function Receiver
func (mock MockIAMListUsersAPIClient) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockUser) NewClient() apiclient.IAMListUsersAPIClient {
	client := MockIAMListUsersAPIClient(func(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
		userList := []types.User{
			{UserName: aws.String("test1@example.com")},
			{UserName: aws.String("test2@example.com")},
		}
		result := &iam.ListUsersOutput{Users: userList}
		return result, nil
	})
	return client
}

// Mock Receiver
type MockUserIdentity struct{}

//Mock Function
type MockIAMGetUserAPIClient func(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)

// Implement AWS IAM GetUser Method with Mock Function Receiver
func (mock MockIAMGetUserAPIClient) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockUserIdentity) NewClient() apiclient.IAMGetUserAPIClient {
	client := MockIAMGetUserAPIClient(func(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
		user := &types.User{
			Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
			CreateDate:       &time.Time{},
			UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
			UserName:         aws.String("test@example.com"),
			PasswordLastUsed: &time.Time{},
		}
		result := &iam.GetUserOutput{User: user}
		return result, nil
	})
	return client
}

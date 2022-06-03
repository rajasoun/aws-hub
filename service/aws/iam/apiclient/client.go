package apiclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

// Interface for Amazon IAM ListUsers API
// This will enable TDD using mocking
type IAMListUsersAPIClient interface {
	iam.ListUsersAPIClient // Only for Refernce to Actual Client
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// Interface for Amazon IAM GetUser API
// This will enable TDD using mocking
type IAMGetUserAPIClient interface {
	iam.GetUserAPIClient // Only for Refernce to Actual Client
	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type UserList struct {
	Count int `json:"usercount"`
}

// Interface for Amazon IAM ListUsers API
// This will enable TDD using mocking
type IAMListUsersAPIClient interface {
	iam.ListUsersAPIClient // Only for Refernce to Actual Client
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// GetUserCount retrieves the user accounty for an AWS account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a Users object containing the count and nil.
//     Otherwise, nil and an error from the call.
func GetUserCount(client IAMListUsersAPIClient) (UserList, error) {
	var ctx context.Context = context.TODO()
	input := &iam.ListUsersInput{}
	result, err := client.ListUsers(ctx, input)
	if err != nil {
		return UserList{Count: 0}, err
	}
	return UserList{Count: len(result.Users)}, nil
}

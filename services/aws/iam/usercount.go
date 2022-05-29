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
type IAMListUsersAPI interface {
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// Wrapper Function to AWS IAM GetUser API with client as argument
// This will enable TDD by passing mock client
func ListUsers(ctx context.Context, client IAMListUsersAPI,
	input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	return client.ListUsers(ctx, input)
}

// GetUserCount retrieves the user accounty for an AWS account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a Users object containing the count and nil.
//     Otherwise, nil and an error from the call.
func GetUserCount(client IAMListUsersAPI) (UserList, error) {
	var ctx context.Context = context.TODO()
	input := &iam.ListUsersInput{}
	result, err := ListUsers(ctx, client, input)
	if err != nil {
		return UserList{Count: 0}, err
	}
	return UserList{Count: len(result.Users)}, nil
}

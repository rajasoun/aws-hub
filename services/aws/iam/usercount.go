package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type UserList struct {
	Count int `json:"usercount"`
}

// Interface for Amazon IAM API operations required by ListUsers function
// This will enable TDD using mocking the wrapped function
type IAMListUsersAPI interface {
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// Amazon IAM client’s ListUsers method,
func ListUsers(ctx context.Context, client IAMListUsersAPI,
	input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	return client.ListUsers(ctx, input)
}

// GetUserCount retrieves the user accounty for an AWS account.
// Inputs:
//     cfg is the context of the method call, which includes the AWS Region.
// Output:
//     If successful, a Users object containing the count and nil.
//     Otherwise, nil and an error from the call.
func GetUserCount(cfg aws.Config, client IAMListUsersAPI) (UserList, error) {
	ctx := context.TODO()
	//client := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(ctx, client, input)
	if err != nil {
		return UserList{Count: 0}, err
	}
	return UserList{Count: len(result.Users)}, nil
}

func GetUserCountForAccount(cfg aws.Config) (UserList, error) {
	return GetUserCount(cfg, iam.NewFromConfig(cfg))
}

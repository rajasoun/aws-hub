package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type UserList struct {
	Count int `json:"usercount"`
}

// Interface wraps up the underlying AWS Function
// This will enable TDD using mocking the wrapped function

type ListUsersAPI interface {
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// Wrapper Function to ListUsers with api to be called as argument
// This will enable TDD using mocking the wrapped function
func ListUsers(c context.Context, client ListUsersAPI,
	input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	log.Println("[usercount.go][ListUsers] client -> ", client)
	return client.ListUsers(c, input)
}

// GetUserCount retrieves the user accounty for an AWS account.
// Inputs:
//     cfg is the context of the method call, which includes the AWS Region.
// Output:
//     If successful, a Users object containing the count and nil.
//     Otherwise, nil and an error from the call.
func GetUserCount(cfg aws.Config) (UserList, error) {
	ctx := context.TODO()
	client := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(ctx, client, input)
	if err != nil {
		return UserList{Count: 0}, err
	}
	return UserList{Count: len(result.Users)}, nil
}

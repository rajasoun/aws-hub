package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type User struct {
	Username         string    `json:"username"`
	ARN              string    `json:"arn"`
	CreateDate       time.Time `json:"createDate"`
	PasswordLastUsed time.Time `json:"passwordLastUsed"`
	UserId           string    `json:"userId"`
}

// Interface wraps up the underlying AWS Function
// This will enable TDD using mocking the wrapped function
type GetUserAPI interface {
	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}

// Wrapper Function to ListUsers with api to be called as argument
// This will enable TDD using mocking the wrapped function
func GetUser(c context.Context, api GetUserAPI,
	input *iam.GetUserInput) (*iam.GetUserOutput, error) {
	return api.GetUser(c, input)
}

// GetUserIdentity retrieves the user details from an AWS account.
// Inputs:
//     cfg is the context of the method call, which includes the AWS Region.
// Output:
//     If successful, a Users object containing the account details and nil.
//     Otherwise, nil and an error from the call.
func GetUserIdentity(cfg aws.Config) (User, error) {
	api := iam.NewFromConfig(cfg)
	input := &iam.GetUserInput{}
	result, err := GetUser(context.TODO(), api, input)

	if err != nil {
		return User{}, err
	}

	lastUsed := time.Now()
	if result.User.PasswordLastUsed != nil {
		lastUsed = *result.User.PasswordLastUsed
	}

	userAccount := User{
		Username:         *result.User.UserName,
		ARN:              *result.User.Arn,
		CreateDate:       *result.User.CreateDate,
		UserId:           *result.User.UserId,
		PasswordLastUsed: lastUsed,
	}
	return userAccount, nil
}

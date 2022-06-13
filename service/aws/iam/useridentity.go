package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type User struct {
	Username         string    `json:"username"`
	ARN              string    `json:"arn"`
	CreateDate       time.Time `json:"createDate"`
	PasswordLastUsed time.Time `json:"passwordLastUsed"`
	UserID           string    `json:"userId"`
}

// Interface for Amazon IAM GetUser API
// This will enable TDD using mocking
type IAMGetUserAPIClient interface {
	iam.GetUserAPIClient // Only for Reference to Actual Client
	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}

// GetUserIdentity retrieves the user details from an AWS account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a Users object containing the account details and nil.
//     Otherwise, nil and an error from the call.
func GetUserIdentity(client IAMGetUserAPIClient) (User, error) {
	emptyContext := context.TODO()
	input := &iam.GetUserInput{}
	result, err := client.GetUser(emptyContext, input)

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
		UserID:           *result.User.UserId,
		PasswordLastUsed: lastUsed,
	}
	return userAccount, nil
}

package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IAMUser struct {
	Username         string    `json:"username"`
	ARN              string    `json:"arn"`
	CreateDate       time.Time `json:"createDate"`
	PasswordLastUsed time.Time `json:"passwordLastUsed"`
	UserId           string    `json:"userId"`
}

// IAMListUsersAPI defines the interface for the ListUsers function.
// We use this interface to test the function using a mocked service.
type IAMListUsersAPI interface {
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

// ListUsers retrieves a list of your AWS Identity and Access Management (IAM) users.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If successful, a ListUsersOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ListUsers.
func ListUsers(c context.Context, api IAMListUsersAPI, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	return api.ListUsers(c, input)
}

func (aws AWS) IAMListUsers(cfg aws.Config) (int, error) {
	svc := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(context.TODO(), svc, input)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return len(result.Users), nil
}

func (aws AWS) IAMUser(cfg aws.Config) (IAMUser, error) {
	svc := iam.NewFromConfig(cfg)
	input := &iam.GetUserInput{}
	result, err := svc.GetUser(context.TODO(), input)

	if err != nil {
		return IAMUser{}, err
	}

	lastUsed := time.Now()
	if result.User.PasswordLastUsed != nil {
		lastUsed = *result.User.PasswordLastUsed
	}

	return IAMUser{
		Username:         *result.User.UserName,
		ARN:              *result.User.Arn,
		CreateDate:       *result.User.CreateDate,
		UserId:           *result.User.UserId,
		PasswordLastUsed: lastUsed,
	}, nil
}

func GetUser(context context.Context, svc *iam.Client, input *iam.ListUsersInput) {
	panic("unimplemented")
}

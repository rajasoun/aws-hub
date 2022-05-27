package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IAMUsersCount struct {
	UsersCount int `json:"usercount"`
}

// IAMListUsersAPI defines the interface for the ListUsers function.
type IAMListUsersAPI interface {
	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
}

func ListUsers(c context.Context, api IAMListUsersAPI, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	return api.ListUsers(c, input)
}

func (aws AWS) IAMListUsers(cfg aws.Config) (IAMUsersCount, error) {
	svc := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(context.TODO(), svc, input)
	if err != nil {
		fmt.Println(err)
		return IAMUsersCount{
			UsersCount: 0,
		}, err
	}
	return IAMUsersCount{
		UsersCount: len(result.Users),
	}, nil
}

type IAMUser struct {
	Username         string    `json:"username"`
	ARN              string    `json:"arn"`
	CreateDate       time.Time `json:"createDate"`
	PasswordLastUsed time.Time `json:"passwordLastUsed"`
	UserId           string    `json:"userId"`
}

// IAMGetUserAPI defines the interface for the GetUser function.
type IAMGetUserAPI interface {
	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}

func GetUser(c context.Context, api IAMGetUserAPI, input *iam.GetUserInput) (*iam.GetUserOutput, error) {
	return api.GetUser(c, input)
}

func (aws AWS) IAMUser(cfg aws.Config) (IAMUser, error) {
	svc := iam.NewFromConfig(cfg)
	input := &iam.GetUserInput{}
	result, err := GetUser(context.TODO(), svc, input)

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

package aws

import (
	"context"
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

func (aws AWS) DescribeIAMUser(cfg aws.Config) (IAMUser, error) {
	svc := iam.New(cfg)
	req := svc.GetUserRequest(&iam.GetUserInput{})
	result, err := req.Send(context.Background())
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

func (aws AWS) DescribeIAMUsers(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListUsersRequest(&iam.ListUsersInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Users), nil
}

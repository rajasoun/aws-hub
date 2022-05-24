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

func (aws AWS) DescribeIAMUser(cfg aws.Config) (IAMUser, error) {
	// svc := iam.New(cfg)
	// req := svc.GetUserRequest(&iam.GetUserInput{})
	// result, err := req.Send(context.Background())

	client := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(context.TODO(), client, input)

	if err != nil {
		return IAMUser{}, err
	}

	lastUsed := time.Now()
	if result.Users.PasswordLastUsed != nil {
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
	client := iam.NewFromConfig(cfg)
	input := &iam.ListUsersInput{}
	result, err := ListUsers(context.TODO(), client, input)

	if err != nil {
		return 0, err
	}
	return len(result.Users), nil
	// for _, user := range result.Users {
	// 	fmt.Println(*user.UserName+" created on", *user.CreateDate)
	// }

	// svc := iam.New(cfg)
	// req := svc.ListUsersRequest(&iam.ListUsersInput{})
	// result, err := req.Send(context.Background())
	// if err != nil {
	// 	return 0, err
	// }
	//return len(result.Users), nil
}

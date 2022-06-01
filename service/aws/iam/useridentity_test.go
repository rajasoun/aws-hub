package iam

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

var testUserName string = "test@example.com"

// Mock Receiver
type MockUserIdentity struct{}

//Mock Function
type MockIAMGetUserAPIClient func(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)

// Implement AWS IAM GetUser Method with Mock Function Receiver
func (mock MockIAMGetUserAPIClient) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockUserIdentity) NewMockClient() IAMGetUserAPIClient {
	client := MockIAMGetUserAPIClient(func(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
		user := &types.User{
			Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
			CreateDate:       &time.Time{},
			UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
			UserName:         aws.String(testUserName),
			PasswordLastUsed: &time.Time{},
		}
		result := &iam.GetUserOutput{User: user}
		return result, nil
	})
	return client
}

func TestGetUserIdentity(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := MockUserIdentity{}

	cases := []struct {
		name string
		want string
	}{
		{"Check GetUserIdentity For Account", testUserName},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserIdentity(mock.NewMockClient())
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.Username, "got GetUserIdentity = %v, want = %v", got.Username, tt.want)
		})
	}
	t.Run("Check GetUserIdentity returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg) //mock.NewMockClient(emptyCfg)
		_, err := GetUserIdentity(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

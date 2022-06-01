package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

// Mock Receiver
type MockUser struct{}

//Mock Function
type MockIAMListUsersAPIClient func(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

// Implement AWS IAM ListUsers Method with Mock Function Receiver
func (mock MockIAMListUsersAPIClient) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockUser) NewMockClient() IAMListUsersAPIClient {
	client := MockIAMListUsersAPIClient(func(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
		userList := []types.User{
			{UserName: aws.String("test1@example.com")},
			{UserName: aws.String("test2@example.com")},
		}
		result := &iam.ListUsersOutput{Users: userList}
		return result, nil
	})
	return client
}

func TestGetUserCount(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := MockUser{}

	cases := []struct {
		name string
		want int
	}{
		{"Check Get User Count", 2},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserCount(mock.NewMockClient())
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.Count, "got GetUserCount = %v, want = %v", got.Count, tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg) //mock.NewMockClient(emptyCfg)
		_, err := GetUserCount(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

package iam

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// To mock AWS operations
type MockListUsers struct {
	mock.Mock
}

/**
* Mock using testify Framework
 */

// List Users Mock
func (c *MockListUsers) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(1) != nil {
		return args.Get(0).(*iam.ListUsersOutput), args.Error(1)
	}
	return args.Get(0).(*iam.ListUsersOutput), args.Error(1)
}

func TestGetUserCountViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		input   []types.User
		want    int
		wantErr error
	}{
		{
			name: "Check ListUsers via Mocking Framework",
			input: []types.User{
				{UserName: aws.String("test1@example.com")},
				{UserName: aws.String("test2@example.com")},
			},
			want:    2,
			wantErr: nil,
		},
		{
			name:    "Check ListUsers via Mocking Framework with Err",
			input:   []types.User{},
			want:    0,
			wantErr: errors.New("simulated error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := new(MockListUsers)
			expectedOutput := &iam.ListUsersOutput{Users: tt.input}
			// Inject Mock Function to be Called along with Resturn values as Parameter
			client.
				On("ListUsers", mock.Anything, mock.Anything, mock.Anything).
				Return(expectedOutput, tt.wantErr)
			got, err := GetUserCount(client)
			assert.Equal(tt.want, got.Count, "GetUserCount() = %v, want = %v", got.Count, tt.want)
			if tt.wantErr != nil {
				assert.Error(err, "GetUserCount() %v", err)
				return
			}
			assert.NoError(err, "GetUserCount() %v", err)
		})
	}
}

/**
* Mock via manual creation - Just For Reference
 */

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

func (mock MockUser) NewClient() IAMListUsersAPIClient {
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

	cases := []struct {
		name string
		want int
	}{
		{"Check Get User Count", 2},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mock := MockUser{}
			client := mock.NewClient()
			got, err := GetUserCount(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.Count, "got GetUserCount = %v, want = %v", got.Count, tt.want)
		})
	}
}

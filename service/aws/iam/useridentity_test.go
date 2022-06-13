package iam

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const testUserName = "test@example.com"

// Mock Receiver
type MockUserIdentity struct {
	mock.Mock
}

/**
* Mock using testify Framework
 */

// Get User Mock
func (c *MockUserIdentity) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*iam.GetUserOutput), args.Error(1)
}

func TestGetUserIdentityViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check ListUsers via Mocking Framework ", func(t *testing.T) {
		client := new(MockUserIdentity)
		user := &types.User{
			Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
			CreateDate:       &time.Time{},
			UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
			UserName:         aws.String(testUserName),
			PasswordLastUsed: &time.Time{},
		}
		expectedOutput := &iam.GetUserOutput{User: user}

		// Inject Mock Function to be Called along with Resturn values as Parameter
		client.
			On("GetUser", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedOutput, nil)
		want := "test@example.com"
		got, err := GetUserIdentity(client)
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(want, got.Username, "got GetUserIdentity = %v, want = %v", got.Username, want)
	})
	t.Run("Check GetUserIdentity returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg) //mock.NewMockClient(emptyCfg)
		_, err := GetUserIdentity(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

/**
* Mock via manual creation - Just For Reference
 */

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

func NewGetUserMockClient() IAMGetUserAPIClient {
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

	cases := []struct {
		name string
		want string
	}{
		{"Check GetUserIdentity For Account", testUserName},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// mock := MockUserIdentity{}
			client := NewGetUserMockClient()
			got, err := GetUserIdentity(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.Username, "got GetUserIdentity = %v, want = %v", got.Username, tt.want)
		})
	}
}

package iam

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const testUserName = "test@example.com"

/**
* Mock using testify Framework
 */

// Mock Function to AWS GetUser
// Technique: Interface Substitution
// Interface Substitution is done by mocking methods that are implemented by an interface.
// Steps:
//	1. make an object of struct
//	2. implements all methods in the interface for mocking real implementation
func (mockClient *MockClient) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	args := mockClient.Called(ctx, params, optFns)
	// On Error
	if args.Get(1) != nil {
		return args.Get(0).(*iam.GetUserOutput), args.Error(1)
	}
	// If No Error
	return args.Get(0).(*iam.GetUserOutput), nil
}

func TestGetUserIdentityViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		input   *types.User
		want    string
		wantErr error
	}{
		{
			name: "Check ListUsers via Mocking Framework",
			input: &types.User{
				Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
				CreateDate:       &time.Time{},
				UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
				UserName:         aws.String(testUserName),
				PasswordLastUsed: &time.Time{},
			},
			want:    "test@example.com",
			wantErr: nil,
		},
		{
			name:    "Check ListUsers via Mocking Framework with Err",
			input:   &types.User{},
			want:    "",
			wantErr: errors.New("simulated error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := new(MockClient)
			expectedOutput := &iam.GetUserOutput{
				User: tt.input,
			}
			// Inject Mock Function to be Called along with Resturn values as Parameter
			client.
				On("GetUser", mock.Anything, mock.Anything, mock.Anything).
				Return(expectedOutput, tt.wantErr)
			got, err := GetUserIdentity(client)
			if tt.wantErr != nil {
				assert.Error(err, " GetUserIdentity() %v", err)
				assert.Empty(got.Username, "GetUserIdentity().Username = %v", got.Username)
				return
			}
			assert.NoError(err, " GetUserIdentity() %v", err)
			assert.Equal(tt.want, got.Username, "GetUserIdentity().Username = %v, want = %v", got.Username, tt.name)
		})

	}
}

/**
* Mock via manual creation - Just For Reference
* Technique : Interface Substitution
 */

// Implement AWS IAM GetUser Method with Mock Receiver struct
func (mockReceiver MockReciever) GetUser(ctx context.Context,
	params *iam.GetUserInput,
	optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	if mockReceiver.wantErr != nil {
		return &iam.GetUserOutput{User: &types.User{}}, errors.New("simulated error")
	}
	user := &types.User{
		Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
		CreateDate:       &time.Time{},
		UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
		UserName:         aws.String(testUserName),
		PasswordLastUsed: &time.Time{},
	}
	result := &iam.GetUserOutput{User: user}
	return result, nil
}

func TestGetUserIdentity(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name    string
		client  MockReciever
		want    string
		wantErr bool
	}{
		{
			name:    "Check GetUserIdentity For Account",
			client:  MockReciever{wantErr: nil},
			want:    testUserName,
			wantErr: false,
		},
		{
			name:    "Check GetUserIdentity For Account with Err",
			client:  MockReciever{wantErr: errors.New("simulated error")},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserIdentity(tt.client)
			if tt.wantErr {
				assert.Error(err, "GetUserIdentity() %v", err)
				return
			}
			assert.NoError(err, "GetUserIdentity() %v", err)
			assert.Equal(tt.want, got.Username, "GetUserIdentity() = %v, want = %v", got.Username, tt.want)
		})
	}
}

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

/**
* Mock using testify Framework
 */

// Mock Function to AWS ListUsers
// Technique: Interface Substitution
// Interface Substitution is done by mocking methods that are implemented by an interface.
// Steps:
//	1. make an object of struct
//	2. implements all methods in the interface for mocking real implementation
func (mockClient *MockClient) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	// Mock ListUsers of AWS
	// Mocked ListUsers Function will be Called and Results Injected
	args := mockClient.Called(ctx, params, optFns)
	// On Error
	if args.Get(1) != nil {
		return args.Get(0).(*iam.ListUsersOutput), args.Error(1)
	}
	// If No Error
	return args.Get(0).(*iam.ListUsersOutput), nil
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
			client := new(MockClient)
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
* Technique : Interface Substitution
 */

// Implement AWS IAM GetUser Method with Mock Receiver struct
func (mockReceiver MockReciever) ListUsers(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	if mockReceiver.wantErr != nil {
		return &iam.ListUsersOutput{Users: []types.User{}}, errors.New("simulated error")
	}
	userList := []types.User{
		{UserName: aws.String("test1@example.com")},
		{UserName: aws.String("test2@example.com")},
	}
	result := &iam.ListUsersOutput{Users: userList}
	return result, nil
}

func TestGetUserCountviaManualMock(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name    string
		client  MockReciever
		want    int
		wantErr bool
	}{
		{
			name:    "Check GetUserCount For Account",
			client:  MockReciever{wantErr: nil},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Check GetUserCount For Account with Err",
			client:  MockReciever{wantErr: errors.New("simulated error")},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserCount(tt.client)
			if tt.wantErr {
				assert.Error(err, "GetUserCount() %v", err)
				return
			}
			assert.NoError(err, "GetUserCount() %v", err)
			assert.Equal(tt.want, got.Count, "GetUserCount() = %v, want = %v", got.Count, tt.want)
		})
	}
}

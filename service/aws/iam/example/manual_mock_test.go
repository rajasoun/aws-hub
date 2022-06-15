// Manual Mocking of AWS APIs Example
package example

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	service "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/stretchr/testify/assert"
)

/**
* Mock via manual creation - Just For Reference
* Technique : Interface Substitution
 */

// Implement AWS IAM ListAccountAliases Method with Mock Receiver struct
func (mockReceiver MockReciever) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	if mockReceiver.wantErr != nil {
		return &iam.ListAccountAliasesOutput{AccountAliases: []string{}}, errors.New(testErrMsg)
	}
	aliases := []string{testAlias}
	result := &iam.ListAccountAliasesOutput{AccountAliases: aliases}
	return result, nil
}

func TestGetAliasesviaHandMadeMock(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name    string
		client  MockReciever
		want    int
		wantErr bool
	}{
		{
			name:    "Check GetUserIdentity For Account",
			client:  MockReciever{wantErr: nil},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Check GetUserIdentity For Account with Err",
			client:  MockReciever{wantErr: errors.New(testErrMsg)},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetAliases(tt.client)
			assert.Equal(tt.want, len(got.List), "GetAliases() = %v, want = %v", len(got.List), tt.want)
			if tt.wantErr {
				assert.Error(err, "GetAliases() %v", err)
				return
			}
			assert.NoError(err, "GetAliases() %v", err)
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
		return &iam.ListUsersOutput{Users: []types.User{}}, errors.New(testErrMsg)
	}
	userList := []types.User{
		{UserName: aws.String(testUsers[0])},
		{UserName: aws.String(testUsers[1])},
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
			client:  MockReciever{wantErr: errors.New(testErrMsg)},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetUserCount(tt.client)
			if tt.wantErr {
				assert.Error(err, "GetUserCount() %v", err)
				return
			}
			assert.NoError(err, "GetUserCount() %v", err)
			assert.Equal(tt.want, got.Count, "GetUserCount() = %v, want = %v", got.Count, tt.want)
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
		return &iam.GetUserOutput{User: &types.User{}}, errors.New(testErrMsg)
	}
	user := &types.User{
		Arn:              aws.String(testARN),
		CreateDate:       &time.Time{},
		UserId:           aws.String(testUserID),
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
			client:  MockReciever{wantErr: errors.New(testErrMsg)},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetUserIdentity(tt.client)
			if tt.wantErr {
				assert.Error(err, "GetUserIdentity() %v", err)
				return
			}
			assert.NoError(err, "GetUserIdentity() %v", err)
			assert.Equal(tt.want, got.Username, "GetUserIdentity() = %v, want = %v", got.Username, tt.want)
		})
	}
}

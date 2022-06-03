package iam

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
	"github.com/stretchr/testify/assert"
)

var testUserName string = "test@example.com"

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
			mock := iammock.MockUserIdentity{}
			client := mock.NewClient()
			got, err := GetUserIdentity(client)
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

func TestGetUserIdentityViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check ListUsers via Mocking Framework ", func(t *testing.T) {
		client := new(iammock.MockClient)
		user := &types.User{
			Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
			CreateDate:       &time.Time{},
			UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
			UserName:         aws.String("test@example.com"),
			PasswordLastUsed: &time.Time{},
		}
		expectedOutput := &iam.GetUserOutput{User: user}

		client.InjectFunctionMock(client, "GetUser", expectedOutput)
		want := "test@example.com"
		got, err := GetUserIdentity(client)
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(want, got.Username, "got GetUserIdentity = %v, want = %v", got.Username, want)
	})
}

package iam

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

//Mock Interface
type MockListUsersAPI func(ctx context.Context,
	params *iam.ListUsersInput,
	optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

//ListUsers Mock
func (mock MockListUsersAPI) ListUsers(ctx context.Context,
	params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	log.Println("[usercount_test.go] (mock MockListUsersAPI) ListUsers ")
	userList := []types.User{
		{UserName: aws.String("test1@example.com")},
		{UserName: aws.String("test2@example.com")},
	}
	return &iam.ListUsersOutput{Users: userList}, nil
}

func TestListUsers(t *testing.T) {
	t.Run("Check GetListUsers", func(t *testing.T) {
		assert := assert.New(t)
		client := new(MockListUsersAPI)
		input := &iam.ListUsersInput{}
		want := 2
		got, err := ListUsers(context.TODO(), client, input)
		assert.NoError(err, "err = %v, want = nil", err)
		assert.Equal(want, len(got.Users), "got = %v , want = %v", got, want)
	})
}

func TestGetUserCount(t *testing.T) {
	t.Run("Check GetUserCount returns err with Invalid aws.Config{}", func(t *testing.T) {
		assert := assert.New(t)
		_, err := GetUserCount(aws.Config{})
		assert.Error(err, "err = %v, want = nil", err)
	})
}

package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

type mockGetUserImpl struct{}

func (mock mockGetUserImpl) GetUser(ctx context.Context,
	params *iam.GetUserInput, optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	return &iam.GetUserOutput{
		User: &types.User{
			UserName: new(string),
		},
	}, nil
}

func TestGetUser(t *testing.T) {
	t.Run("Check GetUser", func(t *testing.T) {
		assert := assert.New(t)
		api := &mockGetUserImpl{}
		input := &iam.GetUserInput{}
		want := &iam.GetUserOutput{User: &types.User{UserName: new(string)}}
		got, err := GetUser(context.TODO(), api, input)
		assert.NoError(err, "err = %v, want = nil", err)
		assert.Equal(got, want, "got = %v , want = %v", got, want)
	})
}

func TestGetUserIdentity(t *testing.T) {
	t.Run("Check GetUserIdentity returns err with aws.Config{}", func(t *testing.T) {
		assert := assert.New(t)
		_, err := GetUserIdentity(aws.Config{})
		assert.Error(err, "err = %v, want = nil", err)
	})
}

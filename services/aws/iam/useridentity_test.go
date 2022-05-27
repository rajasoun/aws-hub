package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

type GetUserImpl struct{}

func (mock GetUserImpl) GetUser(ctx context.Context,
	params *iam.GetUserInput, optFns ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	return &iam.GetUserOutput{
		User: &types.User{
			UserName: new(string),
		},
	}, nil
}

func TestGetUserIdentity(t *testing.T) {
	assert := assert.New(t)
	api := &GetUserImpl{}
	input := &iam.GetUserInput{}
	want := &iam.GetUserOutput{User: &types.User{UserName: new(string)}}
	got, err := GetUser(context.TODO(), api, input)
	assert.NoError(err, "err = %v, want = nil", err)
	assert.Equal(got, want, "got = %v , want = %v", got, want)
}

package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/stretchr/testify/assert"
)

type ListUsersImpl struct{}

func (mock ListUsersImpl) ListUsers(ctx context.Context,
	params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	userList := []types.User{
		{
			UserName: new(string),
		},
	}
	return &iam.ListUsersOutput{Users: userList}, nil
}

func TestListUsers(t *testing.T) {
	assert := assert.New(t)
	api := &ListUsersImpl{}
	input := &iam.ListUsersInput{}
	want := 1
	got, err := ListUsers(context.TODO(), api, input)
	assert.NoError(err, "err = %v, want = nil", err)
	assert.Equal(want, len(got.Users), "got = %v , want = %v", got, want)
}

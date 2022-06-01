package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
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
			mock := apiclient.MockUserIdentity{}
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

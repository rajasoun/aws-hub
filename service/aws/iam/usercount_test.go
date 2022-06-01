package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
	"github.com/stretchr/testify/assert"
)

func TestGetUserCount(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name string
		want int
	}{
		{"Check Get User Count", 2},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mock := apiclient.MockUser{}
			client := mock.NewClient()
			got, err := GetUserCount(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.Count, "got GetUserCount = %v, want = %v", got.Count, tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg) //mock.NewMockClient(emptyCfg)
		_, err := GetUserCount(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

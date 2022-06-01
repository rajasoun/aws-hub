package iammock

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"

	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
)

func TestGetAliases(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := MockAccountAliases{}

	cases := []struct {
		name string
		want string
	}{
		{"Check Get Aliases", testAlias},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := mock.NewMockClient()
			got, err := hubIAM.GetAliases(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetUserCount = %v, want = %v", got.List[0], tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg)
		_, err := hubIAM.GetAliases(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

package iammock

import (
	"testing"

	"github.com/stretchr/testify/assert"

	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
)

func TestGetAliases(t *testing.T) {
	var testAlias string = "aws-test-account-alias"
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
			assert.Equal(tt.want, got.List[0], "got GetAliases = %v, want = %v", got.List[0], tt.want)
		})
	}
}

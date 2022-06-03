package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/stretchr/testify/assert"

	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
)

func TestGetAliases(t *testing.T) {
	var testAlias string = "aws-test-account-alias"
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name string
		want string
	}{
		{"Check Get Aliases", testAlias},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mock := iammock.MockAccountAliases{}
			client := mock.NewClient()
			got, err := GetAliases(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetUserCount = %v, want = %v", got.List[0], tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg) //mock.NewMockClient(emptyCfg)
		_, err := GetAliases(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

func TestListAccountAliasesViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check ListAccountAliases via Mocking Framework ", func(t *testing.T) {
		client := new(iammock.MockClient)
		var testAlias string = "aws-test-account-alias"
		aliases := []string{testAlias}
		expectedOutput := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		client.InjectFunctionMock(client, "ListAccountAliases", expectedOutput)
		got, err := GetAliases(client)
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(testAlias, got.List[0], "got GetAliases = %v, want = %v", got.List[0], testAlias)
	})
}
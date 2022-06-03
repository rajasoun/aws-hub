package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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

// Client provides the API client to mock AWS operations
type MockIt struct {
	mock.Mock
}

// List Account Aliases Mock
func (c *MockIt) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	args := c.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*iam.ListAccountAliasesOutput), args.Error(1)
}

func TestListAccountAliasesViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check ListAccountAliases via Mocking Framework ", func(t *testing.T) {
		//client := new(iammock.MockClient)
		client := new(MockIt)
		var testAlias string = "aws-test-account-alias"
		aliases := []string{testAlias}
		expectedOutput := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		//client.InjectFunctionMock(client, "ListAccountAliases", expectedOutput)
		client.
			On("ListAccountAliases", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedOutput, nil)
		got, err := GetAliases(client)
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(testAlias, got.List[0], "got GetAliases = %v, want = %v", got.List[0], testAlias)
	})
}

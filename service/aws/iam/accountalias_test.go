package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// To mock AWS operations
type MockAccountAliases struct {
	mock.Mock
}

/**
* Mock using testify Framework
 */

// List Account Aliases Mock
func (mockFunc *MockAccountAliases) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	// Call Mock Function call
	// The Function to be Called and Result will be Injected in Test
	result := mockFunc.Called(ctx, params, optFns)
	//Return Result On Error
	if result.Get(0) == nil {
		return nil, result.Error(1)
	}
	//Return Result If No Error
	return result.Get(0).(*iam.ListAccountAliasesOutput), result.Error(1)
}

var testAlias string = "aws-test-account-alias"

func TestListAccountAliasesViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check ListAccountAliases", func(t *testing.T) {
		//client := new(iammock.MockClient)
		client := new(MockAccountAliases)
		aliases := []string{testAlias}
		expectedOutput := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		// Inject Mock Function to be Called along with Resturn values as Parameter
		client.
			On("ListAccountAliases", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedOutput, nil)
		got, err := GetAliases(client)
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(testAlias, got.List[0], "got GetAliases = %v, want = %v", got.List[0], testAlias)
	})
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg)
		_, err := GetAliases(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

/**
* Mock via manual creation - Just For Reference
 */

//Mock Function
type MockListAccountAliasesAPIClient func(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)

// Implement AWS IAM ListAccountAliases Interface with mock reciever
func (mock MockListAccountAliasesAPIClient) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	return mock(ctx, params, optFns...)
}

func (mock MockAccountAliases) NewClient() IAMListAccountAliasesAPIClient {
	fn := func(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
		aliases := []string{testAlias}
		result := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		return result, nil
	}
	client := MockListAccountAliasesAPIClient(fn)
	return client
}
func TestGetAliasesviaHandMadeMock(t *testing.T) {
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
			mock := MockAccountAliases{}
			client := mock.NewClient()
			got, err := GetAliases(client)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetUserCount = %v, want = %v", got.List[0], tt.want)
		})
	}
}

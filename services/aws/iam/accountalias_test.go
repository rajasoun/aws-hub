package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
)

var testAlias string = "aws-test-account-alias"

//Mock Function
type MockIAMListAccountAliasesAPI func(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)

// AWS IAM ListAccountAliases Method with Mock Function Receiver
func (mock MockIAMListAccountAliasesAPI) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	return mock(ctx, params, optFns...)
}

func mockIAMListAccountAliasesAPIOutput() IAMListAccountAliasesAPI {
	return MockIAMListAccountAliasesAPI(func(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
		aliases := []string{testAlias}
		result := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		return result, nil
	})
}

func TestGetAliases(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	cases := []struct {
		name   string
		client func() IAMListAccountAliasesAPI
		want   string
	}{
		{
			name: "Check Get Aliases",
			client: func() IAMListAccountAliasesAPI {
				return mockIAMListAccountAliasesAPIOutput()
			},
			want: testAlias,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAliases(tt.client())
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetUserCount = %v, want = %v", got.List[0], tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg)
		_, err := GetAliases(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

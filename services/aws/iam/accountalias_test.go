package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
)

//type ListAccountAliasesImpl struct{}

type mockListAccountAliases struct{}

func (dt mockListAccountAliases) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	aliases := []string{
		"aws-docs-account-alias1",
		"aws-docs-account-alias2",
	}
	output := &iam.ListAccountAliasesOutput{
		AccountAliases: aliases,
	}
	return output, nil
}

func TestGetAccountAliases(t *testing.T) {
	t.Run("Check GetAccountAliases", func(t *testing.T) {
		assert := assert.New(t)
		client := &mockListAccountAliases{}
		input := &iam.ListAccountAliasesInput{}
		want := 2
		got, err := GetAccountAliases(context.Background(), client, input)
		assert.NoError(err, "err = %v, want = nil", err)
		assert.Equal(want, len(got.AccountAliases), "got = %v , want = %v", got, want)
	})
}

func TestGetAliases(t *testing.T) {
	t.Run("Check GetAliases returns err with aws.Config{}", func(t *testing.T) {
		assert := assert.New(t)
		_, err := GetAliases(aws.Config{})
		assert.Error(err, "err = %v, want = nil", err)
	})
}

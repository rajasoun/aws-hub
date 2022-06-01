package iammock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

// Mock Receiver
type MockAccountAliases struct{}

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

func (mock MockAccountAliases) NewMockClient() IAMListAccountAliasesAPIClient {
	client := MockListAccountAliasesAPIClient(func(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
		aliases := []string{testAlias}
		result := &iam.ListAccountAliasesOutput{
			AccountAliases: aliases,
		}
		return result, nil
	})
	return client
}

package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type Aliases struct {
	List []string `json:"list"`
}

//Interface for Amazon IAM ListAccountAliases API
//This will enable TDD using mocking
type IAMListAccountAliasesAPIClient interface {
	iam.ListAccountAliasesAPIClient // Only for Refernce to Actual Client
	ListAccountAliases(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)
}

// GetAccountAliases retrieves the aliases for your AWS Identity and Access Management (IAM) account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a ListAccountAliasesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ListAccountAliases.
func GetAliases(client IAMListAccountAliasesAPIClient) (Aliases, error) {
	var ctx context.Context = context.TODO()
	input := &iam.ListAccountAliasesInput{}
	result, err := client.ListAccountAliases(ctx, input)
	if err != nil {
		log.Println("Got an error retrieving account aliases")
		return Aliases{List: []string{}}, err
	}
	return Aliases{List: result.AccountAliases}, nil
}

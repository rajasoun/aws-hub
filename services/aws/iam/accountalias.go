package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type Aliases struct {
	List []string `json:"list"`
}

// Interface wraps up the underlying AWS Function
// This will enable TDD using mocking the wrapped function
type ListAccountAliasesAPI interface {
	ListAccountAliases(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)
}

// GetAccountAliases retrieves the aliases for your AWS Identity and Access Management (IAM) account.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If successful, a ListAccountAliasesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ListAccountAliases.
func GetAccountAliases(c context.Context, api ListAccountAliasesAPI,
	input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
	return api.ListAccountAliases(c, input)

}

func GetAliases(cfg aws.Config) (Aliases, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.ListAccountAliasesInput{}
	result, err := GetAccountAliases(context.TODO(), client, input)
	if err != nil {
		log.Println("Got an error retrieving account aliases")
		return Aliases{List: []string{}}, err
	}
	return Aliases{List: result.AccountAliases}, nil
}

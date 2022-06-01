package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	clientAPI "github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
)

type Aliases struct {
	List []string `json:"list"`
}

// GetAccountAliases retrieves the aliases for your AWS Identity and Access Management (IAM) account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a ListAccountAliasesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to ListAccountAliases.
func GetAliases(client clientAPI.IAMListAccountAliasesAPIClient) (Aliases, error) {
	var ctx context.Context = context.TODO()
	input := &iam.ListAccountAliasesInput{}
	result, err := client.ListAccountAliases(ctx, input)
	if err != nil {
		log.Println("Got an error retrieving account aliases")
		return Aliases{List: []string{}}, err
	}
	return Aliases{List: result.AccountAliases}, nil
}

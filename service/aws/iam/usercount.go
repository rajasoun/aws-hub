package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/apiclient"
)

type UserList struct {
	Count int `json:"usercount"`
}

// GetUserCount retrieves the user accounty for an AWS account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a Users object containing the count and nil.
//     Otherwise, nil and an error from the call.
func GetUserCount(client apiclient.IAMListUsersAPIClient) (UserList, error) {
	var ctx context.Context = context.TODO()
	input := &iam.ListUsersInput{}
	result, err := client.ListUsers(ctx, input)
	if err != nil {
		return UserList{Count: 0}, err
	}
	return UserList{Count: len(result.Users)}, nil
}

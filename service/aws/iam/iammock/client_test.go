package iammock

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/stretchr/testify/mock"
)

func TestListAccountAliases(t *testing.T) {
	client := new(MockClient)
	methodName := "ListAccountAliases"
	output := iam.ListAccountAliasesOutput{}
	client.
		On(methodName, mock.Anything, mock.Anything, mock.Anything).
		Return(&output, nil)
	fn := func(c hubIAM.IAMAPIClient) {
		_, _ = c.ListAccountAliases(context.Background(), nil, nil)
	}
	fn(client)
	client.AssertExpectations(t)
	// client.
	// 	On(methodName, mock.Anything, mock.Anything, mock.Anything).
	// 	Return(&output, errors.New("Mock Error"))

	// fn(client)
	// client.AssertExpectations(t)
}

// func TestListUsers(t *testing.T) {
// 	client := new(MockClient)
// 	methodName := "ListUsers"
// 	output := iam.ListUsersOutput{}
// 	client.
// 		On(methodName, mock.Anything, mock.Anything, mock.Anything).
// 		Return(&output, nil)
// 	fn := func(c hubIAM.IamClient) {
// 		_, _ = c.ListUsers(context.Background(), nil, nil)
// 	}
// 	fn(client)
// 	client.AssertExpectations(t)
// }

// func TestGetUser(t *testing.T) {
// 	client := new(MockClient)
// 	methodName := "GetUser"
// 	output := iam.GetUserOutput{}
// 	client.
// 		On(methodName, mock.Anything, mock.Anything, mock.Anything).
// 		Return(&output, nil)
// 	fn := func(c hubIAM.IamClient) {
// 		_, _ = c.GetUser(context.Background(), nil, nil)
// 	}
// 	fn(client)
// 	client.AssertExpectations(t)
// }

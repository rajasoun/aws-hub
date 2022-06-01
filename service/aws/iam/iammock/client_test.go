package iammock

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"

	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
)

// import (
// 	"context"
// 	"testing"

// 	"github.com/aws/aws-sdk-go-v2/service/iam"
// 	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
// 	"github.com/stretchr/testify/mock"
// )

// func TestListAccountAliases(t *testing.T) {
// 	fn := func(c hubIAM.IamClient) {
// 		_, _ = c.ListAccountAliases(context.Background(), nil, nil)
// 	}
// 	client := new(MockClient)
// 	methodName := "ListAccountAliases"
// 	output := iam.ListAccountAliasesOutput{}
// 	client.
// 		On(methodName, mock.Anything, mock.Anything, mock.Anything).
// 		Return(&output, nil)

// 	fn(client)
// 	client.AssertExpectations(t)
// 	// client.
// 	// 	On(methodName, mock.Anything, mock.Anything, mock.Anything).
// 	// 	Return(&output, errors.New("Mock Error"))

// 	// fn(client)
// 	// client.AssertExpectations(t)
// }

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

func TestGetAliases(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	mock := MockAccountAliases{}

	cases := []struct {
		name string
		want string
	}{
		{"Check Get Aliases", testAlias},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hubIAM.GetAliases(mock.NewMockClient())
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetUserCount = %v, want = %v", got.List[0], tt.want)
		})
	}
	t.Run("Check GetUserCount returns err with Empty aws.Config{}", func(t *testing.T) {
		emptyCfg := aws.Config{}
		noOpClient := iam.NewFromConfig(emptyCfg)
		_, err := hubIAM.GetAliases(noOpClient)
		assert.Error(err, "err = %v, want = nil", err)
	})
}

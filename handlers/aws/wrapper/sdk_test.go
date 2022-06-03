package wrapper

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
	"github.com/stretchr/testify/assert"
)

var client = new(iammock.MockClient)

func TestAccountExecute(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check Account Execute", func(t *testing.T) {
		injectFunction := "ListAccountAliases"
		expectedOutput := &iam.ListAccountAliasesOutput{
			AccountAliases: []string{"aws-test-account-alias"},
		}
		want := "aws-test-account-alias"
		client.InjectFunctionMock(client, injectFunction, expectedOutput)
		receiver := New(&Account{
			client:   client,
			response: hubIAM.Aliases{},
		})
		err := receiver.Wrapper.Execute()
		got := receiver.Wrapper.response
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(want, got.List[0], "got GetAliases = %v, want = %v", got.List[0], want)
	})

}

func TestUserCountExecute(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Get User Count Execute", func(t *testing.T) {
		injectFunction := "ListUsers"
		userList := []types.User{
			{UserName: aws.String("test1@example.com")},
			{UserName: aws.String("test2@example.com")},
		}
		expectedOutput := &iam.ListUsersOutput{Users: userList}
		want := 2
		client.InjectFunctionMock(client, injectFunction, expectedOutput)
		receiver := New(&UserCount{
			client:   client,
			response: hubIAM.UserList{},
		})
		err := receiver.Wrapper.Execute()
		got := receiver.Wrapper.response
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(want, got.Count, "got GetAliases = %v, want = %v", got.Count, want)
	})

}

func TestUserExecute(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	t.Run("Check Get User Identity Execute", func(t *testing.T) {
		injectFunction := "GetUser"
		user := &types.User{
			Arn:              aws.String("arn:aws:iam::000123456789:user/test@example.com"),
			CreateDate:       &time.Time{},
			UserId:           aws.String("ABCDEFGHIJKLMNOPQRST"),
			UserName:         aws.String("test@example.com"),
			PasswordLastUsed: &time.Time{},
		}
		expectedOutput := &iam.GetUserOutput{User: user}
		want := "test@example.com"
		client.InjectFunctionMock(client, injectFunction, expectedOutput)
		receiver := New(&UserIdentity{
			client:   client,
			response: hubIAM.User{},
		})
		err := receiver.Wrapper.Execute()
		got := receiver.Wrapper.response
		assert.NoError(err, "expect no error, got %v", err)
		assert.Equal(want, got.Username, "got GetAliases = %v, want = %v", got.Username, want)
	})
}

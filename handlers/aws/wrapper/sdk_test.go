package wrapper

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
	"github.com/stretchr/testify/assert"
)

func TestAccountExecute(t *testing.T) {
	assert := assert.New(t)
	client := new(iammock.MockClient)
	t.Parallel()
	cases := []struct {
		name           string
		injectFunction string
		expectedOutput interface{}
		want           string
	}{
		{
			name:           "Check ListAccountAliases via Mocking Framework",
			injectFunction: "ListAccountAliases",
			expectedOutput: &iam.ListAccountAliasesOutput{AccountAliases: []string{"aws-test-account-alias"}},
			want:           "aws-test-account-alias",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client.InjectFunctionMock(client, tt.injectFunction, tt.expectedOutput)
			receiver := NewAccount(client)
			err := receiver.Wrapper.Execute()
			got := receiver.Wrapper.alias
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetAliases = %v, want = %v", got.List[0], tt.want)
		})
	}
}

func NewAccount(client *iammock.MockClient) *SDK[*Account] {
	receiver := New(&Account{
		client: client,
		alias:  hubIAM.Aliases{},
	})
	return receiver
}

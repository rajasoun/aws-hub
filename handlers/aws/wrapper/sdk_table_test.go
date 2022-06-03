package wrapper

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
	"github.com/stretchr/testify/assert"
)

// Modularization of Above Tests using testsuites
// ToDo: Generic Variable Declaration

func TestAccountExecuteWithTesSuites(t *testing.T) {
	var clientMock = new(iammock.MockClient)
	assert := assert.New(t)

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
			clientMock.InjectFunctionMock(clientMock, tt.injectFunction, tt.expectedOutput)
			account := Account{}
			receiver := account.NewReceiver(clientMock)
			err := receiver.Wrapper.Execute()
			got := receiver.Wrapper.response
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetAliases = %v, want = %v", got.List[0], tt.want)
		})
	}
}

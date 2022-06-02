package wrapper

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/service/aws/iam/iammock"
	"github.com/stretchr/testify/assert"
)

var client = new(iammock.MockClient)

func TestAccountExecute(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	cases := []struct {
		name           string
		injectFunction string
		expectedOutput interface{}
		want           string
		receiver       Account
	}{
		{
			name:           "Check ListAccountAliases via Mocking Framework",
			injectFunction: "ListAccountAliases",
			expectedOutput: &iam.ListAccountAliasesOutput{AccountAliases: []string{"aws-test-account-alias"}},
			want:           "aws-test-account-alias",
			receiver:       Account{client: client},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sdk := new(SDK)
			client.InjectFunctionMock(client, tt.injectFunction, tt.expectedOutput)
			assert.NoError(nil, "expect no error, got %v", nil)

			receiver := &tt.receiver //Account{client: client}
			err := sdk.ExecuteAPI(receiver)
			got := receiver.alias

			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.want, got.List[0], "got GetAliases = %v, want = %v", got.List[0], tt.want)
			//assert.Equal(tt.want, got, "got GetAliases = %v, want = %v", got, tt.want)
		})
	}
}

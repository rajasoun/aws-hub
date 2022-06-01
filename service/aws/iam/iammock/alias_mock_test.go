package iammock

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	hubIAM "github.com/rajasoun/aws-hub/service/aws/iam"
	"github.com/stretchr/testify/assert"
)

func TestListAccountAliases(t *testing.T) {
	assert := assert.New(t)
	client := new(MockClient)
	var testAlias string = "aws-test-account-alias"
	aliases := []string{testAlias}
	expectedOutput := &iam.ListAccountAliasesOutput{
		AccountAliases: aliases,
	}
	client.InjectFunctionMock(client, "ListAccountAliases", expectedOutput)
	got, err := hubIAM.GetAliases(client)
	assert.NoError(err, "expect no error, got %v", err)
	assert.Equal(testAlias, got.List[0], "got GetAliases = %v, want = %v", got.List[0], testAlias)
}

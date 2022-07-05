// Integration test for the AWS IAM APIs

package integration_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/provider/credential"
	hubiam "github.com/rajasoun/aws-hub/service/external/aws/iam"
	"github.com/stretchr/testify/assert"
)

func New() (*iam.Client, error) {
	cfgLoader := credential.New()
	cfg, err := cfgLoader.LoadDefaultConfig()
	if err != nil {
		return nil, err
	}

	client := iam.NewFromConfig(cfg)
	return client, nil
}

func setUp(t *testing.T) bool {
	var outputToFile = false
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}

	if os.Getenv("OUTPUT_TO_FILE") != "" {
		outputToFile = true
	}
	return outputToFile
}

func TestIAMAccountAlias(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	outputToFile := setUp(t)
	client, err := New()
	assert.NoError(err, "IAM Client New() = %v", err)
	filePath := "testdata/iam/alias.json"

	t.Run("Check Account Alias", func(t *testing.T) {
		got, err := hubiam.GetAliases(client)
		assert.NoError(err, "GetAliases() = %v", err)

		if outputToFile {
			jsonWriteErr := ToJSONFile(filePath, got)
			assert.NoError(jsonWriteErr, "ToJSONFile() = %v", jsonWriteErr)
		}

		want := hubiam.Aliases{}
		jsonReadErr := FromJSONFile(filePath, &want)
		assert.NoError(jsonReadErr, "FromJSONFile() = %v", jsonReadErr)

		assert.Equal(want, got, "GetAliases() = %v, want = %v", got, want)

	})
}

func TestIAMUsersCount(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	outputToFile := setUp(t)
	client, err := New()
	assert.NoError(err, "IAM Client New() = %v", err)
	filePath := "testdata/iam/users.json"

	t.Run("Check Users Count", func(t *testing.T) {
		got, err := hubiam.GetUserCount(client)
		assert.NoError(err, "GetUserCount() = %v", err)

		if outputToFile {
			jsonWriteErr := ToJSONFile(filePath, got)
			assert.NoError(jsonWriteErr, "ToJSONFile() = %v", jsonWriteErr)
		}

		want := hubiam.UserList{}
		jsonReadErr := FromJSONFile(filePath, &want)
		assert.NoError(jsonReadErr, "FromJSONFile() = %v", jsonReadErr)

		assert.Equal(want, got, "GetUserCount() = %v, want = %v", got, want)

	})
}

func TestIAMUserIdentity(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	outputToFile := setUp(t)
	client, err := New()
	assert.NoError(err, "IAM Client New() = %v", err)
	filePath := "testdata/iam/identity.json"

	t.Run("Check User Identity", func(t *testing.T) {
		got, err := hubiam.GetUserIdentity(client)
		assert.NoError(err, "GetUserIdentity() = %v", err)

		if outputToFile {
			jsonWriteErr := ToJSONFile(filePath, got)
			assert.NoError(jsonWriteErr, "ToJSONFile() = %v", jsonWriteErr)
		}

		want := hubiam.User{}
		jsonReadErr := FromJSONFile(filePath, &want)
		assert.NoError(jsonReadErr, "FromJSONFile() = %v", jsonReadErr)

		assert.Equal(want.ARN, got.ARN, "GetUserIdentity() ARN = %v, want = %v", got.ARN, want.ARN)
		assert.Equal(want.CreateDate, got.CreateDate, "GetUserIdentity() CreateDate = %v, want = %v", got.CreateDate, want.CreateDate)
		assert.Equal(want.UserID, got.UserID, "GetUserIdentity() UserID = %v, want = %v", got.UserID, want.UserID)
		assert.Equal(want.Username, got.Username, "GetUserIdentity() Username = %v, want = %v", got.Username, want.Username)

	})
}

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

func TestIAM(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}

	assert := assert.New(t)
	t.Parallel()

	// Load Default Configuration
	cfgLoader := credential.New()
	cfg, err := cfgLoader.LoadDefaultConfig()
	assert.NoError(err, "LoadDefaultConfig() = %v", err)
	// Create Client
	client := iam.NewFromConfig(cfg)

	t.Run("Check Account Alias", func(t *testing.T) {
		got, err := hubiam.GetAliases(client)
		assert.NoError(err, "GetAliases() = %v", err)
		want := hubiam.Aliases{}
		assert.Equal(want, got.List[0], "GetAliases() = %v, want = %v", got.List[0], want)
	})
	t.Run("Check Users Count", func(t *testing.T) {
		got, err := hubiam.GetUserCount(client)
		assert.NoError(err, "GetUserCount() = %v", err)
		want := hubiam.UserList{}
		assert.Equal(want, got.Count, "GetUserCount() = %v, want = %v", got.Count, want)
	})
	t.Run("Check User Identity", func(t *testing.T) {
		got, err := hubiam.GetUserIdentity(client)
		assert.NoError(err, "GetUserIdentity() = %v", err)
		want := hubiam.User{}
		assert.Equal(want, got, "GetUserIdentity() = %v, want = %v", got, want)
	})
}

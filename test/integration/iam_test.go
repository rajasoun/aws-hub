// Integration test for the AWS IAM APIs

package integration_test

import (
	"encoding/json"
	"io/ioutil"
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
	tests := []struct {
		name string
	}{
		{
			name: "Check Account Alias",
		},
		{
			name: "Check Users Count",
		},
		{
			name: "Check User Identity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hubiam.GetAliases(client)
			assert.NoError(err, "GetAliases() = %v", err)
			want := hubiam.Aliases{}
			loadTestData("testdata/iam/alias.json", &want)
			assert.Equal(want, got, "GetAliases() = %v, want = %v", got, want)
		})
		t.Run(tt.name, func(t *testing.T) {
			got, err := hubiam.GetUserCount(client)
			assert.NoError(err, "GetUserCount() = %v", err)
			want := hubiam.UserList{}
			loadTestData("testdata/iam/users.json", &want)
			assert.Equal(want, got, "GetUserCount() = %v, want = %v", got, want)
		})
		t.Run(tt.name, func(t *testing.T) {
			got, err := hubiam.GetUserIdentity(client)
			assert.NoError(err, "GetUserIdentity() = %v", err)
			want := hubiam.User{}
			loadTestData("testdata/iam/identity.json", &want)
			assert.Equal(want.ARN, got.ARN, "GetUserIdentity() ARN = %v, want = %v", got.ARN, want.ARN)
			assert.Equal(want.CreateDate, got.CreateDate, "GetUserIdentity() CreateDate = %v, want = %v", got.CreateDate, want.CreateDate)
			assert.Equal(want.UserID, got.UserID, "GetUserIdentity() UserID = %v, want = %v", got.UserID, want.UserID)
			assert.Equal(want.Username, got.Username, "GetUserIdentity() Username = %v, want = %v", got.Username, want.Username)
		})
	}
}

func loadTestData(filePath string, ds interface{}) {
	file, _ := ioutil.ReadFile(filePath)
	json.Unmarshal(file, &ds)
}

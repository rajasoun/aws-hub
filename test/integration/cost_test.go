// Integration test for the Cost Explorer APIs

package integration_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/rajasoun/aws-hub/provider/credential"

	"github.com/stretchr/testify/assert"

	"github.com/rajasoun/aws-hub/service/external/aws/cost"
)

func TestCurrentBill(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}

	assert := assert.New(t)
	t.Parallel()
	t.Run("Check Current Bill", func(t *testing.T) {
		// Load Default Configuration
		cfgLoader := credential.New()
		cfg, err := cfgLoader.LoadDefaultConfig()
		assert.NoError(err, "LoadDefaultConfig() = %v", err)

		// Create Client From Configuration
		client := costexplorer.NewFromConfig(cfg)

		// Execute API
		got, err := cost.CurrentBill(client)

		// Assert
		assert.NoError(err, "CurrentBill() = %v", err)
		assert.NotZero(got.Total, "GetCost() = %v, want %v", got.Total)
	})
}

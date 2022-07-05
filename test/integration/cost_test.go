// Integration test for the Cost Explorer APIs

package integration_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/rajasoun/aws-hub/provider/credential"

	"github.com/stretchr/testify/assert"

	"github.com/rajasoun/aws-hub/service/external/aws/cost"
	"github.com/rajasoun/aws-hub/service/external/aws/cost/model"
)

func TestCurrentBill(t *testing.T) {
	var outputToFile = false
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}

	if os.Getenv("OUTPUT_TO_FILE") != "" {
		outputToFile = true
	}

	assert := assert.New(t)
	t.Parallel()
	// Load Default Configuration
	cfgLoader := credential.New()
	cfg, err := cfgLoader.LoadDefaultConfig()
	assert.NoError(err, "LoadDefaultConfig() = %v", err)

	// Create Client From Configuration
	client := costexplorer.NewFromConfig(cfg)

	tests := []struct {
		name     string
		filePath string
	}{
		{
			name:     "Check Current Bill",
			filePath: "testdata/cost/bill.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Execute API
			got, err := cost.CurrentBill(client)

			if outputToFile {
				jsonWriteErr := ToJSONFile(tt.filePath, got)
				assert.NoError(jsonWriteErr, "ToJSONFile() = %v", jsonWriteErr)
			}

			want := model.Bill{}
			jsonReadErr := FromJSONFile(tt.filePath, &want)
			assert.NoError(jsonReadErr, "FromJSONFile() = %v", jsonReadErr)

			assert.Equal(want, got, "CurrentBill() = %v, want = %v", got, want)

			assert.NoError(err, "CurrentBill() = %v", err)
			assert.NotZero(got.Total, "GetCost() = %v, want %v", got.Total)
		})
	}
}

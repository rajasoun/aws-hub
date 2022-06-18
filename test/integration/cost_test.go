// Integration test for the Cost Explorer APIs

package cost

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	cfg, _ := credential.New().LoadDefaultConfig()
	type args struct {
		client cost.GetCostAndUsageAPI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Check CurrentBill With Valid Config",
			args: args{
				client: costexplorer.NewFromConfig(cfg),
			},
			wantErr: false,
		},
		{
			name: "Check CurrentBill With Empty Config",
			args: args{
				client: costexplorer.NewFromConfig(aws.Config{}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Fail()
			got, err := cost.CurrentBill(tt.args.client)
			if tt.wantErr {
				assert.Error(err, "GetCost() error = %v", err)
				assert.Zero(got.Total, "GetCost() = %v, want %v", got.Total)
				return
			}
			assert.NotZero(got.Total, "GetCost() = %v, want %v", got.Total)
			assert.NoError(err, "GetCost() error = %v", err)
		})
	}
}

package cost

import (
	"testing"

	awsutil "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	model "github.com/rajasoun/aws-hub/service/external/aws/cost/model"
	"github.com/stretchr/testify/assert"
)

func TestGroupByMetrics(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type args struct {
		res types.ResultByTime
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []model.Group
	}{
		{
			name: "Check GroupBMetrics for Empty Result",
			args: args{
				res: types.ResultByTime{Groups: []types.Group{}},
			},
			want:  "USD",
			want1: []model.Group{},
		},
		{
			name: "Check GroupBMetrics for Non Empty Result",
			args: args{
				res: types.ResultByTime{Groups: []types.Group{
					{
						Keys: []string{"BlendedCost"},
						Metrics: map[string]types.MetricValue{
							"BlendedCost": {
								Amount: awsutil.String("100"),
								Unit:   awsutil.String("USD"),
							},
						},
					},
				}},
			},
			want:  "USD",
			want1: []model.Group{{Key: "BlendedCost", Amount: 100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := groupByMetrics(tt.args.res)
			assert.Equal(tt.want, got, "groupByMetrics() got = %v, want %v", got, tt.want)
			assert.Equal(tt.want1, got1, "groupByMetrics() got = %v, want %v", got1, tt.want1)
		})
	}
}

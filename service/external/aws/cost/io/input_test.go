package io

import (
	"testing"
	"time"

	awsutil "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/stretchr/testify/assert"
)

func TestTimePeriod(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		years  int
		months int
		days   int
	}
	tests := []struct {
		name string
		args args
		want *types.DateInterval
	}{
		{
			name: "Check TimePeriod For Last 6 Months",
			args: args{
				years:  0,
				months: -6,
				days:   0,
			},
			want: &types.DateInterval{
				Start: awsutil.String(time.Now().AddDate(0, -6, 0).Format(formatLayout)),
				End:   awsutil.String(time.Now().Format(formatLayout)),
			},
		},
		{
			name: "Check TimePeriod For Last 1 Year",
			args: args{
				years:  -1,
				months: 0,
				days:   0,
			},
			want: &types.DateInterval{
				Start: awsutil.String(time.Now().AddDate(-1, 0, 0).Format(formatLayout)),
				End:   awsutil.String(time.Now().Format(formatLayout)),
			},
		},
		{
			name: "Check TimePeriod For Last 30 Days",
			args: args{
				years:  0,
				months: 0,
				days:   -30,
			},
			want: &types.DateInterval{
				Start: awsutil.String(time.Now().AddDate(0, 0, -30).Format(formatLayout)),
				End:   awsutil.String(time.Now().Format(formatLayout)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timePeriod(tt.args.years, tt.args.months, tt.args.days)
			assert.Equal(tt.want, got, "timePeriod() = %v, want %v", got, tt.want)
		})
	}
}

func TestCostAndUsageInput(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type want struct {
		Granularity         types.Granularity
		Metrics             []string
		GroupDefinitionSize int
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Check CostAndUsageInput",
			want: want{
				Granularity:         types.Granularity("MONTHLY"),
				Metrics:             []string{"BlendedCost"},
				GroupDefinitionSize: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CostAndUsageInput()
			assert.Equal(tt.want.Granularity, got.Granularity, "CostAndUsageInput() = %v, want %v", got.Granularity, tt.want.Granularity)
			assert.Equal(tt.want.Metrics, got.Metrics, "CostAndUsageInput() = %v, want %v", got.Granularity, tt.want.Metrics)
			assert.Equal(tt.want.GroupDefinitionSize, len(got.GroupBy), "CostAndUsageInput() = %v, want %v", len(got.GroupBy), tt.want.GroupDefinitionSize)
		})
	}
}

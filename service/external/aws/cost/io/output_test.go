package io

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"

	awsutil "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
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

func TestAppendToList(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		costs  []model.Cost
		start  time.Time
		end    time.Time
		unit   string
		groups []model.Group
	}
	tests := []struct {
		name string
		args args
		want []model.Cost
	}{
		{
			name: "Check AppendToList with Empty Cost",
			args: args{
				costs:  make([]model.Cost, 0),
				start:  time.Time{},
				end:    time.Time{},
				unit:   "USD",
				groups: []model.Group{},
			},
			want: []model.Cost{
				{
					Start:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					End:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					Unit:   "USD",
					Groups: []model.Group{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendToList(tt.args.costs, tt.args.start, tt.args.end, tt.args.unit, tt.args.groups)
			assert.Equal(tt.want, got, "appendToList() = %v, want %v", got, tt.want)
			assert.Equal(1, len(got), "appendToList() = %v, want %v", len(got), 1)
		})
	}
}

func TestCostHistory(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		result *costexplorer.GetCostAndUsageOutput
	}
	tests := []struct {
		name string
		args args
		want []model.Cost
	}{
		{
			name: "Check CostHistory with Empty result",
			args: args{
				result: &costexplorer.GetCostAndUsageOutput{},
			},
			want: []model.Cost{},
		},
		{
			name: "Check CostHistory with Non Empty result",
			args: args{
				result: &costexplorer.GetCostAndUsageOutput{},
			},
			want: []model.Cost{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := costHistory(tt.args.result)
			assert.Equal(tt.want, got, "costHistory() = %v, want %v", got, tt.want)
		})
	}
}

func TestCurrentBill(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		costHistory []model.Cost
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Check Current Bill",
			args: args{
				costHistory: []model.Cost{
					{
						Unit:   "USD",
						Groups: []model.Group{{Key: "Total", Amount: 0}},
					}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := currentBill(tt.args.costHistory)
			assert.Equal(tt.want, got, "currentBill() = %v, want %v", got, tt.want)
		})
	}
}

func TestGetCostHistoryAndBill(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	jsonData := jsonTestData()

	type args struct {
		result *costexplorer.GetCostAndUsageOutput
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Check GetCostHistoryAndBill",
			args: args{
				result: &costexplorer.GetCostAndUsageOutput{},
			},
			want: 0,
		},
		{
			name: "Check GetCostHistoryAndBill",
			args: args{
				result: &jsonData,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := TotalBillAndHistory(tt.args.result)
			assert.Equal(tt.want, got1, "GetCostHistoryAndBill() got = %v, want %v", got1, tt.want)
		})
	}
}

func jsonTestData() costexplorer.GetCostAndUsageOutput {
	jsonData := costexplorer.GetCostAndUsageOutput{}
	file, _ := ioutil.ReadFile("testdata/test.json")
	err := json.Unmarshal(file, &jsonData)
	if err != nil {
		log.Printf("Test Data UnMarshall Err = %v", err)
		return costexplorer.GetCostAndUsageOutput{}
	}
	return jsonData
}

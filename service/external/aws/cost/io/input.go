package cost

import (
	"time"

	awsutil "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

const formatLayout = "2006-01-02"

func timePeriod(years, months, days int) *types.DateInterval {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(years, months, days).Format(formatLayout)
	end := currentTime.Format(formatLayout)

	timePeriod := &types.DateInterval{
		Start: &start,
		End:   &end,
	}
	return timePeriod
}

func CostAndUsageInput() *costexplorer.GetCostAndUsageInput {
	// Last Six Months Period.
	timePeriod := timePeriod(0, -6, 0)

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"BlendedCost"},
		TimePeriod:  timePeriod,
		GroupBy: []types.GroupDefinition{
			{
				Key:  awsutil.String("SERVICE"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
	}
	return input
}

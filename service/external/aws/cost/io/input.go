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

func groupDefinition(query string) types.GroupDefinition {
	groupDefinition := types.GroupDefinition{
		Key:  awsutil.String(query),
		Type: types.GroupDefinitionTypeDimension,
	}
	return groupDefinition
}

func CostAndUsageInput() *costexplorer.GetCostAndUsageInput {
	metrics := []string{"BlendedCost"}

	timePeriod := timePeriod(0, -6, 0)
	groupDefinition := groupDefinition("SERVICE")

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		Metrics:     metrics,
		TimePeriod:  timePeriod,
		GroupBy:     []types.GroupDefinition{groupDefinition},
	}
	return input
}

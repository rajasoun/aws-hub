package cost

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	model "github.com/rajasoun/aws-hub/service/external/aws/cost/model"
)

// Interface for Amazon CE GetCostAndUsage API
// This will enable TDD using mocking.
type GetCostAndUsageAPI interface {
	GetCostAndUsage(ctx context.Context, params *costexplorer.GetCostAndUsageInput,
		optFns ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageOutput, error)
}

// GetCost retrieves the Cost for AWS  account.
// Inputs:
//     client is iam.NewFromConfig(cfg) & cfg is the context of the method call
// Output:
//     If successful, a GetCostAndUsage object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to GetCostAndUsage.
func CurrentBill(client GetCostAndUsageAPI) (model.Bill, error) {
	emptyContext := context.TODO()

	metrics := []string{"BlendedCost"}
	// Time Period of Last Six Months
	timePeriod := timePeriod(0, -6, 0)
	groupDefinition := groupDefinition("SERVICE")

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		Metrics:     metrics,
		TimePeriod:  timePeriod,
		GroupBy:     []types.GroupDefinition{groupDefinition},
	}

	result, err := client.GetCostAndUsage(emptyContext, input)
	if err != nil {
		log.Println("Got an error retrieving account aliases")
		return model.Bill{}, err
	}

	costHistory, currentBill := GetCostHistoryAndBill(result)

	return model.Bill{
		Total:   currentBill,
		History: costHistory,
	}, nil
}

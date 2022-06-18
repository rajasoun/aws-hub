package cost

import (
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	awsModel "github.com/rajasoun/aws-hub/service/external/aws/cost/model"
)

const bitSize = 64

func groupByMetrics(res types.ResultByTime) (string, []awsModel.Group) {
	unit := "USD"
	groups := make([]awsModel.Group, 0)
	for _, group := range res.Groups {
		amount, _ := strconv.ParseFloat(*group.Metrics["BlendedCost"].Amount, bitSize)
		groups = append(groups, awsModel.Group{
			Key:    group.Keys[0],
			Amount: amount,
		})
		unit = *group.Metrics["BlendedCost"].Unit
	}
	return unit, groups
}

func appendToList(costs []awsModel.Cost, start time.Time,
	end time.Time, unit string, groups []awsModel.Group) []awsModel.Cost {
	costs = append(costs, awsModel.Cost{
		Start:  start,
		End:    end,
		Unit:   unit,
		Groups: groups,
	})
	return costs
}

func costHistory(result *costexplorer.GetCostAndUsageOutput) []awsModel.Cost {
	costList := make([]awsModel.Cost, 0)
	for _, res := range result.ResultsByTime {
		unit, groups := groupByMetrics(res)

		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Amount > groups[j].Amount
		})

		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)

		costList = appendToList(costList, start, end, unit, groups)
	}
	return costList
}

func currentBill(costHistory []awsModel.Cost) float64 {
	var currentBill float64
	for _, group := range costHistory[len(costHistory)-1].Groups {
		currentBill += group.Amount
	}
	return currentBill
}

func GetCostHistoryAndBill(result *costexplorer.GetCostAndUsageOutput) ([]awsModel.Cost, float64) {
	costHistory := costHistory(result)
	currentBill := currentBill(costHistory)
	return costHistory, currentBill
}

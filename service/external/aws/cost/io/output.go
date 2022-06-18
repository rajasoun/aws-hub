package io

import (
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	model "github.com/rajasoun/aws-hub/service/external/aws/cost/model"
)

const bitSize = 64

func groupByMetrics(res types.ResultByTime) (string, []model.Group) {
	unit := "USD"
	groups := make([]model.Group, 0)
	if len(res.Groups) == 0 {
		log.Println("Got Empty Groups")
		return unit, groups
	}
	for _, group := range res.Groups {
		amount, _ := strconv.ParseFloat(*group.Metrics["BlendedCost"].Amount, bitSize)
		groups = append(groups, model.Group{
			Key:    group.Keys[0],
			Amount: amount,
		})
		unit = *group.Metrics["BlendedCost"].Unit
	}
	return unit, groups
}

func appendToList(costs []model.Cost, start time.Time,
	end time.Time, unit string, groups []model.Group) []model.Cost {
	costs = append(costs, model.Cost{
		Start:  start,
		End:    end,
		Unit:   unit,
		Groups: groups,
	})
	return costs
}

func costHistory(result *costexplorer.GetCostAndUsageOutput) []model.Cost {
	costList := make([]model.Cost, 0)

	if len(result.ResultsByTime) == 0 {
		log.Println("Got Empty result.ResultsByTime")
		return costList
	}

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

func currentBill(costHistory []model.Cost) float64 {
	var currentBill float64
	for _, group := range costHistory[len(costHistory)-1].Groups {
		currentBill += group.Amount
	}
	return currentBill
}

func TotalBillAndHistory(result *costexplorer.GetCostAndUsageOutput) ([]model.Cost, float64) {
	costHistory := costHistory(result)
	if len(costHistory) == 0 {
		log.Println("Cost History is Empty")
		return []model.Cost{}, 0
	}
	currentBill := currentBill(costHistory)
	return costHistory, currentBill
}

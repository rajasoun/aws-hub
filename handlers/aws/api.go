package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func (handler *AWSHandler) AwsAPI(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "IAMListUsers":
		response, err = handler.aws.IAMListUsers(cfg)
	case apiName == "IAMUser":
		response, err = handler.aws.IAMUser(cfg)
		// case apiName == "DescribeCostAndUsage":
		// 	response, err = handler.aws.DescribeCostAndUsage(cfg)
		// case apiName == "DescribeCostAndUsagePerInstanceType":
		// 	response, err = handler.aws.DescribeCostAndUsagePerInstanceType(cfg)
		// case apiName == "DescribeForecastPrice":
		// 	response, err = handler.aws.DescribeForecastPrice(cfg)
	}
	return response, err
}

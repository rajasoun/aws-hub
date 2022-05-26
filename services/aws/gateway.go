package aws

import "github.com/aws/aws-sdk-go-v2/aws"

// ToDo Technical Debt - Use Interface to call the right method
// Use Dependency Injection
func (aws AWS) ExternalServiceGateway(cfg aws.Config, apiName string) (interface{}, error) {
	var response interface{}
	var err error
	switch {
	case apiName == "IAMListUsers":
		response, err = aws.IAMListUsers(cfg)
	case apiName == "IAMUser":
		response, err = aws.IAMUser(cfg)
	case apiName == "DescribeCostAndUsage":
		response, err = aws.DescribeCostAndUsage(cfg)
	case apiName == "DescribeCostAndUsagePerInstanceType":
		response, err = aws.DescribeCostAndUsagePerInstanceType(cfg)
	case apiName == "DescribeForecastPrice":
		response, err = aws.DescribeForecastPrice(cfg)
	}
	return response, err
}

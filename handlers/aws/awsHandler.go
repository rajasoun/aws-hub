package aws

import (
	"time"

	"github.com/rajasoun/aws-hub/services/cache"
)

func awsHandler() *AWSHandler {
	cache := &cache.Memory{
		Expiration: time.Duration(30),
	}
	awsHandler := NewAWSHandler(cache, true)
	return awsHandler
}

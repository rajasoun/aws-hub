package handlers

import (
	"time"

	"github.com/rajasoun/aws-hub/service/aws"
	"github.com/rajasoun/aws-hub/service/cache"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
	aws      aws.AWS
}

func NewDefaultAWSHandler(multiple bool) *AWSHandler {
	cache := &cache.Memory{Expiration: time.Duration(30)}
	cache.Connect()
	return NewAWSHandler(cache, multiple)
}

func NewAWSHandler(cache cache.Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws:      aws.AWS{},
	}
	return &awsHandler
}

func (handler *AWSHandler) GetAWSHandler() aws.AWS {
	return handler.aws
}

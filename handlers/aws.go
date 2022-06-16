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
	defaultCacheDuration := 30
	cacheHandler := &cache.Memory{Expiration: time.Duration(defaultCacheDuration)}
	cacheHandler.Connect()
	return NewAWSHandler(cacheHandler, multiple)
}

func NewAWSHandler(cacheHandler cache.Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cacheHandler,
		multiple: multiple,
		aws:      aws.AWS{},
	}
	return &awsHandler
}

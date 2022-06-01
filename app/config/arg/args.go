package arg

import (
	"time"

	"github.com/rajasoun/aws-hub/service/cache"

	"github.com/urfave/cli/v2"
)

const (
	DefaultPort      = 3000
	DefaultDuration  = 30
	DefaultCacheType = "in-memory"
)

type CliContext struct {
	ctx                  *cli.Context
	port                 int
	duration             int
	cache                cache.Cache
	isMultipleAWSProfile bool
}

func NewCliContext(appCtx *cli.Context) *CliContext {
	ctx := CliContext{ctx: appCtx}
	ctx.port = ctx.GetPort()
	ctx.duration = ctx.GetDuration()
	ctx.cache = ctx.GetCache()
	ctx.isMultipleAWSProfile = ctx.GetAwsProfileType()
	return &ctx
}

func (cli *CliContext) GetPort() int {
	port := cli.ctx.Int("port")
	if port == 0 {
		port = DefaultPort
	}
	return port
}

func (cli *CliContext) GetDuration() int {
	duration := cli.ctx.Int("duration")
	if duration == 0 {
		duration = DefaultDuration
	}
	return duration
}

func (cli *CliContext) GetCache() cache.Cache {
	var cacheHandler cache.Cache
	cache := cli.ctx.String("cache")
	if cache == "" {
		cache = DefaultCacheType
	}
	cacheHandler = cli.GetCacheHandler(cache)
	return cacheHandler
}

func (cli *CliContext) GetCacheHandler(cacheType string) cache.Cache {
	var cacheHandler cache.Cache
	duration := cli.GetDuration()
	switch {
	case cacheType == "in-memory":
		cacheHandler = cli.GetInMemoryCachehandler(duration)
	case cacheType == "redis":
		cacheHandler = cli.GetRedisCachehandler(cacheType, duration)
	}
	return cacheHandler
}

func (cli *CliContext) GetInMemoryCachehandler(duration int) cache.Cache {
	cacheHandler := &cache.Memory{
		Expiration: time.Duration(duration),
	}
	return cacheHandler
}

func (cli *CliContext) GetRedisCachehandler(redis string, duration int) cache.Cache {
	cacheHandler := &cache.Redis{
		Addr:       redis,
		Expiration: time.Duration(duration),
	}
	return cacheHandler
}

func (cli *CliContext) GetAwsProfileType() bool {
	return cli.ctx.Bool("multiple")
}

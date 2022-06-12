package arg

import (
	"time"

	"github.com/rajasoun/aws-hub/service/cache"

	cli "github.com/urfave/cli/v2"
)

const (
	DefaultPort      = 3000
	DefaultDuration  = 30
	DefaultCacheType = "in-memory"
)

type CliContext struct {
	args                 *cli.Context
	port                 int
	duration             int
	cache                cache.Cache
	isMultipleAWSProfile bool
}

func NewCliContext(appCtx *cli.Context) *CliContext {
	ctx := CliContext{args: appCtx}
	ctx.port = ctx.GetPort()
	ctx.duration = ctx.GetDuration()
	ctx.cache = ctx.GetCache()
	ctx.isMultipleAWSProfile = ctx.GetAwsProfileType()
	return &ctx
}

func (cliCtx *CliContext) GetPort() int {
	port := cliCtx.args.Int("port")
	if port == 0 {
		port = DefaultPort
	}
	return port
}

func (cliCtx *CliContext) GetDuration() int {
	duration := cliCtx.args.Int("duration")
	if duration == 0 {
		duration = DefaultDuration
	}
	return duration
}

func (cliCtx *CliContext) GetCache() cache.Cache {
	var cacheHandler cache.Cache
	cacheType := cliCtx.args.String("cache")
	if cacheType == "" {
		cacheType = DefaultCacheType
	}
	cacheHandler = cliCtx.GetCacheHandler(cacheType)
	return cacheHandler
}

func (cliCtx *CliContext) GetCacheHandler(cacheType string) cache.Cache {
	var cacheHandler cache.Cache
	duration := cliCtx.GetDuration()
	switch {
	case cacheType == "in-memory":
		cacheHandler = cliCtx.GetInMemoryCachehandler(duration)
	case cacheType == "redis":
		cacheHandler = cliCtx.GetRedisCachehandler(cacheType, duration)
	}
	return cacheHandler
}

func (cliCtx *CliContext) GetInMemoryCachehandler(duration int) cache.Cache {
	cacheHandler := &cache.Memory{
		Expiration: time.Duration(duration),
	}
	return cacheHandler
}

func (cliCtx *CliContext) GetRedisCachehandler(redis string, duration int) cache.Cache {
	cacheHandler := &cache.Redis{
		Addr:       redis,
		Expiration: time.Duration(duration),
	}
	return cacheHandler
}

func (cliCtx *CliContext) GetAwsProfileType() bool {
	return cliCtx.args.Bool("multiple")
}

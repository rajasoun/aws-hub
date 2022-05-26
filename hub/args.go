package hub

import (
	"time"

	"github.com/rajasoun/aws-hub/services/cache"

	"github.com/urfave/cli/v2"
)

const (
	DefaultPort     = 3000
	DefaultDuration = 30
)

type CliContext struct {
	ctx *cli.Context
}

func NewCliContext(appCtx *cli.Context) *CliContext {
	ctx := CliContext{
		ctx: appCtx,
	}
	return &ctx
}

func (cli *CliContext) Port() int {
	port := cli.ctx.Int("port")
	if port == 0 {
		port = DefaultPort
	}
	return port
}

func (cli *CliContext) Duration() int {
	duration := cli.ctx.Int("duration")
	if duration == 0 {
		duration = DefaultPort
	}
	return duration
}

func (cli *CliContext) Cache() cache.Cache {
	var cacheHandler cache.Cache
	redis := cli.ctx.String("redis")
	if redis == "" {
		cacheHandler = &cache.Memory{
			Expiration: time.Duration(cli.Duration()),
		}
	} else {
		cacheHandler = &cache.Redis{
			Addr:       redis,
			Expiration: time.Duration(cli.Duration()),
		}
	}
	return cacheHandler
}

func (cli *CliContext) IsMultipleAwsProfiles() bool {
	return cli.ctx.Bool("multiple")
}

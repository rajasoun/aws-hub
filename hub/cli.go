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

func parseArgs(c *cli.Context) (int, cache.Cache, bool) {
	port := c.Int("port")
	duration := c.Int("duration")
	redis := c.String("redis")
	multiple := c.Bool("multiple")

	var cacheHandler cache.Cache

	if port == 0 {
		port = DefaultPort
	}
	if duration == 0 {
		duration = DefaultDuration
	}

	if redis == "" {
		cacheHandler = &cache.Memory{
			Expiration: time.Duration(duration),
		}
	} else {
		cacheHandler = &cache.Redis{
			Addr:       redis,
			Expiration: time.Duration(duration),
		}
	}
	return port, cacheHandler, multiple
}

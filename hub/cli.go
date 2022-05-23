package hub

import (
	"time"

	"github.com/rajasoun/aws-hub/services/cache"
	structs "github.com/rajasoun/go-ds"
	"github.com/urfave/cli/v2"
)

const (
	DefaultPort     = 3000
	DefaultDuration = 30
)

func getAuthors() []*cli.Author {
	authors := []*cli.Author{
		{
			Name:  "Raja Soundaramourty",
			Email: "rajasoun@cisco.com",
		},
	}
	return authors
}

func setUpApp(app *cli.App) {
	app.Name = "AWS Hub"
	app.Usage = "AWS Cost Explorer"
	app.Version = "0.0.1"
	app.Authors = getAuthors()
	app.Compiled = time.Now()
}

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

func StructToMap(app *cli.App) map[string]interface{} {
	s := structs.New(app)
	m := s.Map()
	return m
}

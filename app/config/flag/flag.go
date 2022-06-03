package flag

import (
	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/urfave/cli/v2"
)

func GetFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Usage: "Server port",
			Value: arg.DefaultPort,
		},
		&cli.IntFlag{
			Name:  "duration, d",
			Usage: "Cache expiration time",
			Value: arg.DefaultDuration,
		},
		&cli.StringFlag{
			Name:  "cache, c",
			Usage: "Cache Type",
			Value: arg.DefaultCacheType,
		},
		&cli.BoolFlag{
			Name:  "multiple, m",
			Usage: "Enable multiple AWS accounts",
		},
	}
	return flags
}

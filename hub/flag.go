package hub

import "github.com/urfave/cli/v2"

var flags []cli.Flag

func setFlags(app *cli.App) {
	flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Usage: "Server port",
			Value: DefaultPort,
		},
		&cli.IntFlag{
			Name:  "duration, d",
			Usage: "Cache expiration time",
			Value: DefaultDuration,
		},
		&cli.StringFlag{
			Name:  "redis, r",
			Usage: "Redis server",
		},
		&cli.StringFlag{
			Name:  "dataset, ds",
			Usage: "BigQuery Bill dataset",
		},
		&cli.BoolFlag{
			Name:  "multiple, m",
			Usage: "Enable multiple AWS accounts",
		},
	}
}

func getFlags() []cli.Flag {
	return flags
}

func sliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}

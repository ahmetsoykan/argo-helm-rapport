package watch

import (
	"github.com/urfave/cli/v2"
)

func Commands() (cmd []*cli.Command) {

	cmd = []*cli.Command{
		&cli.Command{
			Name:  "watch",
			Usage: "set charts to be detected and rendered afterwards",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:   "chart",
					Usage:  "add new chart repositories",
					Action: WatchChartAction,
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "name", Required: true, Aliases: []string{"n"}},
					},
				},
			},
		},
	}
	return
}

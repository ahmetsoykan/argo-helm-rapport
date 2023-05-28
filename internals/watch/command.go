package watch

import (
	"github.com/urfave/cli/v2"
)

func Commands() (cmd []*cli.Command) {

	cmd = []*cli.Command{
		&cli.Command{
			Name:  "watch",
			Usage: "list the argo applications under the given path",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:   "chart",
					Usage:  "add new chart repositories",
					Action: WatchChartAction,
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "name", Required: true, Aliases: []string{"n"}},
						&cli.StringFlag{Name: "prev", Required: true, Aliases: []string{"p"}},
						&cli.StringFlag{Name: "curr", Required: true, Aliases: []string{"c"}},
					},
				},
			},
		},
	}
	return
}

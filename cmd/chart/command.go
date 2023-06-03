package chart

import (
	"github.com/urfave/cli/v2"
)

func Commands() (cmd []*cli.Command) {

	cmd = []*cli.Command{
		&cli.Command{
			Name:  "chart",
			Usage: "chart actions",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:   "render",
					Usage:  "templates the charts",
					Action: ChartRenderAction,
				},
				&cli.Command{
					Name:   "compare",
					Usage:  "compare the templates",
					Action: ChartCompareAction,
				},
			},
		},
	}
	return
}

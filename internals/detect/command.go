package detect

import (
	"github.com/urfave/cli/v2"
)

func Commands() (cmd []*cli.Command) {

	cmd = []*cli.Command{
		&cli.Command{
			Name:   "detect",
			Usage:  "list the argo applications under the given path",
			Action: DetectAction,
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "path", Required: true, Aliases: []string{"p"}},
			},
		},
	}
	return
}

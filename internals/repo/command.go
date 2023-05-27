package repo

import (
	"github.com/urfave/cli/v2"
)

func Commands() (cmd []*cli.Command) {

	cmd = []*cli.Command{
		&cli.Command{
			Name:  "repo",
			Usage: "chart repository actions",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:   "list",
					Usage:  "list added chart repositories",
					Action: RepoListAction,
				},
				&cli.Command{
					Name:   "add",
					Usage:  "add new chart repositories",
					Action: RepoAddAction,
					Flags: []cli.Flag{
						&cli.BoolFlag{Value: false, Name: "private"},
						&cli.StringFlag{Name: "name", Required: true, Aliases: []string{"n"}},
						&cli.StringFlag{Name: "host", Required: true},
						&cli.StringFlag{Name: "username", DefaultText: "", Aliases: []string{"u"}},
						&cli.StringFlag{Name: "password", DefaultText: "", Aliases: []string{"p"}},
					},
				},
			},
		},
	}
	return
}

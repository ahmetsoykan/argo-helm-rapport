package repo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestRepo(t *testing.T) {
	cases := []struct {
		name string
		got  []string
		want error
	}{
		{name: "RepoAddSucess", got: []string{"test", "repo", "add", "--name", "chartname", "--host", "chartname.io"}, want: nil},
		{name: "RepoAddFailedWithMissingHost", got: []string{"test", "repo", "add", "--name", "chartname"}, want: errors.New("Required flag \"host\" not set")},
		{name: "RepoAddFailedWithMissingName", got: []string{"test", "repo", "add", "--host", "cahrtname.io"}, want: errors.New("Required flag \"name\" not set")},
		{name: "RepoAddFailedWithMissingNameAndHost", got: []string{"test", "repo", "add"}, want: errors.New("Required flags \"name, host\" not set")},
		{name: "RepoListSucess", got: []string{"test", "repo", "list"}, want: nil},
	}

	for _, c := range cases {

		app := &cli.App{
			Name:     "test",
			Usage:    "Shows you the chart differences",
			Commands: Commands(),
		}

		if err := app.Run(c.got); err != nil {
			if fmt.Sprintf(c.want.Error()) != fmt.Sprintf(err.Error()) {
				t.Errorf("%s, test failed with error %s", c.name, err)
			}
		}
	}
}

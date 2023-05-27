package repo

import (
	"fmt"

	"github.com/ahmetsoykan/helm-rapport/internals/data"
	"github.com/urfave/cli/v2"
)

func RepoListAction(ctx *cli.Context) error {
	repositories, err := data.GetRepositories()
	if err != nil {
		return err
	}
	for i, repository := range repositories {
		fmt.Printf("%d| %s| %s| %v|\n", i, repository.Name, repository.Host, repository.Private)
	}

	return nil
}

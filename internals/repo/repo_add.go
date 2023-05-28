package repo

import (
	"errors"
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/urfave/cli/v2"
)

func RepoAddAction(ctx *cli.Context) error {

	var _repository data.Repository
	name := ctx.String("name")
	host := ctx.String("host")
	private := ctx.String("private")

	if private == "true" {
		username := ctx.String("username")
		password := ctx.String("password")

		if (username == "") || (password == "") {
			return errors.New("credentials need to be passed if the chart repository is private")
		}

		_repository = data.Repository{
			Name:    name,
			Host:    host,
			Private: true,
			Credentials: data.Credentials{
				Username: username,
				Password: password,
			},
		}
	} else {

		_repository = data.Repository{
			Name:    name,
			Host:    host,
			Private: false,
		}

	}

	err := data.WriteRepositoryToFile(_repository)
	if err != nil {
		return err
	} else {
		fmt.Println("Repository added.")
	}
	return nil
}

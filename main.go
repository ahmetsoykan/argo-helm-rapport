package main

import (
	"log"
	"os"

	"github.com/ahmetsoykan/helm-rapport/internals/repo"
	"github.com/urfave/cli/v2"
)

var (
	appName string = "argo-helm-rapport"
)

func main() {

	//* CLI Initialization
	app := &cli.App{
		Name:  appName,
		Usage: "Shows you the helm chart differences between versions",
	}
	app.Commands = append(app.Commands, repo.Commands()...)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

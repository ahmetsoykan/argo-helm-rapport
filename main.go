package main

import (
	"log"
	"os"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/detect"
	"github.com/ahmetsoykan/argo-helm-rapport/internals/repo"
	"github.com/ahmetsoykan/argo-helm-rapport/internals/watch"

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
	app.Commands = append(app.Commands, detect.Commands()...)
	app.Commands = append(app.Commands, watch.Commands()...)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package detect

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/ahmetsoykan/argo-helm-rapport/internals/watch"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var (
	apps []data.App
)

func DetectAction(ctx *cli.Context) error {

	path := ctx.String("path")

	// walk through all files under the given path
	var filePaths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// find apps folders and its application values
		if strings.Contains(path, "apps") && strings.Contains(path, "values.yaml") {
			filePaths = append(filePaths, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	// collect information about applications
	for _, file := range filePaths {
		f, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		var app data.App
		if err := yaml.Unmarshal(f, &app); err != nil {
			log.Fatal(err)
		}

		// filter charts detected
		charts, err := watch.GetWatchedCharts()
		if err != nil {
			return err
		}

		for k, _ := range app {
			for _, chart := range charts {
				if chart.Name == k {
					apps = append(apps, app)
				}
			}
		}
	}

	// Write found charts to as a file
	err = data.WriteChartsToFile(apps)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

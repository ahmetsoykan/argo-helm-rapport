package chart

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	helmclient "github.com/mittwald/go-helm-client"
	"github.com/urfave/cli/v2"
)

func ChartRenderAction(ctx *cli.Context) error {

	apps, err := data.GetApps()
	if err != nil {
		return err
	}

	repos, err := data.GetRepositories()
	if err != nil {
		return err
	}

	// checks
	for k, v := range apps {
		for i, app := range v {

			if len(app.Versions) == 1 {
				return errors.New("at least two version required to compare")
			}
			if app.Versions[0] == app.Versions[1] {
				apps[k][i].DiffVersions = false
			} else {
				apps[k][i].DiffVersions = true
			}
			if !bytes.Equal(app.MergedValueFiles[0], app.MergedValueFiles[1]) {
				apps[k][i].DiffValues = true
			} else {
				apps[k][i].DiffValues = false
			}

			var found bool
			for _, repo := range repos {
				if repo.Name == app.ChartRepository {
					found = true
				}
			}

			if !found {
				return errors.New("this chart repository has to be recognised:" + app.ChartRepository)
			}

		}
	}

	// helm client
	helmClient, err := NewHelmClient()
	if err != nil {
		return nil
	}
	// run `helm repo add`
	helmClient.AddChartRepo()

	for k, v := range apps {
		for i, app := range v {

			if app.DiffVersions || app.DiffValues {

				var renderedFiles []string
				for x, ver := range app.Versions {

					chartSpec := helmclient.ChartSpec{
						ReleaseName:      app.Name,
						ChartName:        app.ChartRepository + "/" + app.DependencyName,
						Namespace:        app.Namespace,
						DependencyUpdate: true,
						UpgradeCRDs:      true,
						Wait:             true,
						Version:          ver,
						ValuesYaml:       string(app.MergedValueFiles[x]),
					}

					yamlData, err := helmClient.TemplateChart(&chartSpec, &helmclient.HelmTemplateOptions{})
					if err != nil {
						return err
					}

					tempStr := data.RandStringBytes(5)
					err = data.WriteYamlToFile(k+"_"+ver+tempStr, yamlData)
					if err != nil {
						return err
					}
					renderedFiles = append(renderedFiles, data.UserDir+k+"_"+ver+tempStr)

				}
				apps[k][i].RenderedFiles = renderedFiles

			}

		}
	}

	// fmt.Printf("%+v\n", apps) // Uncomment this to debug locally
	// Write found charts to as a file
	err = data.WriteAppsToFile(apps)
	if err != nil {
		return err
	} else {
		fmt.Println("Charts rendered successfully.")
	}

	return nil
}

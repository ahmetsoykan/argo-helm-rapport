package chart

import (
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	helmclient "github.com/mittwald/go-helm-client"
	"github.com/urfave/cli/v2"
)

func ChartRenderAction(ctx *cli.Context) error {

	apps, err := data.GetCharts()
	if err != nil {
		return err
	}
	fmt.Println(apps)

	for _, app := range apps {
		for k, v := range app {
			for _, ver := range v.Versions {
				helmClient, _ := NewHelmClient()
				helmClient.AddChartRepo()

				chartSpec := helmclient.ChartSpec{
					ReleaseName: k,
					ChartName:   v.ChartRepo + "/" + k,
					Namespace:   v.Namespace,
					UpgradeCRDs: true,
					Wait:        true,
					Version:     ver,
					ValuesYaml:  ``,
				}

				yamlData, err := helmClient.TemplateChart(&chartSpec, &helmclient.HelmTemplateOptions{})
				if err != nil {
					return err
				}

				err = data.WriteYamlToFile(k+"_"+ver, yamlData)
				if err != nil {
					return nil
				}

			}
		}
	}

	return nil
}

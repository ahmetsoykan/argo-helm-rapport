package chart

import (
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli/v2"
)

func ChartCompareAction(ctx *cli.Context) error {

	apps, err := data.GetCharts()
	if err != nil {
		return err
	}
	fmt.Println(apps)

	dmp := diffmatchpatch.New()

	for _, app := range apps {
		for k, v := range app {

			text1, _ := data.ReadYamlFromFile(k + "_" + v.Versions[0])
			text2, _ := data.ReadYamlFromFile(k + "_" + v.Versions[1])
			diffs := dmp.DiffMain(string(text1), string(text2), false)

			fmt.Println(k)
			fmt.Println(dmp.DiffPrettyText(diffs))
			fmt.Println("---")

		}
	}

	return nil
}

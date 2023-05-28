package watch

import (
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/urfave/cli/v2"
)

func WatchChartAction(ctx *cli.Context) error {

	var _chart data.Chart
	name := ctx.String("name")
	prev := ctx.String("prev")
	curr := ctx.String("curr")

	_chart = data.Chart{
		Name: name,
		Prev: prev,
		Curr: curr,
	}

	err := data.WriteWatchToFile(_chart)
	if err != nil {
		return err
	} else {
		fmt.Println("Chart watched.")
	}
	return nil
}

func GetWatchedCharts() ([]data.Chart, error) {
	charts, err := data.GetWatchs()
	if err != nil {
		return nil, err
	}

	return charts, nil
}

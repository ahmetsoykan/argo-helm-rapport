package chart

import (
	"errors"
	"log"
	"testing"

	"github.com/ahmetsoykan/argo-helm-rapport/cmd/detect"
	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"

	"github.com/urfave/cli/v2"
)

func TestChart(t *testing.T) {
	setup()

	cases := []struct {
		name string
		got  []string
		want error
	}{
		{name: "DetectSucess", got: []string{"test", "detect", "-p", "../detect/example-folder"}, want: nil},
		{name: "RenderFailWithoutTwoTimesRunOfDetectAction", got: []string{"test", "chart", "render"}, want: errors.New("at least two version required to compare")},
		{name: "DetectSucess", got: []string{"test", "detect", "-p", "../detect/example-folder"}, want: nil},
		{name: "RenderSucess", got: []string{"test", "chart", "render"}, want: nil},
		{name: "RenderSucessWithTheSameVersion", got: []string{"test", "chart", "render"}, want: nil},
		{name: "CompareSuccess", got: []string{"test", "chart", "compare"}, want: nil},
	}

	for _, c := range cases {

		app := &cli.App{
			Name:     "test",
			Usage:    "Shows you the helm chart differences between versions",
			Commands: Commands(),
		}
		app.Commands = append(app.Commands, detect.Commands()...) // required for the second failing test

		if err := app.Run(c.got); err != nil {
			if c.want.Error() != err.Error() {
				t.Errorf("%s, test failed with error %s", c.name, err)
			}
		} else {
			if nil != c.want {
				t.Errorf("%s, test failed with error %s", c.name, err)
			}
		}
	}
}

func setup() {
	err := data.DeleteApps()
	if err != nil {
		log.Fatal(err)
	}
}

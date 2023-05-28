package watch

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestWatch(t *testing.T) {
	cases := []struct {
		name string
		got  []string
		want error
	}{
		{name: "ChartWatchSucess", got: []string{"test", "watch", "chart", "--name", "myapp", "--prev", "0.15.2", "--curr", "0.15.3"}, want: nil},
	}

	for _, c := range cases {

		app := &cli.App{
			Name:     "test",
			Usage:    "Shows you the helm chart differences between versions",
			Commands: Commands(),
		}

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

func TestGetWatchedCharts(t *testing.T) {
	_, err := GetWatchedCharts()
	if err != nil {
		t.Errorf("%s, test failed with error %s", "TestGetWatchedCharts", err)
	}
}

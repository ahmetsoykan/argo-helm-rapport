package chart

import (
	"testing"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
)

func TestAddChartRepo(t *testing.T) {

	_repository := data.Repository{
		Name:    "stable",
		Host:    "https://charts.helm.sh/stable",
		Private: false,
	}
	err := data.WriteRepositoryToFile(_repository)
	if err != nil {
		t.Errorf("%s, test failed with error %s", "TestAddChartRepo", err)
	}

	client, err := NewHelmClient()
	if err != nil {
		t.Errorf("%s, test failed with error %s", "TestAddChartRepo", err)
	}
	err = client.AddChartRepo()
	if err != nil {
		t.Errorf("%s, test failed with error %s", "TestAddChartRepo", err)
	}

}

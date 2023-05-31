package chart

import (
	"fmt"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"helm.sh/helm/v3/pkg/repo"
)

func (c HelmClient) AddChartRepo() error {

	repositories, err := data.GetRepositories()
	if err != nil {
		return err
	}

	for _, repository := range repositories {

		var chartRepo repo.Entry
		if repository.Private {

			chartRepo = repo.Entry{
				Name:               repository.Name,
				URL:                repository.Host,
				Username:           repository.Credentials.Username,
				Password:           repository.Credentials.Password,
				PassCredentialsAll: true,
			}
		} else {

			chartRepo = repo.Entry{
				Name:               repository.Name,
				URL:                repository.Host,
				PassCredentialsAll: true,
			}
		}

		// Add a chart-repository to the client.
		if err := c.AddOrUpdateChartRepo(chartRepo); err != nil {
			return err
		} else {
			fmt.Printf("%s chart repository added.\n", chartRepo.Name)
		}

	}

	return nil
}

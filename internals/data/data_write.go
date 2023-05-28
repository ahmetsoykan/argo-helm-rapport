package data

import (
	"encoding/json"
	"os"
)

func WriteRepositoryToFile(r Repository) error {
	_repositories, err := GetRepositories()
	if err != nil {
		return err
	}

	_repositories = append(_repositories, r)

	file, err := json.MarshalIndent(_repositories, "", " ")
	if err != nil {
		return err
	}

	_ = os.WriteFile(userDir+"repositories.json", file, 0644)

	return nil
}

func WriteWatchToFile(c Chart) error {
	_charts, err := GetWatchs()
	if err != nil {
		return err
	}

	_charts = append(_charts, c)

	file, err := json.MarshalIndent(_charts, "", " ")
	if err != nil {
		return err
	}

	_ = os.WriteFile(userDir+"watchs.json", file, 0644)

	return nil
}

func WriteChartsToFile(a []App) error {
	file, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return err
	}

	_ = os.WriteFile(userDir+"charts.json", file, 0644)

	return nil
}

package data

import (
	"encoding/json"
	"os"
)

var (
	userDir string = "/tmp/"
)

func GetRepositories() ([]Repository, error) {

	if _, err := os.Stat(userDir + "repositories.json"); err != nil {
		return []Repository{}, nil
	}

	file, err := os.ReadFile(userDir + "repositories.json")
	if err != nil {
		return nil, err
	}

	_repositories := []Repository{}
	err = json.Unmarshal(file, &_repositories)
	if err != nil {
		return nil, err
	}

	return _repositories, nil
}

func GetWatchs() ([]Chart, error) {

	if _, err := os.Stat(userDir + "watchs.json"); err != nil {
		return []Chart{}, nil
	}

	file, err := os.ReadFile(userDir + "watchs.json")
	if err != nil {
		return nil, err
	}

	_charts := []Chart{}
	err = json.Unmarshal(file, &_charts)
	if err != nil {
		return nil, err
	}

	return _charts, nil
}

func GetCharts() ([]AppMeta, error) {
	if _, err := os.Stat(userDir + "charts.json"); err != nil {
		return []AppMeta{}, nil
	}

	file, err := os.ReadFile(userDir + "charts.json")
	if err != nil {
		return nil, err
	}

	_charts := []AppMeta{}
	err = json.Unmarshal(file, &_charts)
	if err != nil {
		return nil, err
	}

	return _charts, nil
}

func ReadYamlFromFile(filename string) (string, error) {
	data, err := os.ReadFile(userDir + filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetApps() (map[string][]App, error) {
	if _, err := os.Stat(userDir + "apps.json"); err != nil {
		return map[string][]App{}, nil
	}

	file, err := os.ReadFile(userDir + "apps.json")
	if err != nil {
		return nil, err
	}

	_apps := make(map[string][]App)
	err = json.Unmarshal(file, &_apps)
	if err != nil {
		return nil, err
	}

	return _apps, nil
}

package data

import (
	"encoding/json"
	"os"
)

var (
	UserDir string = "/tmp/"
)

// repo pkg datas
func GetRepositories() ([]Repository, error) {

	if _, err := os.Stat(UserDir + "repositories.json"); err != nil {
		return []Repository{}, nil
	}

	file, err := os.ReadFile(UserDir + "repositories.json")
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

// watch pkg datas
func GetWatchs() ([]Chart, error) {

	if _, err := os.Stat(UserDir + "watchs.json"); err != nil {
		return []Chart{}, nil
	}

	file, err := os.ReadFile(UserDir + "watchs.json")
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

// detects pkg - logs applications that detected, which render and compare works based on these findings
func GetApps() (map[string][]App, error) {
	if _, err := os.Stat(UserDir + "apps.json"); err != nil {
		return map[string][]App{}, nil
	}

	file, err := os.ReadFile(UserDir + "apps.json")
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

// chart pkg - before compare to read rendered yaml files
func ReadYamlFromFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

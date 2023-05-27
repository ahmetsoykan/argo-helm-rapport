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

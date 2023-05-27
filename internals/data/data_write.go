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

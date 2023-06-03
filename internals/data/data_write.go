package data

import (
	"encoding/json"
	"math/rand"
	"os"
)

// repo pkg datas
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

	_ = os.WriteFile(UserDir+"repositories.json", file, 0644)

	return nil
}

// watch pkg datas
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

	_ = os.WriteFile(UserDir+"watchs.json", file, 0644)

	return nil
}

// chart pkg - save rendered yaml
func WriteYamlToFile(filename string, y []byte) error {
	_ = os.WriteFile(UserDir+filename, y, 0644)

	return nil
}

// detect and chart pkg
func WriteAppsToFile(a map[string][]App) error {
	file, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return err
	}

	_ = os.WriteFile(UserDir+"apps.json", file, 0644)

	return nil
}

// chart pkg - test setup
func DeleteApps() error {

	err := os.Remove(UserDir + "apps.json")
	if err != nil {
		return err
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

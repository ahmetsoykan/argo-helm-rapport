package detect

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ahmetsoykan/argo-helm-rapport/cmd/watch"
	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var (
	apps = make(map[string][]data.App)
)

func DetectAction(ctx *cli.Context) error {

	path := ctx.String("path")

	watchedCharts, err := watch.GetWatchedCharts()
	if err != nil {
		return err
	}

	// we can find directories that match with watched chart names under the given folder
	// assumption: chart names and folder names the same
	chartPaths := make(map[string][]string)
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, chart := range watchedCharts {
			if strings.Contains(path, chart.Name) {
				if strings.Contains(path, chart.Name+"/Chart.yaml") {
					pathSplit := strings.Split(path, "/")
					chartPaths[chart.Name] = append(chartPaths[chart.Name], strings.Join(pathSplit[:len(pathSplit)-1], "/"))
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// when the charts found, we can go one upper level directory to find "apps" folder
	for k, v := range chartPaths {
		for _, chartPath := range v {
			splittedPath := strings.Split(chartPath, "/")
			newPath := splittedPath[:len(splittedPath)-1]
			apps[k] = append(apps[k], data.App{DirectoryPath: strings.Join(newPath, "/")})
		}
	}

	// when we know the directory that has "apps" and "actual chart" directory, we can find the app setting file
	// assumption: every chart directory has a apps dicrectory at the same level.
	// assumption: every apps/values.yaml has a key with name of the chart
	for k, v := range apps {
		for i, app := range v {
			err = filepath.Walk(app.DirectoryPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if strings.Contains(path, "apps") && strings.Contains(path, "values.yaml") {

					f, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					var metaData data.AppMeta
					if err := yaml.Unmarshal(f, &metaData); err != nil {
						return err
					}
					v[i].Name = metaData[k].Name
					v[i].Namespace = metaData[k].Namespace

					// read values files from apps's values.yaml
					var valuesEnv []string
					if metaData[k].Plugin.Name == "kustomized-helm" {
						for _, val := range metaData[k].Plugin.Env {
							if val.Name == "HELM_ARGS" {
								tempValues := strings.Split(val.Value, "-f")
								if len(tempValues) < 1 && len(tempValues) > 2 {
									valuesEnv = []string{"./values.yaml"}
								} else {
									valuesEnv = []string{tempValues[1], tempValues[2]}
								}
							}
						}
					} else {

						if metaData[k].ValueFiles != nil {
							valuesEnv = append(valuesEnv, metaData[k].ValueFiles...)
						} else {
							valuesEnv = []string{"./values.yaml"}
						}
					}

					v[i].ValueFiles = valuesEnv

					// reassign it
					apps[k] = v
				}

				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	// we have: [name of the directory that has our chart, chart name, namespace to be deployed, values.yaml files' relative paths]
	// we can find the current version and the its chart repository with these information
	// assumption: every chart has its own folder with the exact name match
	for k, v := range apps {
		for i, app := range v {
			err = filepath.Walk(app.DirectoryPath+"/"+k, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if strings.Contains(path, "Chart.yaml") {

					f, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					var req data.Requirement
					if err := yaml.Unmarshal(f, &req); err != nil {
						return err
					}

					// for _, subcharts := range req.Dependencies {
					// 	if subcharts.Name == k {
					// 		v[i].ChartRepository = subcharts.Repository
					// 		v[i].Versions = append(v[i].Versions, subcharts.Version)
					// 	}
					// }

					// assumption it has only one subchart
					v[i].DependencyName = req.Dependencies[0].Name
					v[i].ChartRepository = req.Dependencies[0].Repository
					v[i].Versions = append(v[i].Versions, req.Dependencies[0].Version)

					apps[k] = v
				}

				if strings.Contains(path, "kustomization.yaml") {

					f, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					var d data.KustomizePacth
					if err := yaml.Unmarshal(f, &d); err != nil {
						return err
					}
					if d.Components != nil {
						v[i].KustomizeComponentRepoURL = d.Components[0]
					}
				}

				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	// minor data manipulation
	for k, v := range apps {
		for i, app := range v {
			// remove @ prefix from chart repository
			apps[k][i].ChartRepository = strings.ReplaceAll(app.ChartRepository, "@", "")
			// values file change to absolute from relative path
			updatedValuesFiles := make([]string, 0)
			for _, val := range app.ValueFiles {
				val = strings.TrimSpace(val)
				updatedValuesFiles = append(updatedValuesFiles, filepath.Join(app.DirectoryPath, fmt.Sprintf("./%s/", app.Name), val))
			}
			apps[k][i].ValueFiles = updatedValuesFiles
		}
	}

	prevApp, err := data.GetApps()
	if err != nil {
		return err
	}
	for k, _ := range prevApp {
		if _, ok := apps[k]; ok {
			//checking the previously obtained versions
			for x, y := range apps[k] {
				for a, b := range prevApp[k] {
					if y.DirectoryPath == b.DirectoryPath {
						if y.Name == b.Name {
							if len(prevApp[k][a].Versions) < 2 {
								apps[k][x].Versions = append(prevApp[k][a].Versions, apps[k][x].Versions...)
							} else {
								apps[k][x].Versions = append(prevApp[k][a].Versions[1:], apps[k][x].Versions...)
							}

						}
					}
				}
			}
		} else {
			apps[k] = prevApp[k]
		}
	}

	for k, v := range apps {
		for i, app := range v {

			opts := &data.Options{
				ValueFiles: app.ValueFiles,
			}
			typedCombinedValues, err := opts.MergeValues()
			if err != nil {
				return err
			}
			byteCombinedValues, err := yaml.Marshal(&typedCombinedValues)
			if err != nil {
				return err
			}

			// apps[k][i].MergedValueFiles = append(prevApp[k][i].MergedValueFiles, byteCombinedValues)
			if _, ok := prevApp[k]; ok {
				for _, b := range prevApp[k] {
					if b.DirectoryPath == app.DirectoryPath {
						if b.Name == app.Name {
							if len(b.MergedValueFiles) >= 2 {
								apps[k][i].MergedValueFiles = append(b.MergedValueFiles[len(b.MergedValueFiles)-1:], byteCombinedValues)
							} else if len(b.MergedValueFiles) != 0 {
								apps[k][i].MergedValueFiles = append(b.MergedValueFiles, byteCombinedValues)
							}
						}
					}
				}
			} else {
				apps[k][i].MergedValueFiles = append(apps[k][i].MergedValueFiles, byteCombinedValues)
			}
		}
	}

	// fmt.Printf("%+v", apps)
	// write found charts to as a file
	err = data.WriteAppsToFile(apps)
	if err != nil {
		return err
	}

	return nil
}

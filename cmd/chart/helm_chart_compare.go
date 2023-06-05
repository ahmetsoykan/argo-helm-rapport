package chart

import (
	"bytes"
	"os"

	"github.com/ahmetsoykan/argo-helm-rapport/internals/data"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kustomize/kyaml/filesys"
	. "sigs.k8s.io/kustomize/kustomize/v5/commands/build"
)

func ChartCompareAction(ctx *cli.Context) error {

	apps, err := data.GetApps()
	if err != nil {
		return err
	}

	for _, v := range apps {
		for _, app := range v {

			if app.DiffVersions || app.DiffValues {
				// if kustomize patch required, apply it one by one

				if app.KustomizeComponentRepoURL != "" {
					kustomData1 := data.KustomizePacth{
						APIVersion: "kustomize.config.k8s.io/v1beta1",
						Kind:       "Kustomization",
						Resources:  []string{"all.yaml"},
						Components: []string{app.KustomizeComponentRepoURL},
					}
					patchedData1, err := kustomize(app.RenderedFiles[0], kustomData1)
					if err != nil {
						return err
					}
					_ = os.WriteFile(app.RenderedFiles[0], patchedData1, 0644)
					kustomData2 := data.KustomizePacth{
						APIVersion: "kustomize.config.k8s.io/v1beta1",
						Kind:       "Kustomization",
						Resources:  []string{"all.yaml"},
						Components: []string{app.KustomizeComponentRepoURL},
					}
					patchedData2, err := kustomize(app.RenderedFiles[1], kustomData2)
					if err != nil {
						return err
					}
					_ = os.WriteFile(app.RenderedFiles[1], patchedData2, 0644)
				}

				Exec(app.RenderedFiles[0], app.RenderedFiles[1], Options{})
			}

		}
	}

	return nil
}

func kustomize(helmManifest string, d data.KustomizePacth) ([]byte, error) {

	// create tmp directory
	fSys := filesys.MakeFsOnDisk()
	fSys.Mkdir("k")

	defer fSys.RemoveAll("k")

	// all.yaml
	all, _ := data.ReadYamlFromFile(helmManifest)
	fSys.WriteFile("k/all.yaml", all)

	// kustomization.yaml
	kustom, _ := yaml.Marshal(d)
	fSys.WriteFile("k/kustomization.yaml", kustom)

	buffy := new(bytes.Buffer)
	cmd := NewCmdBuild(fSys, MakeHelp("foo", "bar"), buffy)
	if err := cmd.RunE(cmd, []string{"./k/"}); err != nil {
		return []byte{}, err
	}

	return buffy.Bytes(), nil
}

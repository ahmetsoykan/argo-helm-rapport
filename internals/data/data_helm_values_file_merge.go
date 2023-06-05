package data

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// Options captures the different ways to specify values
type Options struct {
	ValueFiles []string // -f/--values
}

func (opts *Options) MergeValues() (map[string]interface{}, error) {
	base := map[string]interface{}{}

	// User specified a values files via -f/--values
	for _, filePath := range opts.ValueFiles {
		currentMap := map[string]interface{}{}

		bytes, err := ReadYamlFromFile(filePath)
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
			return nil, errors.New("failed to parse" + filePath + err.Error())
		}
		// Merge with the previous map
		base = mergeMaps(base, currentMap)
	}

	return base, nil
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

//go:build generator

package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

// readPermissions reads PermissionsFile
// Returns map[OperationID][]Permission
func readPermissions(cfg *envConfig) (map[string][]string, error) {
	var config map[string][]string
	err := readYamlFile(cfg.ConfigFile, &config)
	if err != nil {
		return nil, err
	}

	var permissions map[string][]string
	err = readYamlFile(cfg.PermissionsFile, &permissions)
	if err != nil {
		return nil, err
	}

	operationIDs := make(map[string]bool)
	for _, list := range config {
		for _, k := range list {
			operationIDs[k] = true
		}
	}

	for k, v := range permissions {
		if !operationIDs[k] {
			delete(permissions, k)
			continue
		}

		slices.Sort(v)
		permissions[k] = v
		if len(v) == 0 {
			delete(permissions, k)
		}
	}

	// Write permissions back to file
	// Removes all unknown permissions not listed in the config file
	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(yamlTabSize)
	err = encoder.Encode(&permissions)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(cfg.PermissionsFile, buffer.Bytes(), writeMode)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func readYamlFile(path string, out any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, out)
	if err != nil {
		return fmt.Errorf("error parsing yaml file %q: %v", path, err)
	}
	return nil
}

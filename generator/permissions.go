//go:build generator

package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

// readPermissions reads PermissionsFile
// Returns map[OperationID][]Permission
func readPermissions(cfg *envConfig) (map[string][]string, error) {
	b, err := os.ReadFile(cfg.PermissionsFile)
	if err != nil {
		return nil, err
	}

	var permissions map[string][]string
	err = yaml.Unmarshal(b, &permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

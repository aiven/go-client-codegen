package aiven

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed permissions.yaml
var permissionBytes []byte

// Permissions returns the map of operation IDs to permission strings.
func Permissions() (map[string][]string, error) {
	var m map[string][]string
	err := yaml.Unmarshal(permissionBytes, &m)
	return m, err
}

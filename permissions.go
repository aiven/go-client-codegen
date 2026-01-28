package aiven

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed permissions.yaml
var permissionsBytes []byte
var permissionsMap map[string][]string

func init() {
	err := yaml.Unmarshal(permissionsBytes, &permissionsMap)
	if err != nil {
		panic(fmt.Sprintf("Error parsing permissions: %v", err))
	}
}

// Permissions returns the map of operation IDs to permission strings.
func Permissions() map[string][]string {
	return permissionsMap
}

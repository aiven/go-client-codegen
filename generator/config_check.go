//go:build generator

package main

import (
	"fmt"
	"strings"
)

func checkDuplicateEndpoints(config map[string][]string) error {
	endpoints := make(map[string]struct{})
	duplicates := make(map[string]struct{})

	for _, methods := range config {
		for _, method := range methods {
			if _, exists := endpoints[method]; exists {
				duplicates[method] = struct{}{}
			} else {
				endpoints[method] = struct{}{}
			}
		}
	}

	if len(duplicates) > 0 {
		return fmt.Errorf("duplicate endpoints found in config: %v", strings.Join(sortedKeys(duplicates), ", "))
	}
	return nil
}

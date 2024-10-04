package main

import (
	"fmt"
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
		return fmt.Errorf("Duplicate endpoints found in config: %v", keys(duplicates))
	}
	return nil
}

func keys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

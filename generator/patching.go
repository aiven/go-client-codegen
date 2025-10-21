//go:build generator

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

func patchDict[T map[string]any](dest, patch T) error {
	for k, v := range patch {
		result, ok := dest[k]
		switch {
		case !ok:
			result = v
		case reflect.TypeOf(result) != reflect.TypeOf(v):
			return fmt.Errorf("type missmatch for key %s", k)
		case reflect.TypeOf(result).Kind() == reflect.Map:
			err := patchDict(result.(T), v.(T)) //nolint:forcetypeassert
			if err != nil {
				return err
			}
		default:
			result = v
		}

		dest[k] = result
	}
	return nil
}

// readOpenAPIPatched Reads OpenAPI file (JSON) and applies the patch (YAML)
func readOpenAPIPatched(oaFile, patchFile string) ([]byte, error) {
	dest, err := os.ReadFile(filepath.Clean(oaFile))
	if err != nil {
		return nil, err
	}

	patch, err := os.ReadFile(filepath.Clean(patchFile))
	if errors.Is(err, os.ErrExist) {
		// No patch found, exits
		return dest, nil
	}

	if err != nil {
		return nil, err
	}

	var d, p map[string]any
	err = json.Unmarshal(dest, &d)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(patch, &p)
	if err != nil {
		return nil, err
	}

	err = patchDict(d, p)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&d)
}

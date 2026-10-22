//go:build generator

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadOpenAPIPatched(t *testing.T) {
	dir := t.TempDir()
	specFile := filepath.Join(dir, "openapi.json")
	patchFile := filepath.Join(dir, "openapi_patch.yaml")

	spec := `{
		"b": {
			"max": 9223372036854775807,
			"min": -9223372036854775808,
			"umax": 18446744073709551615,
			"float": 1.5,
			"len": 10,
			"list": [1, 2]
		},
		"a": "<b>&</b>"
	}`
	patch := `
b:
  len: 1024
  list: [3]
  added: true
`
	require.NoError(t, os.WriteFile(specFile, []byte(spec), writeMode))
	require.NoError(t, os.WriteFile(patchFile, []byte(patch), writeMode))

	expected := `{"a":"<b>&</b>","b":{"added":true,"float":1.5,"len":1024,"list":[3],` +
		`"max":9223372036854775807,"min":-9223372036854775808,"umax":18446744073709551615}}` + "\n"

	patched, err := readOpenAPIPatched(specFile, patchFile)
	require.NoError(t, err)
	assert.Equal(t, expected, string(patched))

	// The generator writes the result back, so patching it again must be a no-op
	require.NoError(t, os.WriteFile(specFile, patched, writeMode))
	again, err := readOpenAPIPatched(specFile, patchFile)
	require.NoError(t, err)
	assert.Equal(t, expected, string(again))
}

//go:build generator

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadOpenAPIPatched(t *testing.T) {
	applyPatch := func(t *testing.T, spec string, patch []byte) []byte {
		t.Helper()

		dir := t.TempDir()
		specFile := filepath.Join(dir, "openapi.json")
		patchFile := filepath.Join(dir, "openapi_patch.yaml")
		require.NoError(t, os.WriteFile(specFile, []byte(spec), writeMode))
		require.NoError(t, os.WriteFile(patchFile, patch, writeMode))

		patched, err := readOpenAPIPatched(specFile, patchFile)
		require.NoError(t, err)

		// The generator writes the result back, so patching it again must be a no-op.
		require.NoError(t, os.WriteFile(specFile, patched, writeMode))
		again, err := readOpenAPIPatched(specFile, patchFile)
		require.NoError(t, err)
		assert.Equal(t, string(patched), string(again))
		return patched
	}

	t.Run("preserves values and is idempotent", func(t *testing.T) {
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
		patch := []byte(`
b:
  len: 1024
  list: [3]
  added: true
`)

		expected := `{"a":"<b>&</b>","b":{"added":true,"float":1.5,"len":1024,"list":[3],` +
			`"max":9223372036854775807,"min":-9223372036854775808,"umax":18446744073709551615}}` + "\n"

		patched := applyPatch(t, spec, patch)
		assert.Equal(t, expected, string(patched))
	})

	t.Run("Google PSC approval IPv4 format", func(t *testing.T) {
		const requestSchema = "ServicePrivatelinkGoogleConnectionCreateRequestBody"
		spec := `{"components":{"schemas":{"` + requestSchema + `":{
			"type":"object",
			"properties":{"user_ip_address":{"type":"string","maxLength":15,"description":"Endpoint address"}},
			"required":["user_ip_address"]
		}}}}`
		patch, err := os.ReadFile(filepath.Join("..", "openapi_patch.yaml"))
		require.NoError(t, err)

		patched := applyPatch(t, spec, patch)
		var doc struct {
			Components struct {
				Schemas map[string]json.RawMessage `json:"schemas"`
			} `json:"components"`
		}
		require.NoError(t, json.Unmarshal(patched, &doc))
		require.JSONEq(t, `{
			"type":"object",
			"properties":{"user_ip_address":{"type":"string","format":"ipv4","maxLength":15,"description":"Endpoint address"}},
			"required":["user_ip_address"]
		}`, string(doc.Components.Schemas[requestSchema]))
	})
}

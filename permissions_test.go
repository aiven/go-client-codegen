package aiven

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermissions_ProjectGet(t *testing.T) {
	permissions, err := Permissions()
	require.NoError(t, err)
	require.NotNil(t, permissions)

	projectGetPerms, exists := permissions["ProjectGet"]
	require.True(t, exists, "ProjectGet should exist in permissions map")

	expectedPerms := []string{
		"admin",
		"project:services:read",
		"read_only",
		"role:organization:admin",
		"role:services:maintenance",
		"role:services:recover",
		"service:secrets:read",
	}

	assert.ElementsMatch(t, expectedPerms, projectGetPerms, "ProjectGet permissions should match expected values")
}

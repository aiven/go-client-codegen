package aiven

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermissions_ServiceGet(t *testing.T) {
	permissions := Permissions()
	require.NotNil(t, permissions)

	serviceGetPerms, exists := permissions["ServiceGet"]
	require.True(t, exists, "ServiceGet should exist in permissions map")

	expectedPerms := []string{
		"project:services:read",
	}

	assert.ElementsMatch(t, expectedPerms, serviceGetPerms, "ServiceGet permissions should match expected values")
}

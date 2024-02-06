// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	token := os.Getenv("AIVEN_TOKEN")
	if token == "" {
		t.Skip("token is required for the test")
	}

	c, err := NewClient(DebugOpt(true))
	require.NoError(t, err)

	ctx := context.Background()
	tokens, err := c.AccessTokenList(ctx)
	require.NoError(t, err)

	found := 0

	for _, to := range tokens {
		if strings.HasPrefix(token, to.TokenPrefix) {
			found++
		}
	}

	assert.Equal(t, 1, found)
}

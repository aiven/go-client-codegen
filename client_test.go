// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aiven/go-client-codegen/handler/service"
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

func TestServiceCreate(t *testing.T) {
	// Creates a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/project/foo/service", r.URL.Path)

		// Validates request
		expectIn := new(service.ServiceCreateIn)
		err := json.NewDecoder(r.Body).Decode(expectIn)
		assert.NoError(t, err)
		assert.Equal(t, "foo", expectIn.ServiceName)
		assert.Equal(t, "kafka", expectIn.ServiceType)

		// Creates response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"service": {"plan": "wow", "state": "RUNNING"}}`))
		require.NoError(t, err)
	}))
	defer server.Close()

	// Points a new client to the server url
	c, err := NewClient(TokenOpt("token"), HostOpt(server.URL))
	require.NotNil(t, c)
	require.NoError(t, err)

	// Makes create request
	in := &service.ServiceCreateIn{
		ServiceName: "foo",
		ServiceType: "kafka",
	}
	out, err := c.ServiceCreate(context.Background(), "foo", in)
	require.NoError(t, err)
	require.NotNil(t, out)
	assert.Equal(t, "wow", out.Plan)
	assert.Equal(t, "RUNNING", out.State)
}

// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNotFound(t *testing.T) {
	cases := []struct {
		name   string
		err    error
		expect bool
	}{
		{
			name: "not found",
			err: Error{
				Message: "Not Found",
				Status:  404,
			},
			expect: true,
		},
		{
			name: "conflict",
			err: Error{
				Message: "already exists",
				Status:  409,
			},
			expect: false,
		},
		{
			name:   "nil",
			expect: false,
		},
	}

	for _, opt := range cases {
		t.Run(opt.name, func(t *testing.T) {
			assert.Equal(t, opt.expect, IsNotFound(opt.err))
		})
	}
}

func TestIsNotFoundIntegration(t *testing.T) {
	token := os.Getenv("AIVEN_TOKEN")
	if token == "" {
		t.Skip("token is required for the test")
	}

	c, err := NewClient()
	require.NoError(t, err)

	ctx := context.Background()
	out, err := c.AccountGet(ctx, "does_not_exist")
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.True(t, IsNotFound(err))
}

func TestIsAlreadyExists(t *testing.T) {
	cases := []struct {
		name   string
		err    error
		expect bool
	}{
		{
			name: "not found",
			err: Error{
				Message: "Not Found",
				Status:  404,
			},
			expect: false,
		},
		{
			name: "already exists",
			err: Error{
				Message: "already exists",
				Status:  409,
			},
			expect: true,
		},
		{
			name: "conflict",
			err: Error{
				Message: "conflict",
				Status:  409,
			},
			expect: false,
		},
		{
			name:   "nil",
			expect: false,
		},
	}

	for _, opt := range cases {
		t.Run(opt.name, func(t *testing.T) {
			assert.Equal(t, opt.expect, IsAlreadyExists(opt.err))
		})
	}
}

// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

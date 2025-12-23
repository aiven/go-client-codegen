// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	require.Error(t, err)
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

func TestFromResponse(t *testing.T) {
	cases := []struct {
		name           string
		operationID    string
		statusCode     int
		body           []byte
		expectErr      error
		expectStringer string
		expectBytes    []byte
	}{
		{
			name:        "404",
			operationID: "UserAuth",
			statusCode:  http.StatusNotFound,
			body: []byte(`{
			  "message": "Account does not exist",
			  "errors": [
				{
				  "error_code": "account_not_found",
				  "message": "Account does not exist",
				  "status": 404
				}
			  ]
			}`),
			expectStringer: `[404 UserAuth]: Account does not exist`,
			expectBytes:    nil,
			expectErr: Error{
				Message: "Account does not exist",
				Errors: []any{map[string]any{
					"error_code": "account_not_found",
					"message":    "Account does not exist",
					"status":     float64(404),
				}},
				OperationID: "UserAuth",
				Status:      http.StatusNotFound,
			},
		},
		{
			name:        "200",
			operationID: "UserAuth",
			statusCode:  http.StatusOK,
			body:        []byte(`{"success": True}`),
			expectBytes: []byte(`{"success": True}`),
			expectErr:   nil,
		},
		{
			name:           "Invalid json",
			operationID:    "UserAuth",
			statusCode:     http.StatusInternalServerError,
			body:           []byte(`unknown error`),
			expectBytes:    nil,
			expectStringer: `[500 UserAuth]: unknown error`,
			expectErr: Error{
				OperationID: "UserAuth",
				Message:     "unknown error",
				Status:      http.StatusInternalServerError,
			},
		},
		{
			name:           "Valid json, Error fields type mismatch",
			operationID:    "UserAuth",
			statusCode:     http.StatusBadRequest,
			body:           []byte(`{"message": {"key": "value"}}`), // message is not string
			expectStringer: `[400 UserAuth]: json: cannot unmarshal object into Go struct field Error.message of type string`,
			expectBytes:    nil,
			expectErr: Error{
				OperationID: "UserAuth",
				Message:     "json: cannot unmarshal object into Go struct field Error.message of type string",
				Status:      http.StatusBadRequest,
			},
		},
		{
			name:        "add errors, because message is different",
			operationID: "UserAuth",
			statusCode:  http.StatusNotFound,
			body: []byte(`{
			  "message": "Oh no!",
			  "errors": [
				{
				  "error_code": "account_not_found",
				  "message": "Account does not exist",
				  "status": 404
				}
			  ]
			}`),
			expectStringer: `[404 UserAuth]: Oh no! ([{"error_code":"account_not_found","message":"Account does not exist","status":404}])`,
			expectBytes:    nil,
			expectErr: Error{
				Message: "Oh no!",
				Errors: []any{map[string]any{
					"error_code": "account_not_found",
					"message":    "Account does not exist",
					"status":     float64(404),
				}},
				OperationID: "UserAuth",
				Status:      http.StatusNotFound,
			},
		},
	}

	for _, opt := range cases {
		t.Run(opt.name, func(t *testing.T) {
			rsp := &http.Response{
				StatusCode: opt.statusCode,
				Body:       io.NopCloser(bytes.NewReader(opt.body)),
			}

			b, err := fromResponse(opt.operationID, rsp)
			assert.Equal(t, opt.expectBytes, b)
			assert.Empty(t, cmp.Diff(opt.expectErr, err))
			if err != nil {
				assert.Equal(t, opt.expectStringer, err.Error())
			}
		})
	}
}

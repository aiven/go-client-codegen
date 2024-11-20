// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aiven/go-client-codegen/handler/clickhouse"
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
	var callCount int64

	// Creates a test server
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/v1/project/aiven-project/service",
		func(w http.ResponseWriter, r *http.Request) {
			// Validates request
			expectIn := new(service.ServiceCreateIn)
			err := json.NewDecoder(r.Body).Decode(expectIn)
			assert.NoError(t, err)
			assert.Equal(t, "my-clickhouse", expectIn.ServiceName)
			assert.Equal(t, "clickhouse", expectIn.ServiceType)
			assert.Regexp(t, `go-client-codegen/[0-9\.]+ unit-test`, r.Header["User-Agent"])

			// Creates response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`{"service": {"plan": "wow", "state": "RUNNING"}}`))
			require.NoError(t, err)
			atomic.AddInt64(&callCount, 1)
		},
	)
	mux.HandleFunc(
		"/v1/project/aiven-project/service/my-clickhouse/clickhouse/query/stats",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.URL.RawQuery, "limit=1&order_by=max_time%3Aasc")

			// Creates response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"queries": [{"calls": 1}]}`))
			require.NoError(t, err)
			atomic.AddInt64(&callCount, 1)
		},
	)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Points a new client to the server url
	c, err := NewClient(TokenOpt("token"), HostOpt(server.URL), UserAgentOpt("unit-test"))
	require.NotNil(t, c)
	require.NoError(t, err)

	// Makes create request
	in := &service.ServiceCreateIn{
		ServiceName: "my-clickhouse",
		ServiceType: "clickhouse",
	}

	ctx := context.Background()
	project := "aiven-project"
	out, err := c.ServiceCreate(ctx, project, in)
	require.NoError(t, err)
	require.NotNil(t, out)
	assert.Equal(t, "wow", out.Plan)
	assert.Equal(t, service.ServiceStateTypeRunning, out.State)

	// Validates query params
	stats, err := c.ServiceClickHouseQueryStats(
		ctx, project, in.ServiceName,
		clickhouse.ServiceClickHouseQueryStatsLimit(1),
		clickhouse.ServiceClickHouseQueryStatsOrderByType(clickhouse.OrderByTypeMaxTimeAsc),
	)
	require.NoError(t, err)
	assert.Len(t, stats, 1)

	// All calls are received
	assert.EqualValues(t, 2, callCount)
}

func TestServiceCreateErrorsRetries(t *testing.T) {
	cases := []struct {
		Name         string
		ResponseBody string
		ErrorExpect  string
		RetryMax     int
		CallsExpect  int
	}{
		{
			Name:         "message field only",
			ResponseBody: `{"message": "Internal Server Error"}`,
			ErrorExpect:  "[500 ServiceCreate]: Internal Server Error",
			RetryMax:     6,
			CallsExpect:  7,
		},
		{
			Name:         "with errors field, list",
			ResponseBody: `{"message": "Something went wrong", "errors": ["oh!", "no!"]}`,
			ErrorExpect:  `[500 ServiceCreate]: Something went wrong: ["oh!","no!"]`,
			RetryMax:     1,
			CallsExpect:  2,
		},
		{
			Name:         "with errors field, string",
			ResponseBody: `{"message": "Something went wrong", "errors": "wow!"}`,
			ErrorExpect:  `[500 ServiceCreate]: Something went wrong: "wow!"`,
			RetryMax:     1,
			CallsExpect:  2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var callsActual int64

			// Creates a test server
			mux := http.NewServeMux()
			mux.HandleFunc(
				"/v1/project/aiven-project/service",
				func(w http.ResponseWriter, _ *http.Request) {
					// Creates response
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write([]byte(tt.ResponseBody))
					require.NoError(t, err)
					atomic.AddInt64(&callsActual, 1)
				},
			)

			server := httptest.NewServer(mux)

			// Points a new client to the server url
			c, err := NewClient(
				TokenOpt("token"),
				UserAgentOpt("unit-test"),
				HostOpt(server.URL),
				RetryMaxOpt(tt.RetryMax),
				RetryWaitMinOpt(1*time.Millisecond),
				RetryWaitMaxOpt(3*time.Millisecond),
				DebugOpt(false),
			)
			require.NotNil(t, c)
			require.NoError(t, err)

			// Makes create request
			in := &service.ServiceCreateIn{
				ServiceName: "my-clickhouse",
				ServiceType: "clickhouse",
			}

			ctx := context.Background()
			project := "aiven-project"
			out, err := c.ServiceCreate(ctx, project, in)
			assert.Nil(t, out)
			assert.Equal(t, err.Error(), tt.ErrorExpect)

			// All calls are received
			assert.EqualValues(t, tt.CallsExpect, callsActual)
			server.Close()
		})
	}
}

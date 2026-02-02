// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
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
			assert.NoError(t, err)
			atomic.AddInt64(&callCount, 1)
		},
	)
	mux.HandleFunc(
		"/v1/project/aiven-project/service/my-clickhouse/clickhouse/query/stats",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "limit=1&order_by=max_time%3Aasc", r.URL.RawQuery)

			// Creates response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"queries": [{"calls": 1}]}`))
			assert.NoError(t, err)
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
			ErrorExpect:  `[500 ServiceCreate]: Something went wrong (["oh!","no!"])`,
			RetryMax:     1,
			CallsExpect:  2,
		},
		{
			Name:         "with errors field, string",
			ResponseBody: `{"message": "Something went wrong", "errors": "wow!"}`,
			ErrorExpect:  `[500 ServiceCreate]: Something went wrong ("wow!")`,
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
					assert.NoError(t, err)
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
			out, err := c.ServiceCreate(ctx, "aiven-project", in)
			assert.Nil(t, out)
			assert.Equal(t, err.Error(), tt.ErrorExpect)

			// All calls are received
			assert.EqualValues(t, tt.CallsExpect, callsActual)
			server.Close()
		})
	}
}

func TestFmtQuery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		operationID string
		query       [][2]string
		want        string
	}{
		{
			name:        "With no params",
			operationID: "TestOperation",
			query:       nil,
			want:        "limit=999",
		},
		{
			name:        "With existing params",
			operationID: "TestOperation",
			query: [][2]string{
				{"foo", "bar"},
				{"baz", "qux"},
			},
			want: "baz=qux&foo=bar&limit=999",
		},
		{
			name:        "With custom limit",
			operationID: "TestOperation",
			query: [][2]string{
				{"limit", "50"},
			},
			want: "limit=50",
		},
		{
			name:        "Ignored operation",
			operationID: "ServiceKafkaQuotaDescribe",
			query: [][2]string{
				{"foo", "bar"},
			},
			want: "foo=bar",
		},
		{
			name:        "Multiple parameters with same key",
			operationID: "TestOperation",
			query: [][2]string{
				{"tag", "v1"},
				{"tag", "v2"},
			},
			want: "limit=999&tag=v1&tag=v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := fmtQuery(tt.operationID, tt.query...)
			if got != tt.want {
				t.Errorf("fmtQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceIntegrationEndpointGet(t *testing.T) {
	var callCount int64

	// Creates a test server
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/v1/project/{project}/integration_endpoint/{integration_endpoint_id}",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Regexp(t, `go-client-codegen/[0-9\.]+ unit-test`, r.Header["User-Agent"])
			assert.Equal(t, "/v1/project/aiven-endpoint-project/integration_endpoint/foo?include_secrets=true&limit=999", r.RequestURI)

			// Creates response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"service_integration_endpoint": {"endpoint_name": "wow"}}`))
			assert.NoError(t, err)
			atomic.AddInt64(&callCount, 1)
		},
	)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Points a new client to the server url
	c, err := NewClient(TokenOpt("token"), HostOpt(server.URL), UserAgentOpt("unit-test"))
	require.NotNil(t, c)
	require.NoError(t, err)

	ctx := context.Background()
	project := "aiven-endpoint-project"
	out, err := c.ServiceIntegrationEndpointGet(ctx, project, "foo", service.ServiceIntegrationEndpointGetIncludeSecrets(true))
	require.NoError(t, err)
	require.NotNil(t, out)
	assert.Equal(t, "wow", out.EndpointName)

	// All calls are received
	assert.EqualValues(t, 1, callCount)
}

type doerFunc func(*http.Request) (*http.Response, error)

func (f doerFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestSingleFlightDeduplication(t *testing.T) {
	// This test checks that EnableSingleFlight deduplicates concurrent identical requests:
	// the first call reaches the HTTP Doer and blocks, and all other concurrent calls reuse
	// the same in-flight result without reaching the Doer themselves.
	//
	// We use a blocking Doer and testing/synctest to make the execution deterministic:
	// synctest.Wait() pauses the test until all goroutines in the bubble are durably blocked.
	synctest.Test(t, func(t *testing.T) {
		const concurrency = 20

		// 1) Create a client with a Doer that blocks until we release it.
		// releaseNow is called explicitly and also deferred to avoid leaving goroutines stuck if the test fails early.
		var releaseOnce sync.Once
		release := make(chan struct{})
		releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
		defer releaseNow()

		const body = `{"services":[{"service_name":"svc"}]}`

		entered := make(chan struct{})
		c, err := NewClient(
			TokenOpt("token"),
			HostOpt("http://example.com"),
			UserAgentOpt("unit-test"),
			DoerOpt(doerFunc(func(req *http.Request) (*http.Response, error) {
				// 2) Signal that the first request reached the Doer.
				// If a second request reaches the Doer, this will panic (close of closed channel).
				close(entered)
				// 3) Block the first request so the rest of the goroutines can race into singleflight.
				<-release
				return &http.Response{
					StatusCode:    http.StatusOK,
					Header:        http.Header{},
					Body:          io.NopCloser(strings.NewReader(body)),
					ContentLength: int64(len(body)),
					Request:       req,
				}, nil
			})),
			EnableSingleFlightOpt(true),
		)
		require.NoError(t, err)

		ctx := t.Context()

		// 4) Start the first request and wait until it reaches the Doer (and blocks).
		var firstErr error
		var firstRes []service.ServiceOut
		firstDone := make(chan struct{})
		go func() {
			defer close(firstDone)
			firstRes, firstErr = c.ServiceList(ctx, "test-project")
		}()

		<-entered
		synctest.Wait()

		// 5) Start concurrent requests. They mustn't reach the Doer while the first is in-flight.
		res := make([][]service.ServiceOut, concurrency)
		errs := make([]error, concurrency)
		var wg sync.WaitGroup
		for i := range concurrency {
			wg.Go(func() {
				res[i], errs[i] = c.ServiceList(ctx, "test-project")
			})
		}
		synctest.Wait()

		// 6) Release the first request and wait for every goroutine to finish.
		releaseNow()

		<-firstDone
		require.NoError(t, firstErr)
		wg.Wait()

		// 7) All callers should see the same result.
		for _, err := range errs {
			require.NoError(t, err)
		}
		for _, r := range res {
			require.Equal(t, firstRes, r)
		}
	})
}

func TestSingleFlightDeduplicationHTTPServer(t *testing.T) {
	const concurrency = 20

	var callCount int64

	var releaseOnce sync.Once
	release := make(chan struct{})
	releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
	defer releaseNow()

	entered := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc(
		"/v1/project/test-project/service",
		func(w http.ResponseWriter, _ *http.Request) {
			if atomic.AddInt64(&callCount, 1) == 1 {
				close(entered)
			}

			<-release

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"services":[{"service_name":"svc"}]}`))
		},
	)

	server := httptest.NewServer(mux)
	defer server.Close()

	c, err := NewClient(
		TokenOpt("token"),
		HostOpt(server.URL),
		UserAgentOpt("unit-test"),
		DoerOpt(server.Client()),
		EnableSingleFlightOpt(true),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	start := make(chan struct{})
	var ready sync.WaitGroup
	ready.Add(concurrency)

	type callResult struct {
		res []service.ServiceOut
		err error
	}

	results := make(chan callResult, concurrency)
	var wg sync.WaitGroup
	for range concurrency {
		wg.Go(func() {
			ready.Done()
			<-start
			res, err := c.ServiceList(ctx, "test-project")
			results <- callResult{res: res, err: err}
		})
	}

	ready.Wait()
	close(start)

	select {
	case <-entered:
	case <-time.After(5 * time.Second):
		t.Fatal("server wasn't called")
	}

	// Give other goroutines a chance to enter the client call while the server is still blocked.
	time.Sleep(50 * time.Millisecond)

	releaseNow()

	wg.Wait()
	close(results)

	var first []service.ServiceOut
	for r := range results {
		require.NoError(t, r.err)
		if first == nil {
			first = r.res
			continue
		}
		require.Equal(t, first, r.res)
	}

	require.Less(t, atomic.LoadInt64(&callCount), int64(concurrency))
}

// TestSingleFlightDifferentProjects tests that singleflight doesn't merge different paths.
func TestSingleFlightDifferentProjects(t *testing.T) {
	// This test checks that EnableSingleFlight does not merge requests with different paths.
	// Here we use two different project names, so the generated URL path differs.
	synctest.Test(t, func(t *testing.T) {
		// 1) Arrange: a Doer that blocks so requests overlap in time.
		var releaseOnce sync.Once
		release := make(chan struct{})
		releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
		defer releaseNow()

		const (
			bodyA = `{"services":[{"service_name":"svc-a"}]}`
			bodyB = `{"services":[{"service_name":"svc-b"}]}`
		)

		enteredA := make(chan struct{})
		enteredB := make(chan struct{})

		c, err := NewClient(
			TokenOpt("token"),
			HostOpt("http://example.com"),
			UserAgentOpt("unit-test"),
			DoerOpt(doerFunc(func(req *http.Request) (*http.Response, error) {
				var body string
				switch req.URL.Path {
				case "/v1/project/project-a/service":
					// If the same project request reaches the Doer twice, panic (close of closed channel).
					close(enteredA)
					body = bodyA
				case "/v1/project/project-b/service":
					close(enteredB)
					body = bodyB
				default:
					return nil, errors.New("unexpected request path: " + req.URL.Path)
				}

				<-release
				return &http.Response{
					StatusCode:    http.StatusOK,
					Header:        http.Header{},
					Body:          io.NopCloser(strings.NewReader(body)),
					ContentLength: int64(len(body)),
					Request:       req,
				}, nil
			})),
			EnableSingleFlightOpt(true),
		)
		require.NoError(t, err)

		ctx := t.Context()

		type listResult struct {
			res []service.ServiceOut
			err error
		}

		// 2) Start the first request and wait until it reaches the Doer (and blocks).
		projectACh := make(chan listResult, 1)
		go func() {
			res, err := c.ServiceList(ctx, "project-a")
			projectACh <- listResult{res: res, err: err}
		}()

		<-enteredA
		synctest.Wait()

		// 3) Start the second request while the first is still in-flight. It must reach the Doer too.
		projectBCh := make(chan listResult, 1)
		go func() {
			res, err := c.ServiceList(ctx, "project-b")
			projectBCh <- listResult{res: res, err: err}
		}()

		synctest.Wait()
		select {
		case <-enteredB:
		default:
			t.Fatal("project-b request did not reach the Doer; likely deduplicated with project-a")
		}

		// 4) Release both requests and verify results.
		releaseNow()

		projectARes := <-projectACh
		projectBRes := <-projectBCh
		require.NoError(t, projectARes.err)
		require.NoError(t, projectBRes.err)
		require.Len(t, projectARes.res, 1)
		require.Len(t, projectBRes.res, 1)
		require.NotEqual(t, projectARes.res, projectBRes.res)
		require.Equal(t, "svc-a", projectARes.res[0].ServiceName)
		require.Equal(t, "svc-b", projectBRes.res[0].ServiceName)
	})
}

func TestSingleFlightMethodIsolation(t *testing.T) {
	// This test checks that singleflight never merges requests with different HTTP methods.
	// We run several concurrent requests for the same path+query, but with different methods.
	// Each method should start its own in-flight request (and deduplicate only within itself).
	synctest.Test(t, func(t *testing.T) {
		const copies = 5

		// A blocking Doer makes all requests overlap in time.
		var releaseOnce sync.Once
		release := make(chan struct{})
		releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
		defer releaseNow()

		methods := []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodOptions,
			http.MethodTrace,
		}

		entered := make(map[string]chan struct{}, len(methods))
		for _, method := range methods {
			entered[method] = make(chan struct{})
		}

		queryLimit := [2]string{"limit", "1"}

		d := &aivenClient{
			Host:               "http://example.com",
			UserAgent:          "unit-test",
			Token:              "token",
			EnableSingleFlight: true,
			doer: doerFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Path != "/v1/test" {
					return nil, errors.New("unexpected request path: " + req.URL.Path)
				}
				if req.URL.RawQuery != "limit=1" {
					return nil, errors.New("unexpected request query: " + req.URL.RawQuery)
				}

				ch, ok := entered[req.Method]
				if !ok {
					return nil, errors.New("unexpected request method: " + req.Method)
				}

				// If the same method reaches the Doer twice, panic (close of closed channel).
				close(ch)

				<-release
				body := req.Method
				return &http.Response{
					StatusCode:    http.StatusOK,
					Header:        http.Header{},
					Body:          io.NopCloser(strings.NewReader(body)),
					ContentLength: int64(len(body)),
					Request:       req,
				}, nil
			}),
		}

		ctx := t.Context()

		type callResult struct {
			method string
			body   []byte
			err    error
		}

		results := make(chan callResult, copies*len(methods))
		var wg sync.WaitGroup

		for _, method := range methods {
			for range copies {
				wg.Go(func() {
					b, err := d.Do(ctx, "op", method, "/v1/test", nil, queryLimit)
					results <- callResult{method: method, body: b, err: err}
				})
			}
		}

		synctest.Wait()

		// Ensure every method reached the Doer (i.e., didn't merge into another one).
		for _, method := range methods {
			select {
			case <-entered[method]:
			default:
				t.Errorf("%s request did not reach the Doer; likely merged into another method", method)
			}
		}

		releaseNow()

		wg.Wait()
		close(results)

		for r := range results {
			require.NoError(t, r.err)
			require.Equal(t, r.method, string(r.body))
		}
	})
}

package aiven

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aiven/go-client-codegen/handler/service"
)

type lockedWriter struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (w *lockedWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buf.Write(p)
}

func (w *lockedWriter) String() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buf.String()
}

func (w *lockedWriter) lineCount() int {
	return strings.Count(w.String(), "\n")
}

func newDebugLogClient(t *testing.T, host string, opts ...Option) (Client, *lockedWriter) {
	t.Helper()

	logBuf := &lockedWriter{}
	base := []Option{
		TokenOpt("test-token"),
		HostOpt(host),
		UserAgentOpt("unit-test"),
		DebugOpt(true),
		LoggerWriterOpt(logBuf),
		EnableSingleFlightOpt(true),
	}
	c, err := NewClient(append(base, opts...)...)
	require.NoError(t, err)
	return c, logBuf
}

func TestDebugLogsOncePerSuccessfulCall(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/access_token", func(w http.ResponseWriter, r *http.Request) {
		assert.Regexp(t, `go-client-codegen/[0-9\.]+ unit-test`, r.Header.Get("User-Agent"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"tokens":[]}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c, logs := newDebugLogClient(t, server.URL)

	_, err := c.AccessTokenList(t.Context())
	require.NoError(t, err)

	logText := logs.String()
	assert.Equal(t, 1, logs.lineCount())
	assert.Contains(t, logText, "AccessTokenList")
	assert.Contains(t, logText, "GET")
}

func TestDebugLogsOncePerErroredCall(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/access_token", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"bad request"}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c, logs := newDebugLogClient(t, server.URL)

	_, err := c.AccessTokenList(t.Context())
	require.Error(t, err)

	logText := logs.String()
	assert.Equal(t, 1, logs.lineCount())
	assert.Contains(t, logText, "AccessTokenList")
	assert.Contains(t, logText, "GET")
}

func TestDebugLogsOnceAcrossRetries(t *testing.T) {
	var calls int64

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/project/test-project/service", func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt64(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		if n < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"temporary failure"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"service":{"plan":"basic","state":"RUNNING"}}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c, logs := newDebugLogClient(t, server.URL,
		RetryMaxOpt(3),
		RetryWaitMinOpt(1*time.Millisecond),
		RetryWaitMaxOpt(5*time.Millisecond),
	)

	in := &service.ServiceCreateIn{
		ServiceName: "svc",
		ServiceType: "pg",
	}
	_, err := c.ServiceCreate(t.Context(), "test-project", in)
	require.NoError(t, err)
	require.EqualValues(t, 3, calls)

	logText := logs.String()
	assert.Equal(t, 1, logs.lineCount())
	assert.Contains(t, logText, "ServiceCreate")
	assert.Contains(t, logText, "POST")
}

func TestDebugLogsOncePerCallerWithSingleflight(t *testing.T) {
	// zerolog.ConsoleWriter wraps field names in ANSI color codes unless
	// NO_COLOR is set, which splits "shared=true" and breaks the count below.
	t.Setenv("NO_COLOR", "1")
	synctest.Test(t, func(t *testing.T) {
		const concurrency = 2
		const body = `{"services":[{"service_name":"svc"}]}`

		var callCount int64

		var releaseOnce sync.Once
		release := make(chan struct{})
		releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
		defer releaseNow()

		c, logs := newDebugLogClient(t, "http://example.com",
			DoerOpt(doerFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt64(&callCount, 1)
				<-release
				return &http.Response{
					StatusCode:    http.StatusOK,
					Header:        http.Header{},
					Body:          io.NopCloser(strings.NewReader(body)),
					ContentLength: int64(len(body)),
					Request:       req,
				}, nil
			})),
		)

		ctx := t.Context()
		errs := make([]error, concurrency)
		var wg sync.WaitGroup
		for i := range concurrency {
			wg.Go(func() {
				_, errs[i] = c.ServiceList(ctx, "test-project")
			})
		}

		// Every caller is now blocked either in the Doer or in singleflight.
		synctest.Wait()
		releaseNow()
		wg.Wait()

		for _, err := range errs {
			require.NoError(t, err)
		}
		require.EqualValues(t, 1, callCount)
		logText := logs.String()
		assert.Equal(t, concurrency, logs.lineCount())
		assert.Equal(t, concurrency, strings.Count(logText, "shared=true"))
		assert.Contains(t, logText, "ServiceList")
		assert.Contains(t, logText, "GET")
	})
}

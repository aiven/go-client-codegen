package aiven

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
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
	const concurrency = 2

	var callCount int64

	var releaseOnce sync.Once
	release := make(chan struct{})
	releaseNow := func() { releaseOnce.Do(func() { close(release) }) }
	defer releaseNow()

	entered := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/project/test-project/service", func(w http.ResponseWriter, _ *http.Request) {
		if atomic.AddInt64(&callCount, 1) == 1 {
			close(entered)
		}
		<-release

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"services":[{"service_name":"svc"}]}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c, logs := newDebugLogClient(t, server.URL)

	ctx := t.Context()
	start := make(chan struct{})
	var ready sync.WaitGroup
	ready.Add(concurrency)

	var wg sync.WaitGroup
	errs := make(chan error, concurrency)
	for range concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ready.Done()
			<-start
			_, err := c.ServiceList(ctx, "test-project")
			errs <- err
		}()
	}

	ready.Wait()
	close(start)

	select {
	case <-entered:
	case <-time.After(5 * time.Second):
		t.Fatal("server was not called")
	}

	// Grace window for the second caller to reach singleflight.Do and
	// join the in-flight call before we release the server. We can't
	// observe "parked in singleflight" from the outside; the
	// callCount == 1 assertion below self-validates dedup, so a flake
	// here fails loudly rather than silently passing.
	time.Sleep(50 * time.Millisecond)
	releaseNow()
	wg.Wait()
	close(errs)

	for err := range errs {
		require.NoError(t, err)
	}
	require.EqualValues(t, 1, callCount)
	logText := logs.String()
	assert.Equal(t, concurrency, logs.lineCount())
	assert.Contains(t, logText, "ServiceList")
	assert.Contains(t, logText, "GET")
}

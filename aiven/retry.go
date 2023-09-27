package aiven

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-retryablehttp"
)

// checkRetry does basic retries (>=500 && != 501, timeouts, connection errors)
// Plus custom retries, see isRetryable
// If ErrorPropagatedRetryPolicy returns error it is either >=500
// or things you can't retry like an invalid protocol scheme
// Suspends errors, cause that's what retryablehttp.DefaultRetryPolicy does
func checkRetry(ctx context.Context, rsp *http.Response, err error) (bool, error) {
	shouldRetry, err := retryablehttp.ErrorPropagatedRetryPolicy(ctx, rsp, err)
	return shouldRetry || err == nil && isRetryable(rsp), nil
}

// isRetryable
// 501 — some resources like kafka topic and kafka connector may return 501. Which means retry later
// 417 — has dependent resource pending: Application deletion forbidden because it has 1 deployment(s).
// 408 — dependent server time out
// 404 — see retryableChecks
func isRetryable(rsp *http.Response) bool {
	// This might happen in tests only
	if rsp.Request == nil {
		return false
	}

	switch rsp.StatusCode {
	case http.StatusRequestTimeout, http.StatusNotImplemented:
		return true
	case http.StatusExpectationFailed:
		return rsp.Request.Method == "DELETE"
	case http.StatusNotFound:
		// We need to restore the body
		body := rsp.Body
		defer body.Close()

		// Shouldn't be there much of data, ReadAll is ok
		b, err := io.ReadAll(body)
		if err != nil {
			return false
		}

		// Restores the body
		rsp.Body = io.NopCloser(bytes.NewReader(b))

		// Error checks
		s := string(b)
		for _, c := range retryableChecks {
			if c(rsp.Request.Method, s) {
				return true
			}
		}
	}
	return false
}

var retryableChecks = []func(meth, body string) bool{
	isServiceLagError,
	isUserError,
}

var (
	reServiceNotFound = regexp.MustCompile(`Service \S+ does not exist`)
	reUserNotFound    = regexp.MustCompile(`User (avnadmin|root) with component main not found`)
)

// isServiceLagError service is might be ready, but there is a lag that need to wait for ending
func isServiceLagError(meth, body string) bool {
	return meth == "POST" && reServiceNotFound.MatchString(body)
}

// isUserError an internal unknown error
func isUserError(_, body string) bool {
	return reUserNotFound.MatchString(body)
}

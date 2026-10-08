// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"io"
	"time"
)

// Option is a function that configures the client.
type Option func(*aivenClient)

// DebugOpt whether should the client run in debug mode.
// For instance, to output debug information
func DebugOpt(debug bool) Option {
	return func(d *aivenClient) {
		d.Debug = debug
	}
}

// LoggerWriterOpt sets the writer the debug logger emits to.
//
// Defaults to os.Stderr. Passing a custom writer is useful when the caller
// wants to route the client's request logs somewhere else — for example,
// through a CLI spinner that interleaves logs above its animation without
// clobbering it, or to a file. Has no effect unless DebugOpt(true) is also
// set.
//
// The writer must be safe for concurrent Write calls: the client fires
// requests from arbitrary goroutines, and both the per-attempt hook and
// the per-operation summary may call Write from different goroutines at
// the same time. os.Stderr and *os.File satisfy this; a bare bytes.Buffer
// does not — wrap it in a mutex if needed.
func LoggerWriterOpt(w io.Writer) Option {
	return func(d *aivenClient) {
		d.loggerWriter = w
	}
}

// UserAgentOpt sets User-Agent header
func UserAgentOpt(userAgent string) Option {
	return func(d *aivenClient) {
		d.UserAgent = userAgent
	}
}

// TokenOpt runs the client with the given token
func TokenOpt(token string) Option {
	return func(d *aivenClient) {
		d.Token = token
	}
}

// HostOpt API host url
func HostOpt(host string) Option {
	return func(d *aivenClient) {
		d.Host = host
	}
}

// DoerOpt replaces underlying http client in aivenClient
func DoerOpt(doer Doer) Option {
	return func(d *aivenClient) {
		d.doer = doer
	}
}

// RetryMaxOpt sets the maximum number of retries
func RetryMaxOpt(retryMax int) Option {
	return func(d *aivenClient) {
		d.RetryMax = retryMax
	}
}

// RetryWaitMinOpt sets the minimum wait time between retries
func RetryWaitMinOpt(retryWaitMin time.Duration) Option {
	return func(d *aivenClient) {
		d.RetryWaitMin = retryWaitMin
	}
}

// RetryWaitMaxOpt sets the maximum wait time between retries
func RetryWaitMaxOpt(retryWaitMax time.Duration) Option {
	return func(d *aivenClient) {
		d.RetryWaitMax = retryWaitMax
	}
}

// EnableSingleFlightOpt enables singleflight for deduplicating concurrent identical requests
func EnableSingleFlightOpt(enableSingleFlight bool) Option {
	return func(d *aivenClient) {
		d.EnableSingleFlight = enableSingleFlight
	}
}

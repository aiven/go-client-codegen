// Package aiven provides a client for interacting with the Aiven API.
package aiven

import "time"

// Option is a function that configures the client.
type Option func(*aivenClient)

// DebugOpt whether should the client run in debug mode.
// For instance, to output debug information
func DebugOpt(debug bool) Option {
	return func(d *aivenClient) {
		d.Debug = debug
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

// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

var errTokenIsRequired = errors.New(
	"token is required. See https://api.aiven.io/doc/#section/Get-started/Authentication",
)

// Doer aka http.Client
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient creates a new Aiven client.
func NewClient(opts ...Option) (Client, error) {
	d := new(aivenClient)

	err := envconfig.Process("", d)
	if err != nil {
		return nil, err
	}

	// User settings
	for _, opt := range opts {
		opt(d)
	}

	// When DoerOpt is not applied
	if d.doer == nil {
		c := retryablehttp.NewClient()
		c.RetryMax = d.RetryMax
		c.RetryWaitMin = d.RetryWaitMin
		c.RetryWaitMax = d.RetryWaitMax
		c.CheckRetry = checkRetry

		// Disables retryablehttp logger which outputs a lot of debug information
		c.Logger = nil

		// By default, when retryablehttp gets 500 (or any error),
		// it doesn't return the body which might contain useful information.
		// Instead, it returns `giving up after %d attempt(s)` for the given url and method.
		c.ErrorHandler = retryablehttp.PassthroughErrorHandler
		d.doer = c.StandardClient()
	}

	if d.Debug {
		out := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		d.logger = zerolog.New(out).With().Timestamp().Logger()
	}

	if d.Token == "" {
		return nil, errTokenIsRequired
	}

	// Removes trailing / so it is easier later Host + URL
	d.Host = strings.TrimSuffix(d.Host, "/")

	// Formats the user agent
	d.UserAgent = fmt.Sprintf(
		"go-client-codegen/%s %s",
		strings.TrimLeft(Version(), "v"),
		strings.TrimSpace(d.UserAgent),
	)

	return newClient(d), nil
}

// Retry settings explanation:
// Default values (retryWaitMin = 1s, retryWaitMax = 30s, retryMax = 4)
// Changed values (retryWaitMin = 2s, retryWaitMax = 15s, retryMax = 6)
//
// Default values       | Changed values
// Run  Seconds         | Run  Seconds
// 0    0.000           | 0    0.000
// 1    1.000           | 1    2.000
// 2    3.000           | 2    6.000
// 3    7.000           | 3    14.000
// 4    15.000          | 4    15.000 (capped at retryWaitMax)
//
//	| 5    15.000 (capped at retryWaitMax)
//	| 6    15.000 (capped at retryWaitMax)
//
// Maximum wait time if all attempts fail:
// Default values: 26 seconds
// Changed values: 67 seconds
type aivenClient struct {
	Host         string        `envconfig:"AIVEN_WEB_URL" default:"https://api.aiven.io"`
	UserAgent    string        `envconfig:"AIVEN_USER_AGENT" default:"aiven-go-client/v3"`
	Token        string        `envconfig:"AIVEN_TOKEN"`
	Debug        bool          `envconfig:"AIVEN_DEBUG"`
	RetryMax     int           `envconfig:"AIVEN_CLIENT_RETRY_MAX" default:"6"`
	RetryWaitMin time.Duration `envconfig:"AIVEN_CLIENT_RETRY_WAIT_MIN" default:"2s"`
	RetryWaitMax time.Duration `envconfig:"AIVEN_CLIENT_RETRY_WAIT_MAX" default:"15s"`
	logger       zerolog.Logger
	doer         Doer
}

// OperationIDKey is the key used to store the operation ID in the context.
type OperationIDKey struct{}

func (d *aivenClient) Do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) ([]byte, error) {
	ctx = context.WithValue(ctx, OperationIDKey{}, operationID)
	var rsp *http.Response
	var err error

	if d.Debug {
		start := time.Now()
		defer func() {
			end := time.Since(start)

			var event *zerolog.Event
			if err != nil {
				event = d.logger.Error().Err(err)
			} else {
				event = d.logger.Info().Str("status", rsp.Status)
			}

			event.Ctx(ctx).
				Stringer("duration", end).
				Str("operationID", operationID).
				Str("method", method).
				Str("path", path).
				Str("query", fmtQuery(operationID, query...)).
				Send()
		}()
	}

	rsp, err = d.do(ctx, operationID, method, path, in, query...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = multierror.Append(rsp.Body.Close()).ErrorOrNil()
	}()

	return fromResponse(operationID, rsp)
}

func (d *aivenClient) do(ctx context.Context, operationID, method, path string, in any, query ...[2]string) (*http.Response, error) {
	var body io.Reader

	if !(in == nil || isEmpty(in)) {
		b, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, d.Host+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", d.UserAgent)
	req.Header.Set("Authorization", "aivenv1 "+d.Token)
	req.URL.RawQuery = fmtQuery(operationID, query...)

	return d.doer.Do(req)
}

func isEmpty(a any) bool {
	v := reflect.ValueOf(a)
	return v.IsZero() || v.Kind() == reflect.Ptr && v.IsNil()
}

func fmtQuery(operationID string, query ...[2]string) string {
	q := make(url.Values)

	// Add all provided query parameters
	for _, v := range query {
		q.Add(v[0], v[1])
	}

	// Add default limit for GET requests if conditions are met
	const defaultLimit = "999"
	if shouldAddDefaultLimit(operationID, q) {
		// TODO: BAD hack to get around pagination in most cases
		// we should implement this properly at some point but for now
		// that should be its own issue
		q.Add("limit", defaultLimit)
	}

	return q.Encode()
}

// shouldAddDefaultLimit determines if default limit should be added
func shouldAddDefaultLimit(operationID string, q url.Values) bool {
	operationsWithoutLimit := []string{
		"ServiceKafkaQuotaDescribe",
		"ServiceKafkaQuotaDelete",
	}

	return !q.Has("limit") &&
		!slices.Contains(operationsWithoutLimit, operationID)
}

package aiven

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

func NewClient(opts ...Option) (Client, error) {
	d := new(Doer)
	err := envconfig.Process("", d)
	if err != nil {
		return nil, err
	}

	c := retryablehttp.NewClient()
	c.Logger = nil
	c.CheckRetry = checkRetry
	d.Client = c.StandardClient()

	// User settings
	for _, opt := range opts {
		opt(d)
	}

	if d.Debug {
		out := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		d.Logger = zerolog.New(out).With().Timestamp().Logger()
	}

	if d.Token == "" {
		return nil, fmt.Errorf("token is required. See https://api.aiven.io/doc/#section/Get-started/Authentication")
	}

	return newClient(d), nil
}

type Option func(*Doer)

func DebugOpt() Option {
	return func(d *Doer) {
		d.Debug = true
	}
}

type Doer struct {
	Host      string `envconfig:"AIVEN_WEB_URL" default:"https://api.aiven.io/v1"`
	UserAgent string `envconfig:"AIVEN_USER_AGENT" default:"aiven-go-client/v3"`
	Token     string `envconfig:"AIVEN_TOKEN"`
	Debug     bool   `envconfig:"AIVEN_DEBUG"`
	Client    *http.Client
	Logger    zerolog.Logger
}

type OperationIDKey struct{}

func (d *Doer) Do(ctx context.Context, operationID, method, path string, v any) ([]byte, error) {
	ctx = context.WithValue(ctx, OperationIDKey{}, operationID)

	var rsp *http.Response
	var err error
	if d.Debug {
		start := time.Now()
		defer func() {
			end := time.Since(start)
			var event *zerolog.Event
			if err != nil {
				event = d.Logger.Error().Err(err)
			} else {
				event = d.Logger.Info().Str("status", rsp.Status)
			}
			event.Ctx(ctx).
				Stringer("took", end).
				Str("operationID", operationID).
				Str("method", method).
				Str("path", path).
				Send()
		}()
	}

	rsp, err = d.do(ctx, method, path, v)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = multierror.Append(rsp.Body.Close())
	}()

	b, err := io.ReadAll(rsp.Body)
	if err != nil || rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		return nil, Error{Message: string(b), Status: rsp.StatusCode}
	}

	return b, err
}

func (d *Doer) do(ctx context.Context, method, path string, v any) (*http.Response, error) {
	var body io.Reader
	if !(v == nil || isEmpty(v)) {
		b, err := json.Marshal(v)
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

	// TODO: BAD hack to get around pagination in most cases
	// we should implement this properly at some point but for now
	// that should be its own issue
	query := req.URL.Query()
	query.Add("limit", "999")
	req.URL.RawQuery = query.Encode()
	return d.Client.Do(req)
}

func isEmpty(a any) bool {
	v := reflect.ValueOf(a)
	return !v.IsZero() || v.Kind() == reflect.Ptr && v.IsNil()
}

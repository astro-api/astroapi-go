// Package categories provides base HTTP functionality shared by all API category clients.
package categories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	astroerrors "github.com/astro-api/astroapi-go/errors"
	"github.com/astro-api/astroapi-go/internal/requestconfig"
	"github.com/astro-api/astroapi-go/internal/transport"
	"github.com/astro-api/astroapi-go/internal/validator"
	"github.com/astro-api/astroapi-go/option"
)

// BaseCategoryClient is the shared HTTP foundation for all category clients.
type BaseCategoryClient struct {
	Config *requestconfig.RequestConfig
}

// BuildURL joins the Config.BaseURL with the provided path segments.
// Leading and trailing slashes are trimmed from each segment.
func (b *BaseCategoryClient) BuildURL(segments ...string) string {
	base := strings.TrimRight(b.Config.BaseURL, "/")
	var parts []string
	for _, s := range segments {
		s = strings.Trim(s, "/")
		if s != "" {
			parts = append(parts, s)
		}
	}
	if len(parts) == 0 {
		return base
	}
	return base + "/" + strings.Join(parts, "/")
}

// Get executes a GET request. Params (if not nil) are encoded as query parameters.
func (b *BaseCategoryClient) Get(ctx context.Context, rawURL string, params any, out any, opts ...option.RequestOption) error {
	return b.MakeRequest(ctx, http.MethodGet, rawURL, params, out, opts...)
}

// Post executes a POST request with params marshalled as JSON body.
func (b *BaseCategoryClient) Post(ctx context.Context, rawURL string, params any, out any, opts ...option.RequestOption) error {
	return b.MakeRequest(ctx, http.MethodPost, rawURL, params, out, opts...)
}

// Put executes a PUT request with params marshalled as JSON body.
func (b *BaseCategoryClient) Put(ctx context.Context, rawURL string, params any, out any, opts ...option.RequestOption) error {
	return b.MakeRequest(ctx, http.MethodPut, rawURL, params, out, opts...)
}

// Delete executes a DELETE request with params marshalled as JSON body.
func (b *BaseCategoryClient) Delete(ctx context.Context, rawURL string, params any, out any, opts ...option.RequestOption) error {
	return b.MakeRequest(ctx, http.MethodDelete, rawURL, params, out, opts...)
}

// MakeRequest is the core request method used by all category clients.
func (b *BaseCategoryClient) MakeRequest(ctx context.Context, method, rawURL string, params any, out any, opts ...option.RequestOption) error {
	// Validate params.
	if err := validator.Validate(params); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Clone and apply per-request options.
	cfg := b.Config.Clone()
	rawOpts := make([]func(*requestconfig.RequestConfig), len(opts))
	for i, o := range opts {
		rawOpts[i] = o
	}
	cfg.Apply(rawOpts)

	// Build the HTTP request.
	var bodyReader io.Reader
	finalURL := rawURL

	if method == http.MethodGet {
		if params != nil {
			qp, err := encodeQueryParams(params)
			if err != nil {
				return fmt.Errorf("encoding query params: %w", err)
			}
			if len(qp) > 0 {
				finalURL = rawURL + "?" + qp.Encode()
			}
		}
	} else {
		if params != nil {
			data, err := json.Marshal(params)
			if err != nil {
				return fmt.Errorf("marshalling request body: %w", err)
			}
			bodyReader = bytes.NewReader(data)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, finalURL, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Apply extra headers from config.
	for key, values := range cfg.ExtraHeaders {
		for _, v := range values {
			req.Header.Set(key, v)
		}
	}

	// Build and execute with auth + retry transports.
	httpClient := buildHTTPClient(cfg)
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	// Store raw response if requested.
	if cfg.ResponseInto != nil {
		*cfg.ResponseInto = resp
	}

	// Handle non-2xx responses.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return astroerrors.NewFromResponse(req, resp)
	}

	// If out is nil, discard the body.
	if out == nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}

	// If out is *string, return raw body text.
	if sp, ok := out.(*string); ok {
		b2, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("reading response body: %w", err)
		}
		*sp = string(b2)
		return nil
	}

	return unmarshalResponse(resp.Body, out)
}

// unmarshalResponse reads the response body and unmarshals it into out.
// It tries the {"data":...} and {"result":...} envelope formats first,
// then falls back to direct unmarshalling.
func unmarshalResponse(body io.Reader, out any) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Try envelope unwrapping.
	var env map[string]json.RawMessage
	if jsonErr := json.Unmarshal(data, &env); jsonErr == nil {
		for _, key := range []string{"data", "result"} {
			if raw, ok := env[key]; ok {
				if err := json.Unmarshal(raw, out); err == nil {
					return nil
				}
			}
		}
	}

	// Fall back to direct unmarshal.
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("unmarshalling response: %w", err)
	}
	return nil
}

// buildHTTPClient creates an http.Client with AuthTransport wrapping RetryTransport.
func buildHTTPClient(cfg *requestconfig.RequestConfig) *http.Client {
	var base http.RoundTripper = http.DefaultTransport
	if cfg.HTTPClient != nil {
		base = cfg.HTTPClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}
	}

	retry := &transport.RetryTransport{
		Base:             base,
		MaxRetries:       cfg.MaxRetries,
		InitialDelay:     cfg.RetryDelay,
		RetryStatusCodes: cfg.RetryStatusCodes,
	}

	auth := &transport.AuthTransport{
		APIKey: cfg.APIKey,
		Base:   retry,
	}

	timeout := cfg.RequestTimeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	return &http.Client{
		Transport: auth,
		Timeout:   timeout,
	}
}

// encodeQueryParams converts a struct to url.Values using json tags or `url` tags.
// Uses JSON marshalling as an intermediate step for simplicity.
func encodeQueryParams(params any) (url.Values, error) {
	if params == nil {
		return nil, nil
	}

	// Marshal to JSON map, then to url.Values.
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	values := url.Values{}
	flattenQueryParams(values, m)
	return values, nil
}

func flattenQueryParams(values url.Values, m map[string]any) {
	for k, v := range m {
		if v == nil {
			continue
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < rv.Len(); i++ {
				values.Add(k, fmt.Sprintf("%v", rv.Index(i).Interface()))
			}
		default:
			values.Set(k, fmt.Sprintf("%v", v))
		}
	}
}

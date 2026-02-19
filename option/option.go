// Package option provides functional options for configuring API requests.
package option

import (
	"net/http"
	"time"

	"github.com/astro-api/astroapi-go/internal/requestconfig"
)

// RequestOption is a function that modifies a RequestConfig.
type RequestOption func(*requestconfig.RequestConfig)

// WithAPIKey sets the API key used for authentication.
func WithAPIKey(key string) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.APIKey = key
	}
}

// WithBaseURL sets the base URL for API requests.
func WithBaseURL(url string) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.BaseURL = url
	}
}

// WithHTTPClient sets a custom *http.Client.
func WithHTTPClient(c *http.Client) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.HTTPClient = c
	}
}

// WithMaxRetries sets the maximum number of retry attempts (default: 2).
func WithMaxRetries(n int) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.MaxRetries = n
	}
}

// WithRetryDelay sets the initial backoff delay between retries (default: 500ms).
func WithRetryDelay(d time.Duration) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.RetryDelay = d
	}
}

// WithRetryableStatusCodes overrides which HTTP status codes trigger a retry.
func WithRetryableStatusCodes(codes ...int) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.RetryStatusCodes = append([]int(nil), codes...)
	}
}

// WithRequestTimeout sets the per-request timeout.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.RequestTimeout = d
	}
}

// WithHeader adds an extra header to every request.
func WithHeader(key, value string) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		if rc.ExtraHeaders == nil {
			rc.ExtraHeaders = make(http.Header)
		}
		rc.ExtraHeaders.Set(key, value)
	}
}

// WithResponseInto stores the raw *http.Response into the given pointer after
// a successful request. Useful for inspecting response headers.
func WithResponseInto(resp **http.Response) RequestOption {
	return func(rc *requestconfig.RequestConfig) {
		rc.ResponseInto = resp
	}
}

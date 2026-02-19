// Package requestconfig provides the RequestConfig struct used throughout
// the SDK to configure HTTP requests.
package requestconfig

import (
	"net/http"
	"time"
)

const (
	DefaultBaseURL       = "https://api.astrology-api.io"
	DefaultMaxRetries    = 2
	DefaultRetryDelay    = 500 * time.Millisecond
	DefaultRequestTimeout = 30 * time.Second
)

// DefaultRetryStatusCodes are HTTP status codes that trigger a retry.
var DefaultRetryStatusCodes = []int{408, 429, 500, 502, 503, 504}

// RequestConfig holds all configuration for a single HTTP request.
type RequestConfig struct {
	APIKey             string
	BaseURL            string
	HTTPClient         *http.Client
	MaxRetries         int
	RetryDelay         time.Duration
	RetryStatusCodes   []int
	RequestTimeout     time.Duration
	ExtraHeaders       http.Header
	ResponseInto       **http.Response
}

// NewDefault returns a RequestConfig populated with default values.
func NewDefault() *RequestConfig {
	return &RequestConfig{
		BaseURL:          DefaultBaseURL,
		MaxRetries:       DefaultMaxRetries,
		RetryDelay:       DefaultRetryDelay,
		RequestTimeout:   DefaultRequestTimeout,
		RetryStatusCodes: append([]int(nil), DefaultRetryStatusCodes...),
		ExtraHeaders:     make(http.Header),
	}
}

// Clone returns a shallow copy of the config.
func (rc *RequestConfig) Clone() *RequestConfig {
	if rc == nil {
		return NewDefault()
	}
	clone := *rc
	// Deep-copy slices and maps to prevent mutation
	if rc.RetryStatusCodes != nil {
		clone.RetryStatusCodes = append([]int(nil), rc.RetryStatusCodes...)
	}
	if rc.ExtraHeaders != nil {
		clone.ExtraHeaders = rc.ExtraHeaders.Clone()
	}
	return &clone
}

// Apply applies a list of option functions to the config.
func (rc *RequestConfig) Apply(opts []func(*RequestConfig)) {
	for _, opt := range opts {
		if opt != nil {
			opt(rc)
		}
	}
}

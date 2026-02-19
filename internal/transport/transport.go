// Package transport provides HTTP RoundTripper implementations for authentication
// and retry logic.
package transport

import (
	"fmt"
	"net/http"
)

// AuthTransport injects Authorization and Content-Type headers into every request.
type AuthTransport struct {
	APIKey string
	Base   http.RoundTripper
}

// RoundTrip clones the request, injects auth headers, and delegates to the base transport.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	if clone.Header == nil {
		clone.Header = make(http.Header)
	}
	if t.APIKey != "" && clone.Header.Get("Authorization") == "" {
		clone.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.APIKey))
	}
	if clone.Header.Get("Content-Type") == "" && req.Body != nil {
		clone.Header.Set("Content-Type", "application/json")
	}
	if clone.Header.Get("Accept") == "" {
		clone.Header.Set("Accept", "application/json")
	}
	return t.base().RoundTrip(clone)
}

func (t *AuthTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

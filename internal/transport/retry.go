package transport

import (
	"bytes"
	"io"
	"math"
	"net/http"
	"time"
)

const maxBackoffCap = 30 * time.Second

// RetryTransport wraps a base RoundTripper with configurable retry logic.
// It buffers the request body so retries work correctly.
type RetryTransport struct {
	Base             http.RoundTripper
	MaxRetries       int
	InitialDelay     time.Duration
	RetryStatusCodes []int
}

// RoundTrip executes the request, retrying on network errors or configured status codes
// with exponential backoff capped at maxBackoffCap.
func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the body so it can be replayed on retry.
	var bodyBytes []byte
	if req.Body != nil && req.Body != http.NoBody {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		_ = req.Body.Close()
	}

	var (
		resp *http.Response
		err  error
	)

	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		// Check context before each attempt (except the first).
		if attempt > 0 {
			delay := t.backoff(attempt - 1)
			select {
			case <-req.Context().Done():
				return nil, req.Context().Err()
			case <-time.After(delay):
			}
		}

		// Rebuild request body for each attempt.
		cloned := req.Clone(req.Context())
		if bodyBytes != nil {
			cloned.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			cloned.ContentLength = int64(len(bodyBytes))
		}

		resp, err = t.base().RoundTrip(cloned)
		if err != nil {
			// Network error — retry.
			if attempt < t.MaxRetries {
				continue
			}
			return nil, err
		}

		if t.shouldRetryStatus(resp.StatusCode) && attempt < t.MaxRetries {
			// Drain and close the body to allow connection reuse.
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
			continue
		}

		return resp, nil
	}

	return resp, err
}

func (t *RetryTransport) shouldRetryStatus(code int) bool {
	for _, s := range t.RetryStatusCodes {
		if s == code {
			return true
		}
	}
	return false
}

func (t *RetryTransport) backoff(attempt int) time.Duration {
	delay := t.InitialDelay * time.Duration(math.Pow(2, float64(attempt)))
	if delay > maxBackoffCap {
		delay = maxBackoffCap
	}
	return delay
}

func (t *RetryTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

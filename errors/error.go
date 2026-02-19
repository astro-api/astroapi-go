// Package errors provides typed error types for the Astrology API SDK.
package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// apiErrorBody represents the error structure returned by the Astrology API.
// Shape: {"success": false, "error": {"error_code": "...", "message": "...", "field": "..."}}
type apiErrorBody struct {
	Success bool `json:"success"`
	Error   *struct {
		ErrorCode string `json:"error_code"`
		Message   string `json:"message"`
		Field     string `json:"field"`
	} `json:"error"`
	Message string `json:"message"`
}

// AstrologyError represents an error returned by the Astrology API.
type AstrologyError struct {
	StatusCode int
	Request    *http.Request
	Response   *http.Response
	// Body is the raw response body text.
	Body string
	// Message is the human-readable error message extracted from the body.
	Message string
	// Code is the machine-readable error code extracted from the body.
	Code string
}

// Error implements the error interface.
func (e *AstrologyError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("astrology API error %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("astrology API error %d", e.StatusCode)
}

// IsNotFound reports whether this is a 404 error.
func (e *AstrologyError) IsNotFound() bool { return e.StatusCode == http.StatusNotFound }

// IsRateLimit reports whether this is a 429 rate-limit error.
func (e *AstrologyError) IsRateLimit() bool { return e.StatusCode == http.StatusTooManyRequests }

// IsServerError reports whether this is a 5xx server error.
func (e *AstrologyError) IsServerError() bool { return e.StatusCode >= 500 }

// IsClientError reports whether this is a 4xx client error.
func (e *AstrologyError) IsClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// NewFromResponse constructs an AstrologyError from an http.Response whose
// status code is not 2xx. It reads the response body.
func NewFromResponse(req *http.Request, resp *http.Response) *AstrologyError {
	var bodyStr string
	if resp.Body != nil {
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		bodyStr = string(b)
	}

	ae := &AstrologyError{
		StatusCode: resp.StatusCode,
		Request:    req,
		Response:   resp,
		Body:       bodyStr,
	}

	// Try to parse JSON error body.
	var errBody apiErrorBody
	if err := json.Unmarshal([]byte(bodyStr), &errBody); err == nil {
		if errBody.Error != nil {
			ae.Message = errBody.Error.Message
			ae.Code = errBody.Error.ErrorCode
		} else if errBody.Message != "" {
			ae.Message = errBody.Message
		}
	}

	return ae
}

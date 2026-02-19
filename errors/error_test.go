package errors_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	astroerrors "github.com/astro-api/astroapi-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeResponse(statusCode int, body string) (*http.Request, *http.Response) {
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	resp := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
	return req, resp
}

func TestAstrologyError_ParseBody(t *testing.T) {
	body := `{"success":false,"error":{"error_code":"INVALID_KEY","message":"Invalid API key","field":"api_key"}}`
	req, resp := makeResponse(401, body)

	err := astroerrors.NewFromResponse(req, resp)
	require.NotNil(t, err)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "Invalid API key", err.Message)
	assert.Equal(t, "INVALID_KEY", err.Code)
}

func TestAstrologyError_Error(t *testing.T) {
	req, resp := makeResponse(404, `{"success":false,"error":{"error_code":"NOT_FOUND","message":"Not found"}}`)
	err := astroerrors.NewFromResponse(req, resp)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "Not found")
}

func TestAstrologyError_IsNotFound(t *testing.T) {
	req, resp := makeResponse(404, `{}`)
	err := astroerrors.NewFromResponse(req, resp)
	assert.True(t, err.IsNotFound())
	assert.False(t, err.IsRateLimit())
	assert.False(t, err.IsServerError())
}

func TestAstrologyError_IsRateLimit(t *testing.T) {
	req, resp := makeResponse(429, `{}`)
	err := astroerrors.NewFromResponse(req, resp)
	assert.True(t, err.IsRateLimit())
	assert.True(t, err.IsClientError())
}

func TestAstrologyError_IsServerError(t *testing.T) {
	req, resp := makeResponse(500, `{}`)
	err := astroerrors.NewFromResponse(req, resp)
	assert.True(t, err.IsServerError())
	assert.False(t, err.IsClientError())
}

func TestAstrologyError_FallbackMessage(t *testing.T) {
	body := `{"message":"Something went wrong"}`
	req, resp := makeResponse(400, body)
	err := astroerrors.NewFromResponse(req, resp)
	assert.Equal(t, "Something went wrong", err.Message)
}

func TestAstrologyError_InvalidJSON(t *testing.T) {
	req, resp := makeResponse(503, `not json`)
	err := astroerrors.NewFromResponse(req, resp)
	require.NotNil(t, err)
	assert.Equal(t, 503, err.StatusCode)
	assert.Equal(t, "not json", err.Body)
	assert.Empty(t, err.Message)
}

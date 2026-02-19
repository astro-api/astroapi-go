// Package testutil provides shared test helpers for the Astrology API SDK.
// Unit tests always run against a local mock server.
// Integration tests (named *_Integration) run against the real API when
// ASTROLOGY_API_KEY is set in the environment.
package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const envAPIKey = "ASTROLOGY_API_KEY"

// APIKey returns the API key from the environment, or "" if not set.
func APIKey() string {
	return os.Getenv(envAPIKey)
}

// IsIntegration reports whether ASTROLOGY_API_KEY is set.
func IsIntegration() bool {
	return APIKey() != ""
}

// MockHandler is a function that handles a mock HTTP request in unit tests.
type MockHandler func(w http.ResponseWriter, r *http.Request)

// NewClient always creates an AstrologyClient backed by a local mock httptest.Server.
// The mock handler is always called regardless of whether ASTROLOGY_API_KEY is set.
// Use this for unit tests. The caller must invoke the returned cleanup function.
func NewClient(t *testing.T, mock MockHandler) (client *astroapi.AstrologyClient, cleanup func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(mock))
	t.Cleanup(srv.Close)
	return astroapi.NewClient(
		option.WithAPIKey("test-key"),
		option.WithBaseURL(srv.URL),
		option.WithMaxRetries(0),
	), srv.Close
}

// NewIntegrationClient returns a client pointed at the real API.
// It requires ASTROLOGY_API_KEY to be set; otherwise it calls t.Skip.
func NewIntegrationClient(t *testing.T) *astroapi.AstrologyClient {
	t.Helper()
	SkipIfNoKey(t)
	return astroapi.NewClient(
		option.WithAPIKey(APIKey()),
		option.WithMaxRetries(1),
		option.WithRequestTimeout(30*time.Second),
	)
}

// JSON writes v as a JSON response with status 200.
func JSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// DataEnvelope wraps v in a {"data": v} envelope.
func DataEnvelope(v any) map[string]any { return map[string]any{"data": v} }

// ResultEnvelope wraps v in a {"result": v} envelope.
func ResultEnvelope(v any) map[string]any { return map[string]any{"result": v} }

// SkipIfNoKey skips a test that requires a real API key.
func SkipIfNoKey(t *testing.T) {
	t.Helper()
	if !IsIntegration() {
		t.Skip("skipping integration test: ASTROLOGY_API_KEY not set")
	}
}

// DefaultSubject returns a typical test subject (birth data for London 1990).
func DefaultSubject() shared.Subject {
	return shared.Subject{
		Name: "Test User",
		BirthData: shared.BirthData{
			Year:        1990,
			Month:       5,
			Day:         11,
			Hour:        14,
			Minute:      30,
			City:        "London",
			CountryCode: "GB",
		},
	}
}

// DefaultSubject2 returns a second typical test subject.
func DefaultSubject2() shared.Subject {
	return shared.Subject{
		Name: "Test User 2",
		BirthData: shared.BirthData{
			Year:        1992,
			Month:       3,
			Day:         27,
			Hour:        9,
			Minute:      0,
			City:        "Paris",
			CountryCode: "FR",
		},
	}
}

// DefaultDateTimeLocation returns a typical datetime/location.
func DefaultDateTimeLocation() shared.DateTimeLocation {
	return shared.DateTimeLocation{
		Year:        2024,
		Month:       1,
		Day:         15,
		Hour:        12,
		Minute:      0,
		City:        "London",
		CountryCode: "GB",
	}
}

// AssertNoError is a helper that calls t.Fatalf on error.
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

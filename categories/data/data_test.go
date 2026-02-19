package data_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/data"
	astroerrors "github.com/astro-api/astroapi-go/errors"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

// ---- GetNow ----------------------------------------------------------------

func TestDataClient_GetNow(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/now", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"current_time": "2024-01-15T12:00:00Z"}))
	})
	defer cleanup()

	result, err := client.Data.GetNow(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetPositions ----------------------------------------------------------

func TestDataClient_GetPositions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/positions", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sun": map[string]any{"degree": 45.0}}))
	})
	defer cleanup()

	result, err := client.Data.GetPositions(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDataClient_GetPositions_ValidationError(t *testing.T) {
	// No server needed — validation fails before any HTTP call.
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Data.GetPositions(ctx, data.PositionsParams{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validation")
}

func TestDataClient_GetPositions_APIError(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("skipping mock-only error test in integration mode")
	}
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		testutil.JSON(w, map[string]any{
			"success": false,
			"error": map[string]any{
				"error_code": "UNAUTHORIZED",
				"message":    "Invalid API key",
			},
		})
	})
	defer cleanup()

	_, err := client.Data.GetPositions(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.Error(t, err)
	var apiErr *astroerrors.AstrologyError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, 401, apiErr.StatusCode)
	assert.Equal(t, "Invalid API key", apiErr.Message)
	assert.Equal(t, "UNAUTHORIZED", apiErr.Code)
}

// ---- GetEnhancedPositions -------------------------------------------------

func TestDataClient_GetEnhancedPositions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/positions/enhanced", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"enhanced": true}))
	})
	defer cleanup()

	result, err := client.Data.GetEnhancedPositions(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetGlobalPositions ---------------------------------------------------

func TestDataClient_GetGlobalPositions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/global-positions", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"planets": []any{}}))
	})
	defer cleanup()

	result, err := client.Data.GetGlobalPositions(ctx, data.GlobalPositionsParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetAspects -----------------------------------------------------------

func TestDataClient_GetAspects(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/aspects", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"aspects": []any{}}))
	})
	defer cleanup()

	result, err := client.Data.GetAspects(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetEnhancedAspects ---------------------------------------------------

func TestDataClient_GetEnhancedAspects(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/aspects/enhanced", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"aspects": []any{}}))
	})
	defer cleanup()

	result, err := client.Data.GetEnhancedAspects(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetHouseCusps --------------------------------------------------------

func TestDataClient_GetHouseCusps(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/house-cusps", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"houses": []any{}}))
	})
	defer cleanup()

	result, err := client.Data.GetHouseCusps(ctx, data.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetLunarMetrics ------------------------------------------------------

func TestDataClient_GetLunarMetrics(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/lunar-metrics", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"phase": "full"}))
	})
	defer cleanup()

	result, err := client.Data.GetLunarMetrics(ctx, data.LunarMetricsParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- GetEnhancedLunarMetrics ----------------------------------------------

func TestDataClient_GetEnhancedLunarMetrics(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/data/lunar-metrics/enhanced", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"enhanced": true}))
	})
	defer cleanup()

	result, err := client.Data.GetEnhancedLunarMetrics(ctx, data.LunarMetricsParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

// ---- Envelope unwrapping --------------------------------------------------

func TestDataClient_DataEnvelope(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("envelope test is mock-only")
	}
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"key": "value"}))
	})
	defer cleanup()

	result, err := client.Data.GetNow(ctx)
	require.NoError(t, err)
	m := map[string]any(*result)
	assert.Equal(t, "value", m["key"])
}

func TestDataClient_ResultEnvelope(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("envelope test is mock-only")
	}
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.JSON(w, testutil.ResultEnvelope(map[string]any{"key": "result_value"}))
	})
	defer cleanup()

	result, err := client.Data.GetNow(ctx)
	require.NoError(t, err)
	m := map[string]any(*result)
	assert.Equal(t, "result_value", m["key"])
}

// ---- Retry ----------------------------------------------------------------

func TestDataClient_Retry503(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("retry test is mock-only")
	}
	callCount := 0
	var srvURL string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if callCount < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"retried": true}))
	}))
	defer srv.Close()
	srvURL = srv.URL

	client := astroapi.NewClient(
		option.WithAPIKey("test-key"),
		option.WithBaseURL(srvURL),
		option.WithMaxRetries(3),
	)

	result, err := client.Data.GetNow(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, callCount, "expected 2 503s then 1 success = 3 calls total")
}

// ---- Request body verification --------------------------------------------

func TestDataClient_GetPositions_RequestBody(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("body inspection is mock-only")
	}
	var gotBody map[string]any
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	_ = client
	subject := testutil.DefaultSubject()
	client2, cleanup2 := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup2()

	_, err := client2.Data.GetPositions(ctx, data.PositionsParams{Subject: subject})
	require.NoError(t, err)
	require.NotNil(t, gotBody["subject"])
}

// ---- ResponseInto ---------------------------------------------------------

func TestDataClient_ResponseInto(t *testing.T) {
	if testutil.IsIntegration() {
		t.Skip("ResponseInto test is mock-only")
	}
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "test-req-123")
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	var rawResp *http.Response
	_, err := client.Data.GetNow(ctx, option.WithResponseInto(&rawResp))
	require.NoError(t, err)
	require.NotNil(t, rawResp)
	assert.Equal(t, "test-req-123", rawResp.Header.Get("X-Request-Id"))
}

// ---- Integration-only: smoke test all data endpoints ---------------------

func TestDataClient_Integration_AllEndpoints(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetNow", func(t *testing.T) {
		result, err := client.Data.GetNow(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetPositions", func(t *testing.T) {
		result, err := client.Data.GetPositions(ctx, data.PositionsParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetAspects", func(t *testing.T) {
		result, err := client.Data.GetAspects(ctx, data.PositionsParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetHouseCusps", func(t *testing.T) {
		result, err := client.Data.GetHouseCusps(ctx, data.PositionsParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetLunarMetrics", func(t *testing.T) {
		result, err := client.Data.GetLunarMetrics(ctx, data.LunarMetricsParams{
			DatetimeLocation: testutil.DefaultDateTimeLocation(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetGlobalPositions", func(t *testing.T) {
		result, err := client.Data.GetGlobalPositions(ctx, data.GlobalPositionsParams{
			DatetimeLocation: testutil.DefaultDateTimeLocation(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

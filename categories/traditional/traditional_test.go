package traditional_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/traditional"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestTraditionalClient_GetDignitiesReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/dignities", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"dignities": []any{}}))
	})
	defer cleanup()

	result, err := client.Traditional.GetDignitiesReport(ctx, traditional.AnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetDignitiesReport_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Traditional.GetDignitiesReport(ctx, traditional.AnalysisParams{})
	require.Error(t, err)
}

func TestTraditionalClient_GetProfections(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/profections", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"house": 5}))
	})
	defer cleanup()

	result, err := client.Traditional.GetProfections(ctx, traditional.ProfectionParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetCapabilities(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/capabilities", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"capabilities": []any{}}))
	})
	defer cleanup()

	result, err := client.Traditional.GetCapabilities(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/analysis", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Traditional.GetAnalysis(ctx, traditional.AnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetLotsAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/lots", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"lots": []any{}}))
	})
	defer cleanup()

	result, err := client.Traditional.GetLotsAnalysis(ctx, traditional.AnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetAnnualProfection(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/profections/annual", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"year": 34}))
	})
	defer cleanup()

	result, err := client.Traditional.GetAnnualProfection(ctx, traditional.ProfectionParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetProfectionTimeline(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/profections/timeline", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Traditional.GetProfectionTimeline(ctx, traditional.ProfectionTimelineParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_GetHorary(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/traditional/horary", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Traditional.GetHorary(ctx, traditional.AnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTraditionalClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetDignitiesReport", func(t *testing.T) {
		result, err := client.Traditional.GetDignitiesReport(ctx, traditional.AnalysisParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetProfections", func(t *testing.T) {
		result, err := client.Traditional.GetProfections(ctx, traditional.ProfectionParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

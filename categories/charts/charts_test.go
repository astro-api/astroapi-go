package charts_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/charts"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestChartsClient_GetNatal(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/natal", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "natal"}))
	})
	defer cleanup()

	result, err := client.Charts.GetNatal(ctx, charts.NatalChartParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetNatal_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Charts.GetNatal(ctx, charts.NatalChartParams{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validation")
}

func TestChartsClient_GetSynastry(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/synastry", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "synastry"}))
	})
	defer cleanup()

	result, err := client.Charts.GetSynastry(ctx, charts.SynastryChartParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetComposite(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/composite", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "composite"}))
	})
	defer cleanup()

	result, err := client.Charts.GetComposite(ctx, charts.CompositeChartParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetTransit(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/transit", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "transit"}))
	})
	defer cleanup()

	result, err := client.Charts.GetTransit(ctx, charts.TransitChartParams{
		NatalSubject:    testutil.DefaultSubject(),
		TransitDatetime: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetSolarReturn(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/solar-return", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "solar_return"}))
	})
	defer cleanup()

	result, err := client.Charts.GetSolarReturn(ctx, charts.SolarReturnParams{
		Subject:    testutil.DefaultSubject(),
		ReturnYear: 2024,
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetLunarReturn(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/lunar-return", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "lunar_return"}))
	})
	defer cleanup()

	result, err := client.Charts.GetLunarReturn(ctx, charts.LunarReturnParams{
		Subject:    testutil.DefaultSubject(),
		ReturnDate: "2024-05-11",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetNatalTransits(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/natal-transits", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"events": []any{}}))
	})
	defer cleanup()

	result, err := client.Charts.GetNatalTransits(ctx, charts.NatalTransitsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetProgressions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/progressions", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "progression"}))
	})
	defer cleanup()

	result, err := client.Charts.GetProgressions(ctx, charts.ProgressionParams{
		Subject:         testutil.DefaultSubject(),
		TargetDate:      "2024-05-11",
		ProgressionType: "secondary",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetDirections(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/directions", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "direction"}))
	})
	defer cleanup()

	result, err := client.Charts.GetDirections(ctx, charts.DirectionParams{
		Subject:       testutil.DefaultSubject(),
		TargetDate:    "2024-05-11",
		DirectionType: "solar_arc",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetSolarReturnTransits(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/solar-return-transits", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"transits": []any{}}))
	})
	defer cleanup()

	result, err := client.Charts.GetSolarReturnTransits(ctx, charts.SolarReturnTransitsParams{
		Subject:    testutil.DefaultSubject(),
		ReturnYear: 2024,
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_GetLunarReturnTransits(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/charts/lunar-return-transits", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"transits": []any{}}))
	})
	defer cleanup()

	result, err := client.Charts.GetLunarReturnTransits(ctx, charts.LunarReturnTransitsParams{
		Subject:    testutil.DefaultSubject(),
		ReturnDate: "2024-05-11",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChartsClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetNatal", func(t *testing.T) {
		result, err := client.Charts.GetNatal(ctx, charts.NatalChartParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetSynastry", func(t *testing.T) {
		result, err := client.Charts.GetSynastry(ctx, charts.SynastryChartParams{
			Subject1: testutil.DefaultSubject(),
			Subject2: testutil.DefaultSubject2(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetTransit", func(t *testing.T) {
		result, err := client.Charts.GetTransit(ctx, charts.TransitChartParams{
			NatalSubject:    testutil.DefaultSubject(),
			TransitDatetime: testutil.DefaultDateTimeLocation(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

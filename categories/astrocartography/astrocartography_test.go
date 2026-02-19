package astrocartography_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/astrocartography"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestAstrocartographyClient_GetLines(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/lines", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"lines": []any{}}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GetLines(ctx, astrocartography.LinesParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GetLines_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Astrocartography.GetLines(ctx, astrocartography.LinesParams{})
	require.Error(t, err)
}

func TestAstrocartographyClient_AnalyzeLocation(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/location-analysis", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"score": 75}))
	})
	defer cleanup()

	result, err := client.Astrocartography.AnalyzeLocation(ctx, astrocartography.LocationAnalysisParams{
		Subject:  testutil.DefaultSubject(),
		Location: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GetLineMeanings(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/line-meanings", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"lines": []any{}}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GetLineMeanings(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_CompareLocations(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/compare-locations", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Astrocartography.CompareLocations(ctx, astrocartography.CompareLocationsParams{
		Subject: testutil.DefaultSubject(),
		Locations: []shared.DateTimeLocation{
			testutil.DefaultDateTimeLocation(),
		},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GenerateMap(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/map", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"map_url": "https://example.com/map.svg"}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GenerateMap(ctx, astrocartography.MapParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GenerateParanMap(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/paran-map", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"map_url": "https://example.com/paran.svg"}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GenerateParanMap(ctx, astrocartography.MapParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_FindPowerZones(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/power-zones", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"zones": []any{}}))
	})
	defer cleanup()

	result, err := client.Astrocartography.FindPowerZones(ctx, astrocartography.PowerZonesParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_SearchLocations(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/search-locations", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Astrocartography.SearchLocations(ctx, astrocartography.SearchLocationsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GenerateRelocationChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/relocation-chart", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart_type": "relocation"}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GenerateRelocationChart(ctx, astrocartography.RelocationChartParams{
		Subject:  testutil.DefaultSubject(),
		Location: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_GetReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/astrocartography/report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Astrocartography.GetReport(ctx, astrocartography.LinesParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAstrocartographyClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetLines", func(t *testing.T) {
		result, err := client.Astrocartography.GetLines(ctx, astrocartography.LinesParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetLineMeanings", func(t *testing.T) {
		result, err := client.Astrocartography.GetLineMeanings(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("AnalyzeLocation_Tokyo", func(t *testing.T) {
		result, err := client.Astrocartography.AnalyzeLocation(ctx, astrocartography.LocationAnalysisParams{
			Subject: testutil.DefaultSubject(),
			Location: shared.DateTimeLocation{
				City:        "Tokyo",
				CountryCode: "JP",
				Year:        2024, Month: 1, Day: 15, Hour: 12,
			},
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

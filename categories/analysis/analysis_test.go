package analysis_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/analysis"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestAnalysisClient_GetNatalReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/natal-report", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetNatalReport(ctx, analysis.NatalReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetNatalReport_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Analysis.GetNatalReport(ctx, analysis.NatalReportParams{})
	require.Error(t, err)
}

func TestAnalysisClient_GetSynastryReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/synastry-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"report_title": "Synastry Report"}))
	})
	defer cleanup()

	result, err := client.Analysis.GetSynastryReport(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"overall_score": 85}))
	})
	defer cleanup()

	result, err := client.Analysis.GetCompatibility(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetCompatibilityScore(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/compatibility-score", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"score": 78}))
	})
	defer cleanup()

	result, err := client.Analysis.GetCompatibilityScore(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetTransitReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/transit-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"transits": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetTransitReport(ctx, analysis.TransitReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetCareerAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/career", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"area": "career"}))
	})
	defer cleanup()

	result, err := client.Analysis.GetCareerAnalysis(ctx, analysis.NatalReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetLunarAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/lunar-analysis", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"moon_sign": "Cancer"}))
	})
	defer cleanup()

	result, err := client.Analysis.GetLunarAnalysis(ctx, analysis.LunarAnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetCompositeReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/composite-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetCompositeReport(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetRelationship(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/relationship", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetRelationship(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetRelationshipScore(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/relationship-score", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"score": 80}))
	})
	defer cleanup()

	result, err := client.Analysis.GetRelationshipScore(ctx, analysis.SynastryReportParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetNatalTransitReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/natal-transit-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"transits": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetNatalTransitReport(ctx, analysis.TransitReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetProgressionReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/progression-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetProgressionReport(ctx, analysis.ProgressionReportParams{
		Subject:         testutil.DefaultSubject(),
		TargetDate:      "2024-05-11",
		ProgressionType: "secondary",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetDirectionReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/direction-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetDirectionReport(ctx, analysis.DirectionReportParams{
		Subject:       testutil.DefaultSubject(),
		TargetDate:    "2024-05-11",
		DirectionType: "solar_arc",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetLunarReturnReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/lunar-return-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetLunarReturnReport(ctx, analysis.LunarReturnReportParams{
		Subject:    testutil.DefaultSubject(),
		ReturnDate: "2024-05-11",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetSolarReturnReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/solar-return-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Analysis.GetSolarReturnReport(ctx, analysis.SolarReturnReportParams{
		Subject:    testutil.DefaultSubject(),
		ReturnYear: 2024,
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetHealthAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/health", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"area": "health"}))
	})
	defer cleanup()

	result, err := client.Analysis.GetHealthAnalysis(ctx, analysis.NatalReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_GetKarmicAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/analysis/karmic", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"area": "karmic"}))
	})
	defer cleanup()

	result, err := client.Analysis.GetKarmicAnalysis(ctx, analysis.NatalReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAnalysisClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetNatalReport", func(t *testing.T) {
		result, err := client.Analysis.GetNatalReport(ctx, analysis.NatalReportParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCompatibilityScore", func(t *testing.T) {
		result, err := client.Analysis.GetCompatibilityScore(ctx, analysis.SynastryReportParams{
			Subject1: testutil.DefaultSubject(),
			Subject2: testutil.DefaultSubject2(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCareerAnalysis", func(t *testing.T) {
		result, err := client.Analysis.GetCareerAnalysis(ctx, analysis.NatalReportParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

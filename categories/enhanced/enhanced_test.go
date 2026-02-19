package enhanced_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/enhanced"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestEnhancedClient_GetPersonalAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/enhanced/personal", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"analysis": "enhanced"}))
	})
	defer cleanup()

	result, err := client.Enhanced.GetPersonalAnalysis(ctx, enhanced.PersonalAnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEnhancedClient_GetPersonalAnalysis_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Enhanced.GetPersonalAnalysis(ctx, enhanced.PersonalAnalysisParams{})
	require.Error(t, err)
}

func TestEnhancedClient_GetGlobalAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/enhanced/global", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"global": true}))
	})
	defer cleanup()

	result, err := client.Enhanced.GetGlobalAnalysis(ctx, enhanced.GlobalAnalysisParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEnhancedClient_GetGlobalAnalysisChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/enhanced_charts/global", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart": "global"}))
	})
	defer cleanup()

	result, err := client.Enhanced.GetGlobalAnalysisChart(ctx, enhanced.GlobalAnalysisParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEnhancedClient_GetPersonalAnalysisChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/enhanced_charts/personal", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"chart": "personal"}))
	})
	defer cleanup()

	result, err := client.Enhanced.GetPersonalAnalysisChart(ctx, enhanced.PersonalAnalysisParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEnhancedClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetPersonalAnalysis", func(t *testing.T) {
		result, err := client.Enhanced.GetPersonalAnalysis(ctx, enhanced.PersonalAnalysisParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetGlobalAnalysis", func(t *testing.T) {
		result, err := client.Enhanced.GetGlobalAnalysis(ctx, enhanced.GlobalAnalysisParams{
			DatetimeLocation: testutil.DefaultDateTimeLocation(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

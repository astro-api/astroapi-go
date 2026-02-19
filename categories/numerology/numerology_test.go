package numerology_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/numerology"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestNumerologyClient_GetReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/numerology/core-numbers", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"life_path": 7}))
	})
	defer cleanup()

	result, err := client.Numerology.GetReport(ctx, numerology.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNumerologyClient_GetReport_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Numerology.GetReport(ctx, numerology.SingleSubjectParams{})
	require.Error(t, err)
}

func TestNumerologyClient_GetComprehensiveReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/numerology/comprehensive", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Numerology.GetComprehensiveReport(ctx, numerology.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNumerologyClient_GetCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/numerology/compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"compatibility_score": 85}))
	})
	defer cleanup()

	result, err := client.Numerology.GetCompatibility(ctx, numerology.CompatibilityParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNumerologyClient_GetCoreNumbers(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/numerology/core-numbers", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"life_path": 7}))
	})
	defer cleanup()

	result, err := client.Numerology.GetCoreNumbers(ctx, numerology.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNumerologyClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetCoreNumbers", func(t *testing.T) {
		result, err := client.Numerology.GetCoreNumbers(ctx, numerology.SingleSubjectParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCompatibility", func(t *testing.T) {
		result, err := client.Numerology.GetCompatibility(ctx, numerology.CompatibilityParams{
			Subject1: testutil.DefaultSubject(),
			Subject2: testutil.DefaultSubject2(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

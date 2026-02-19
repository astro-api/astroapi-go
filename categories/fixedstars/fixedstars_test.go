package fixedstars_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/fixedstars"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestFixedStarsClient_GetList(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/fixed-stars/list", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"name": "Regulus"}}))
	})
	defer cleanup()

	result, err := client.FixedStars.GetList(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFixedStarsClient_GetPositions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/fixed-stars/positions", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.FixedStars.GetPositions(ctx, fixedstars.PositionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFixedStarsClient_GetPositions_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.FixedStars.GetPositions(ctx, fixedstars.PositionsParams{})
	require.Error(t, err)
}

func TestFixedStarsClient_GetConjunctions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/fixed-stars/conjunctions", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.FixedStars.GetConjunctions(ctx, fixedstars.ConjunctionsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFixedStarsClient_GetPresets(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/fixed-stars/presets", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{"essential", "traditional"}))
	})
	defer cleanup()

	result, err := client.FixedStars.GetPresets(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFixedStarsClient_GenerateReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/fixed-stars/report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.FixedStars.GenerateReport(ctx, fixedstars.ReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFixedStarsClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetList", func(t *testing.T) {
		result, err := client.FixedStars.GetList(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetPositions", func(t *testing.T) {
		result, err := client.FixedStars.GetPositions(ctx, fixedstars.PositionsParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetConjunctions", func(t *testing.T) {
		result, err := client.FixedStars.GetConjunctions(ctx, fixedstars.ConjunctionsParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

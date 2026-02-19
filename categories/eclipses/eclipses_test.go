package eclipses_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/eclipses"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestEclipsesClient_GetUpcoming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/eclipses/upcoming", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"date": "2024-04-08", "type": "solar"}}))
	})
	defer cleanup()

	result, err := client.Eclipses.GetUpcoming(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEclipsesClient_GetUpcoming_WithParams(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "5", r.URL.Query().Get("limit"))
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Eclipses.GetUpcoming(ctx, &eclipses.UpcomingParams{Limit: 5})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEclipsesClient_CheckNatalImpact(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/eclipses/natal-check", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"impact": "high"}))
	})
	defer cleanup()

	result, err := client.Eclipses.CheckNatalImpact(ctx, eclipses.NatalCheckParams{
		Subject:     testutil.DefaultSubject(),
		EclipseDate: "2024-04-08",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEclipsesClient_CheckNatalImpact_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Eclipses.CheckNatalImpact(ctx, eclipses.NatalCheckParams{})
	require.Error(t, err)
}

func TestEclipsesClient_GetInterpretation(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/eclipses/interpretation", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"interpretation": "Solar eclipse in Aries..."}))
	})
	defer cleanup()

	result, err := client.Eclipses.GetInterpretation(ctx, eclipses.InterpretationParams{
		EclipseDate: "2024-04-08",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEclipsesClient_GetList(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/eclipses/upcoming", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Eclipses.GetList(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestEclipsesClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetUpcoming", func(t *testing.T) {
		result, err := client.Eclipses.GetUpcoming(ctx, &eclipses.UpcomingParams{Limit: 3})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("CheckNatalImpact", func(t *testing.T) {
		result, err := client.Eclipses.CheckNatalImpact(ctx, eclipses.NatalCheckParams{
			Subject:     testutil.DefaultSubject(),
			EclipseDate: "2024-04-08",
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

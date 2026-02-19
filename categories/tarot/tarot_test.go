package tarot_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/tarot"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestTarotClient_GetDraw(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/draw", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"name": "The Fool"}}))
	})
	defer cleanup()

	result, err := client.Tarot.GetDraw(ctx, tarot.DrawCardsParams{Count: 3})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GetDraw_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Tarot.GetDraw(ctx, tarot.DrawCardsParams{})
	require.Error(t, err)
}

func TestTarotClient_GetCard(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/cards/the-fool", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"name": "The Fool", "number": 0}))
	})
	defer cleanup()

	result, err := client.Tarot.GetCard(ctx, "the-fool")
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GetDailyCard(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/daily-card", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"card": "The Star"}))
	})
	defer cleanup()

	result, err := client.Tarot.GetDailyCard(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GetCardsGlossary(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/glossary/cards", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Tarot.GetCardsGlossary(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GetSpreadsGlossary(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/glossary/spreads", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Tarot.GetSpreadsGlossary(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GenerateReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Tarot.GenerateReport(ctx, tarot.TarotReportParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_GenerateSynastryReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/synastry-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sections": []any{}}))
	})
	defer cleanup()

	result, err := client.Tarot.GenerateSynastryReport(ctx, tarot.TarotSynastryParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_SearchCards(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/tarot/cards/search", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Tarot.SearchCards(ctx, &tarot.SearchParams{Query: "fool"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestTarotClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetDraw_3cards", func(t *testing.T) {
		result, err := client.Tarot.GetDraw(ctx, tarot.DrawCardsParams{Count: 3})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetDailyCard", func(t *testing.T) {
		result, err := client.Tarot.GetDailyCard(ctx, nil)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCardsGlossary", func(t *testing.T) {
		result, err := client.Tarot.GetCardsGlossary(ctx, nil)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

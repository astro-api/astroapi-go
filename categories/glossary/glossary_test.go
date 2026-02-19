package glossary_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/astro-api/astroapi-go/categories/glossary"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestGlossaryClient_GetCountries(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/countries", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"code": "GB", "name": "United Kingdom"}}))
	})
	defer cleanup()

	result, err := client.Glossary.GetCountries(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetElements(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/elements", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{"Fire", "Earth", "Air", "Water"}))
	})
	defer cleanup()

	result, err := client.Glossary.GetElements(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetHouseSystems(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/house-systems", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"code": "P", "name": "Placidus"}}))
	})
	defer cleanup()

	result, err := client.Glossary.GetHouseSystems(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetLanguages(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/languages", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{"en", "es", "fr"}))
	})
	defer cleanup()

	result, err := client.Glossary.GetLanguages(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetCities(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/cities", r.URL.Path)
		assert.Equal(t, "London", r.URL.Query().Get("search"))
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"name": "London"}}))
	})
	defer cleanup()

	result, err := client.Glossary.GetCities(ctx, &glossary.CitySearchParams{Search: "London"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetActivePoints(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/active-points", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"name": "Sun"}}))
	})
	defer cleanup()

	result, err := client.Glossary.GetActivePoints(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetThemes(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/themes", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{"light", "dark"}))
	})
	defer cleanup()

	result, err := client.Glossary.GetThemes(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetZodiacTypes(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/zodiac-types", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{"Tropic", "Sidereal"}))
	})
	defer cleanup()

	result, err := client.Glossary.GetZodiacTypes(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetPrimaryActivePoints(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/active-points/primary", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Glossary.GetPrimaryActivePoints(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetFixedStars(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/fixed-stars", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Glossary.GetFixedStars(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetKeywords(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/keywords", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Glossary.GetKeywords(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_GetLifeAreas(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/glossary/life-areas", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Glossary.GetLifeAreas(ctx, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGlossaryClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetCountries", func(t *testing.T) {
		result, err := client.Glossary.GetCountries(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetHouseSystems", func(t *testing.T) {
		result, err := client.Glossary.GetHouseSystems(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetLanguages", func(t *testing.T) {
		result, err := client.Glossary.GetLanguages(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCities_London", func(t *testing.T) {
		result, err := client.Glossary.GetCities(ctx, &glossary.CitySearchParams{Search: "London", CountryCode: "GB"})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

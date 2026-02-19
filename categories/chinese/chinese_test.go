package chinese_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/chinese"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestChineseClient_CalculateBaZi(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/bazi", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"year_pillar": "Geng Wu"}))
	})
	defer cleanup()

	result, err := client.Chinese.CalculateBaZi(ctx, chinese.BaZiParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_CalculateBaZi_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Chinese.CalculateBaZi(ctx, chinese.BaZiParams{})
	require.Error(t, err)
}

func TestChineseClient_CalculateCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"score": 80}))
	})
	defer cleanup()

	result, err := client.Chinese.CalculateCompatibility(ctx, chinese.CompatibilityParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_GetSolarTerms(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/calendar/solar-terms/2024", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Chinese.GetSolarTerms(ctx, 2024)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_GetZodiacAnimal(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/zodiac/dragon", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"animal": "dragon"}))
	})
	defer cleanup()

	result, err := client.Chinese.GetZodiacAnimal(ctx, "dragon")
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_CalculateLuckPillars(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/luck-pillars", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"pillars": []any{}}))
	})
	defer cleanup()

	result, err := client.Chinese.CalculateLuckPillars(ctx, chinese.LuckPillarsParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_CalculateMingGua(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/ming-gua", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"gua": 3}))
	})
	defer cleanup()

	result, err := client.Chinese.CalculateMingGua(ctx, chinese.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_GetYearlyForecast(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/chinese/yearly-forecast", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"year": 2024}))
	})
	defer cleanup()

	result, err := client.Chinese.GetYearlyForecast(ctx, chinese.YearlyForecastParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestChineseClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("CalculateBaZi", func(t *testing.T) {
		result, err := client.Chinese.CalculateBaZi(ctx, chinese.BaZiParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetZodiacAnimal_dragon", func(t *testing.T) {
		result, err := client.Chinese.GetZodiacAnimal(ctx, "dragon")
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetSolarTerms_2024", func(t *testing.T) {
		result, err := client.Chinese.GetSolarTerms(ctx, 2024)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

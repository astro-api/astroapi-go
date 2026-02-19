package horoscope_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/horoscope"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestHoroscopeClient_GetPersonalDaily(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/personal/daily", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"overall_rating": 8}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetPersonalDaily(ctx, horoscope.PersonalDailyParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetPersonalDaily_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Horoscope.GetPersonalDaily(ctx, horoscope.PersonalDailyParams{})
	require.Error(t, err)
}

func TestHoroscopeClient_GetSignDaily(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/daily", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"sign": "aries"}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignDaily(ctx, horoscope.SignHoroscopeParams{
		Sign: "Aries",
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignWeekly(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/weekly", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"week": "2024-W01"}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignWeekly(ctx, horoscope.SignWeeklyParams{Sign: "Taurus"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignMonthly(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/monthly", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"month": 1}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignMonthly(ctx, horoscope.SignMonthlyParams{Sign: "Gemini"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignYearly(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/yearly", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"year": 2024}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignYearly(ctx, horoscope.SignYearlyParams{Sign: "Cancer"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetChinese(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/chinese/bazi", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"animal": "horse"}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetChinese(ctx, horoscope.ChineseHoroscopeParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetPersonalDailyText(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/personal/daily/text", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"text": "Your day looks bright."}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetPersonalDailyText(ctx, horoscope.PersonalTextParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignDailyText(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/daily/text", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"text": "Aries daily text."}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignDailyText(ctx, horoscope.SignHoroscopeParams{Sign: "Aries"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignWeeklyText(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/weekly/text", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"text": "Taurus weekly text."}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignWeeklyText(ctx, horoscope.SignWeeklyParams{Sign: "Taurus"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_GetSignMonthlyText(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/horoscope/sign/monthly/text", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"text": "Gemini monthly text."}))
	})
	defer cleanup()

	result, err := client.Horoscope.GetSignMonthlyText(ctx, horoscope.SignMonthlyParams{Sign: "Gemini"})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestHoroscopeClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetPersonalDaily", func(t *testing.T) {
		result, err := client.Horoscope.GetPersonalDaily(ctx, horoscope.PersonalDailyParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetSignDaily_Aries", func(t *testing.T) {
		result, err := client.Horoscope.GetSignDaily(ctx, horoscope.SignHoroscopeParams{
			Sign: "Aries",
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetSignWeekly_Taurus", func(t *testing.T) {
		result, err := client.Horoscope.GetSignWeekly(ctx, horoscope.SignWeeklyParams{Sign: "Taurus"})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetSignMonthly_Gemini", func(t *testing.T) {
		result, err := client.Horoscope.GetSignMonthly(ctx, horoscope.SignMonthlyParams{Sign: "Gemini"})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

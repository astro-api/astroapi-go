package lunar_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/lunar"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestLunarClient_GetPhase(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/lunar/phases", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"phase": "full", "date": "2024-01-25"}}))
	})
	defer cleanup()

	result, err := client.Lunar.GetPhase(ctx, lunar.PhasesParams{
		DateRange: shared.DateRange{Start: "2024-01-01", End: "2024-01-31"},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLunarClient_GetPhase_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Lunar.GetPhase(ctx, lunar.PhasesParams{})
	require.Error(t, err)
}

func TestLunarClient_GetCalendar(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/lunar/calendar/2024", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"year": 2024, "months": []any{}}))
	})
	defer cleanup()

	result, err := client.Lunar.GetCalendar(ctx, 2024, nil)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLunarClient_GetVoidOfCourse(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/lunar/void-of-course", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{map[string]any{"start": "2024-01-02T10:00:00Z"}}))
	})
	defer cleanup()

	result, err := client.Lunar.GetVoidOfCourse(ctx, lunar.VoidOfCourseParams{
		DateRange: shared.DateRange{Start: "2024-01-01", End: "2024-01-07"},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLunarClient_GetEvents(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/lunar/events", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope([]any{}))
	})
	defer cleanup()

	result, err := client.Lunar.GetEvents(ctx, lunar.EventsParams{
		DateRange: shared.DateRange{Start: "2024-01-01", End: "2024-01-31"},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLunarClient_GetMansions(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/lunar/mansions", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"mansion": 1}))
	})
	defer cleanup()

	result, err := client.Lunar.GetMansions(ctx, lunar.MansionsParams{
		DatetimeLocation: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLunarClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetPhase", func(t *testing.T) {
		result, err := client.Lunar.GetPhase(ctx, lunar.PhasesParams{
			DateRange: shared.DateRange{Start: "2024-01-01", End: "2024-01-31"},
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetCalendar_2024", func(t *testing.T) {
		result, err := client.Lunar.GetCalendar(ctx, 2024, nil)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetVoidOfCourse", func(t *testing.T) {
		result, err := client.Lunar.GetVoidOfCourse(ctx, lunar.VoidOfCourseParams{
			DateRange: shared.DateRange{Start: "2024-01-01", End: "2024-01-14"},
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

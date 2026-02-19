package svg_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/svg"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

const fakeSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="800" height="800"><circle cx="400" cy="400" r="350"/></svg>`

func TestSVGClient_GetNatalChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/svg/natal", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(fakeSVG))
	})
	defer cleanup()

	result, err := client.SVG.GetNatalChart(ctx, svg.NatalChartSVGParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "<svg")
}

func TestSVGClient_GetNatalChart_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.SVG.GetNatalChart(ctx, svg.NatalChartSVGParams{})
	require.Error(t, err)
}

func TestSVGClient_GetSynastryChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/svg/synastry", r.URL.Path)
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(fakeSVG))
	})
	defer cleanup()

	result, err := client.SVG.GetSynastryChart(ctx, svg.SynastryChartSVGParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestSVGClient_GetTransitChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/svg/transit", r.URL.Path)
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(fakeSVG))
	})
	defer cleanup()

	result, err := client.SVG.GetTransitChart(ctx, svg.TransitChartSVGParams{
		NatalSubject:    testutil.DefaultSubject(),
		TransitDatetime: testutil.DefaultDateTimeLocation(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestSVGClient_GetChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/svg/natal", r.URL.Path)
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(`<svg><circle/></svg>`))
	})
	defer cleanup()

	result, err := client.SVG.GetChart(ctx, svg.NatalChartSVGParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestSVGClient_GetCompositeChart(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/svg/composite", r.URL.Path)
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(`<svg><circle/></svg>`))
	})
	defer cleanup()

	result, err := client.SVG.GetCompositeChart(ctx, svg.CompositeChartSVGParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestSVGClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("GetNatalChart", func(t *testing.T) {
		result, err := client.SVG.GetNatalChart(ctx, svg.NatalChartSVGParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, "<svg")
	})

	t.Run("GetSynastryChart", func(t *testing.T) {
		result, err := client.SVG.GetSynastryChart(ctx, svg.SynastryChartSVGParams{
			Subject1: testutil.DefaultSubject(),
			Subject2: testutil.DefaultSubject2(),
		})
		require.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

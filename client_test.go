package astroapi_test

import (
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient_Defaults(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"))
	require.NotNil(t, client)
	assert.NotNil(t, client.Data)
	assert.NotNil(t, client.Charts)
	assert.NotNil(t, client.Horoscope)
	assert.NotNil(t, client.Analysis)
	assert.NotNil(t, client.Glossary)
	assert.NotNil(t, client.Astrocartography)
	assert.NotNil(t, client.Chinese)
	assert.NotNil(t, client.Eclipses)
	assert.NotNil(t, client.Lunar)
	assert.NotNil(t, client.Numerology)
	assert.NotNil(t, client.Tarot)
	assert.NotNil(t, client.Traditional)
	assert.NotNil(t, client.FixedStars)
	assert.NotNil(t, client.Insights)
	assert.NotNil(t, client.SVG)
	assert.NotNil(t, client.Enhanced)
}

func TestNewClient_InsightsSubClients(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"))
	require.NotNil(t, client.Insights)
	assert.NotNil(t, client.Insights.Relationship)
	assert.NotNil(t, client.Insights.Pet)
	assert.NotNil(t, client.Insights.Wellness)
	assert.NotNil(t, client.Insights.Financial)
	assert.NotNil(t, client.Insights.Business)
}

func TestNewClient_NoKey(t *testing.T) {
	// Should not panic even without an API key.
	client := astroapi.NewClient()
	require.NotNil(t, client)
}

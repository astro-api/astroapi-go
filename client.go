package astroapi

import (
	"os"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/categories/analysis"
	"github.com/astro-api/astroapi-go/categories/astrocartography"
	"github.com/astro-api/astroapi-go/categories/charts"
	"github.com/astro-api/astroapi-go/categories/chinese"
	"github.com/astro-api/astroapi-go/categories/data"
	"github.com/astro-api/astroapi-go/categories/eclipses"
	"github.com/astro-api/astroapi-go/categories/enhanced"
	"github.com/astro-api/astroapi-go/categories/fixedstars"
	"github.com/astro-api/astroapi-go/categories/glossary"
	"github.com/astro-api/astroapi-go/categories/horoscope"
	"github.com/astro-api/astroapi-go/categories/insights"
	"github.com/astro-api/astroapi-go/categories/lunar"
	"github.com/astro-api/astroapi-go/categories/numerology"
	"github.com/astro-api/astroapi-go/categories/svg"
	"github.com/astro-api/astroapi-go/categories/tarot"
	"github.com/astro-api/astroapi-go/categories/traditional"
	"github.com/astro-api/astroapi-go/internal/requestconfig"
	"github.com/astro-api/astroapi-go/option"
)

// AstrologyClient is the root client for the Astrology API.
// Create one with NewClient() and use the sub-clients to call specific endpoints.
type AstrologyClient struct {
	Data             *data.Client
	Charts           *charts.Client
	Horoscope        *horoscope.Client
	Analysis         *analysis.Client
	Glossary         *glossary.Client
	Astrocartography *astrocartography.Client
	Chinese          *chinese.Client
	Eclipses         *eclipses.Client
	Lunar            *lunar.Client
	Numerology       *numerology.Client
	Tarot            *tarot.Client
	Traditional      *traditional.Client
	FixedStars       *fixedstars.Client
	Insights         *insights.Client
	SVG              *svg.Client
	Enhanced         *enhanced.Client

	cfg *requestconfig.RequestConfig
}

// NewClient creates a new AstrologyClient. Configuration defaults are applied
// first, then options in order, then environment variable fallbacks.
//
//	client := astroapi.NewClient(
//	    option.WithAPIKey("your-api-key"),
//	)
func NewClient(opts ...option.RequestOption) *AstrologyClient {
	cfg := requestconfig.NewDefault()
	rawOpts := make([]func(*requestconfig.RequestConfig), len(opts))
	for i, o := range opts {
		rawOpts[i] = o
	}
	cfg.Apply(rawOpts)

	// Environment variable fallbacks.
	if cfg.APIKey == "" {
		cfg.APIKey = os.Getenv("ASTROLOGY_API_KEY")
	}
	if cfg.BaseURL == "" || cfg.BaseURL == requestconfig.DefaultBaseURL {
		if envURL := os.Getenv("ASTROLOGY_API_BASE_URL"); envURL != "" {
			cfg.BaseURL = envURL
		}
	}

	base := &categories.BaseCategoryClient{Config: cfg}

	return &AstrologyClient{
		cfg:              cfg,
		Data:             data.NewClient(base),
		Charts:           charts.NewClient(base),
		Horoscope:        horoscope.NewClient(base),
		Analysis:         analysis.NewClient(base),
		Glossary:         glossary.NewClient(base),
		Astrocartography: astrocartography.NewClient(base),
		Chinese:          chinese.NewClient(base),
		Eclipses:         eclipses.NewClient(base),
		Lunar:            lunar.NewClient(base),
		Numerology:       numerology.NewClient(base),
		Tarot:            tarot.NewClient(base),
		Traditional:      traditional.NewClient(base),
		FixedStars:       fixedstars.NewClient(base),
		Insights:         insights.NewClient(base),
		SVG:              svg.NewClient(base),
		Enhanced:         enhanced.NewClient(base),
	}
}

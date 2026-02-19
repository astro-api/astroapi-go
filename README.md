# astroapi-go

Official Go client for the [Astrology API v3](https://api.astrology-api.io).

## Installation

```bash
go get github.com/astro-api/astroapi-go
```

## Quick Start

```go
import (
    astroapi "github.com/astro-api/astroapi-go"
    "github.com/astro-api/astroapi-go/option"
)

client := astroapi.NewClient(
    option.WithAPIKey("your-api-key"),
)
```

## Environment Variables

| Variable | Description |
|---|---|
| `ASTROLOGY_API_KEY` | Your API key (used if not provided via `WithAPIKey`) |
| `ASTROLOGY_API_BASE_URL` | Override the base URL (default: `https://api.astrology-api.io`) |

## Sub-Clients

| Client | Field | API Prefix |
|---|---|---|
| Data | `client.Data` | `/api/v3/data` |
| Charts | `client.Charts` | `/api/v3/charts` |
| Horoscope | `client.Horoscope` | `/api/v3/horoscope` |
| Analysis | `client.Analysis` | `/api/v3/analysis` |
| Glossary | `client.Glossary` | `/api/v3/glossary` |
| Astrocartography | `client.Astrocartography` | `/api/v3/astrocartography` |
| Chinese | `client.Chinese` | `/api/v3/chinese` |
| Eclipses | `client.Eclipses` | `/api/v3/eclipses` |
| Lunar | `client.Lunar` | `/api/v3/lunar` |
| Numerology | `client.Numerology` | `/api/v3/numerology` |
| Tarot | `client.Tarot` | `/api/v3/tarot` |
| Traditional | `client.Traditional` | `/api/v3/traditional` |
| Fixed Stars | `client.FixedStars` | `/api/v3/fixed-stars` |
| Insights | `client.Insights` | `/api/v3/insights` |
| SVG | `client.SVG` | `/api/v3/svg` |
| Enhanced | `client.Enhanced` | `/api/v3/enhanced` |

Insights also exposes nested sub-clients: `Insights.Relationship`, `Insights.Pet`, `Insights.Wellness`, `Insights.Financial`, `Insights.Business`.

## Examples

### Get Current Moment

```go
now, err := client.Data.GetNow(ctx)
```

### Planetary Positions

```go
import "github.com/astro-api/astroapi-go/categories/data"
import "github.com/astro-api/astroapi-go/shared"

positions, err := client.Data.GetPositions(ctx, data.PositionsParams{
    Subject: shared.Subject{
        Name: "Alice",
        BirthData: shared.BirthData{
            Year: 1990, Month: 5, Day: 11,
            Hour: 14, Minute: 30,
            City: "London", CountryCode: "GB",
        },
    },
})
```

### Natal Chart

```go
import "github.com/astro-api/astroapi-go/categories/charts"

chart, err := client.Charts.GetNatal(ctx, charts.NatalChartParams{
    Subject: shared.Subject{
        BirthData: shared.BirthData{
            Year: 1990, Month: 5, Day: 11,
            Hour: 14, Minute: 30,
            City: "London", CountryCode: "GB",
        },
    },
})
```

### Daily Horoscope

```go
import "github.com/astro-api/astroapi-go/categories/horoscope"

daily, err := client.Horoscope.GetSignDaily(ctx, horoscope.SignHoroscopeParams{
    Sign: "Aries",
})
```

### Compatibility

```go
import "github.com/astro-api/astroapi-go/categories/analysis"

compat, err := client.Analysis.GetCompatibility(ctx, analysis.SynastryReportParams{
    Subject1: shared.Subject{BirthData: shared.BirthData{Year: 1990, Month: 5, Day: 11, Hour: 14, Minute: 30, City: "London", CountryCode: "GB"}},
    Subject2: shared.Subject{BirthData: shared.BirthData{Year: 1992, Month: 3, Day: 27, Hour: 9, Minute: 0, City: "Paris", CountryCode: "FR"}},
})
```

### Relationship Insights

```go
import "github.com/astro-api/astroapi-go/categories/insights"

score, err := client.Insights.Relationship.GetCompatibilityScore(ctx, insights.TwoSubjectParams{
    Subject1: subject1,
    Subject2: subject2,
})
```

## Error Handling

```go
import astroerrors "github.com/astro-api/astroapi-go/errors"

positions, err := client.Data.GetPositions(ctx, params)
if err != nil {
    if apiErr, ok := err.(*astroerrors.AstrologyError); ok {
        fmt.Println(apiErr.StatusCode, apiErr.Message, apiErr.Code)
        if apiErr.IsRateLimit() {
            // handle 429
        }
        if apiErr.IsNotFound() {
            // handle 404
        }
    }
}
```

## Per-Request Options

Override configuration for a single call:

```go
import (
    "time"
    "github.com/astro-api/astroapi-go/option"
)

positions, err := client.Data.GetPositions(ctx, params,
    option.WithRequestTimeout(5*time.Second),
    option.WithHeader("X-Custom-Header", "value"),
)
```

## Configuration Options

| Option | Description | Default |
|---|---|---|
| `WithAPIKey(key)` | API key | `""` |
| `WithBaseURL(url)` | Base URL | `https://api.astrology-api.io` |
| `WithHTTPClient(c)` | Custom `*http.Client` | `nil` |
| `WithMaxRetries(n)` | Max retry attempts | `2` |
| `WithRetryDelay(d)` | Initial backoff delay | `500ms` |
| `WithRetryableStatusCodes(codes...)` | Status codes that trigger retry | `[408,429,500,502,503,504]` |
| `WithRequestTimeout(d)` | Per-request timeout | `30s` |
| `WithHeader(key, value)` | Extra request header | — |
| `WithResponseInto(resp)` | Capture raw `*http.Response` | — |

## Versioning & Releases

Versions are managed via **git tags** following [Semantic Versioning](https://semver.org/):

```
v{MAJOR}.{MINOR}.{PATCH}
```

| Change | Bump | Example |
|---|---|---|
| Bug fix, docs | PATCH | v0.1.0 → v0.1.1 |
| New category, new method (backwards-compatible) | MINOR | v0.1.1 → v0.2.0 |
| Breaking API change (removed/renamed types, changed signatures) | MAJOR | v0.2.0 → v1.0.0 |

### Creating a release

```bash
git tag v0.1.0
git push origin v0.1.0
```

CI will automatically:
1. Run tests
2. Create a GitHub Release with changelog
3. Notify Go Module Proxy so the version is available via `go get`

> **Note:** Starting from v2, the module path must include the major version suffix (`github.com/astro-api/astroapi-go/v2`), and all internal imports must be updated accordingly.

## Development

```bash
make build       # build all packages
make test        # run tests with race detector
make test-cover  # run tests with coverage report
make vet         # go vet
make fmt         # gofmt
make tidy        # go mod tidy
```

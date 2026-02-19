// Package astroapi provides the Go client for the Astrology API v3.
//
// # Quick Start
//
//	client := astroapi.NewClient(
//	    option.WithAPIKey("your-api-key"),
//	)
//
//	// Get current planetary positions
//	now, err := client.Data.GetNow(context.Background())
//
// # Environment Variables
//
//   - ASTROLOGY_API_KEY: your API key
//   - ASTROLOGY_API_BASE_URL: override the base URL (default: https://api.astrology-api.io)
package astroapi

import (
	"github.com/astro-api/astroapi-go/categories/data"
	"github.com/astro-api/astroapi-go/internal/apijson"
)

// Field is a generic optional/nullable field type. Use F() to set a value
// and Null() to explicitly set null.
type Field[T any] = apijson.Field[T]

// F creates a Field[T] with the given value.
func F[T any](v T) Field[T] { return apijson.F(v) }

// Null creates a Field[T] that marshals to JSON null.
func Null[T any]() Field[T] { return apijson.Null[T]() }

// String returns a Field[string] set to the given value.
func String(v string) Field[string] { return apijson.F(v) }

// Int returns a Field[int] set to the given value.
func Int(v int) Field[int] { return apijson.F(v) }

// Float returns a Field[float64] set to the given value.
func Float(v float64) Field[float64] { return apijson.F(v) }

// Bool returns a Field[bool] set to the given value.
func Bool(v bool) Field[bool] { return apijson.F(v) }

// DataPositionsParams is a convenience alias to avoid needing to import
// the data sub-package in simple usage.
type DataPositionsParams = data.PositionsParams

// DataLunarMetricsParams is a convenience alias.
type DataLunarMetricsParams = data.LunarMetricsParams

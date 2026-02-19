package data

import "github.com/astro-api/astroapi-go/shared"

// PositionsParams contains request parameters for planetary positions.
type PositionsParams struct {
	Subject shared.Subject    `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

// LunarMetricsParams contains request parameters for lunar metrics.
type LunarMetricsParams struct {
	DatetimeLocation shared.DateTimeLocation `json:"datetime_location" validate:"required"`
	Options          *shared.AstrologyOptions `json:"options,omitempty"`
}

// GlobalPositionsParams contains request parameters for global positions.
type GlobalPositionsParams struct {
	DatetimeLocation shared.DateTimeLocation `json:"datetime_location" validate:"required"`
	Options          *shared.AstrologyOptions `json:"options,omitempty"`
}

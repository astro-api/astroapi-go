package charts

import "github.com/astro-api/astroapi-go/shared"

type NatalChartParams struct {
	Subject shared.Subject    `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type CompositeChartParams struct {
	Subject1 shared.Subject    `json:"subject1" validate:"required"`
	Subject2 shared.Subject    `json:"subject2" validate:"required"`
	Options  *shared.AstrologyOptions `json:"options,omitempty"`
}

type SynastryChartParams struct {
	Subject1 shared.Subject    `json:"subject1" validate:"required"`
	Subject2 shared.Subject    `json:"subject2" validate:"required"`
	Options  *shared.AstrologyOptions `json:"options,omitempty"`
}

type TransitChartParams struct {
	NatalSubject     shared.Subject         `json:"natal_subject" validate:"required"`
	TransitDatetime  shared.DateTimeLocation `json:"transit_datetime" validate:"required"`
	Options          *shared.AstrologyOptions `json:"options,omitempty"`
}

type SolarReturnParams struct {
	Subject        shared.Subject          `json:"subject" validate:"required"`
	ReturnYear     int                     `json:"return_year" validate:"required"`
	ReturnLocation *shared.DateTimeLocation `json:"return_location,omitempty"`
	Options        *shared.AstrologyOptions `json:"options,omitempty"`
}

type LunarReturnParams struct {
	Subject        shared.Subject          `json:"subject" validate:"required"`
	ReturnDate     string                  `json:"return_date" validate:"required"`
	ReturnLocation *shared.DateTimeLocation `json:"return_location,omitempty"`
	Options        *shared.AstrologyOptions `json:"options,omitempty"`
}

type SolarReturnTransitsParams struct {
	Subject        shared.Subject          `json:"subject" validate:"required"`
	ReturnYear     int                     `json:"return_year" validate:"required"`
	DateRange      shared.DateRange        `json:"date_range"`
	Orb            float64                 `json:"orb,omitempty"`
	ReturnLocation *shared.DateTimeLocation `json:"return_location,omitempty"`
}

type LunarReturnTransitsParams struct {
	Subject        shared.Subject          `json:"subject" validate:"required"`
	ReturnDate     string                  `json:"return_date" validate:"required"`
	DateRange      shared.DateRange        `json:"date_range"`
	Orb            float64                 `json:"orb,omitempty"`
	ReturnLocation *shared.DateTimeLocation `json:"return_location,omitempty"`
}

type NatalTransitsParams struct {
	Subject   shared.Subject   `json:"subject" validate:"required"`
	DateRange *shared.DateRange `json:"date_range,omitempty"`
	Orb       float64          `json:"orb,omitempty"`
}

type ProgressionParams struct {
	Subject         shared.Subject          `json:"subject" validate:"required"`
	TargetDate      string                  `json:"target_date" validate:"required"`
	ProgressionType string                  `json:"progression_type" validate:"required"`
	Location        *shared.DateTimeLocation `json:"location,omitempty"`
	Options         *shared.AstrologyOptions `json:"options,omitempty"`
}

type DirectionParams struct {
	Subject       shared.Subject          `json:"subject" validate:"required"`
	TargetDate    string                  `json:"target_date" validate:"required"`
	DirectionType string                  `json:"direction_type" validate:"required"`
	ArcRate       float64                 `json:"arc_rate,omitempty"`
	Options       *shared.AstrologyOptions `json:"options,omitempty"`
}

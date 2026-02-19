package analysis

import "github.com/astro-api/astroapi-go/shared"

type NatalReportParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type SynastryReportParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
	Options  *shared.ReportOptions `json:"options,omitempty"`
}

type TransitReportParams struct {
	Subject   shared.Subject    `json:"subject" validate:"required"`
	DateRange *shared.DateRange `json:"date_range,omitempty"`
	Options   *shared.ReportOptions `json:"options,omitempty"`
}

type LunarReturnReportParams struct {
	Subject    shared.Subject `json:"subject" validate:"required"`
	ReturnDate string         `json:"return_date" validate:"required"`
	Options    *shared.ReportOptions `json:"options,omitempty"`
}

type SolarReturnReportParams struct {
	Subject    shared.Subject `json:"subject" validate:"required"`
	ReturnYear int            `json:"return_year" validate:"required"`
	Options    *shared.ReportOptions `json:"options,omitempty"`
}

type ProgressionReportParams struct {
	Subject         shared.Subject `json:"subject" validate:"required"`
	TargetDate      string         `json:"target_date" validate:"required"`
	ProgressionType string         `json:"progression_type" validate:"required"`
	Options         *shared.ReportOptions `json:"options,omitempty"`
}

type DirectionReportParams struct {
	Subject       shared.Subject `json:"subject" validate:"required"`
	TargetDate    string         `json:"target_date" validate:"required"`
	DirectionType string         `json:"direction_type" validate:"required"`
	ArcRate       float64        `json:"arc_rate,omitempty"`
}

type LunarAnalysisParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

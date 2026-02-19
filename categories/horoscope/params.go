package horoscope

import "github.com/astro-api/astroapi-go/shared"

type PersonalDailyParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type SignHoroscopeParams struct {
	Sign    string `json:"sign" validate:"required"`
	Date    string `json:"date,omitempty"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type SignWeeklyParams struct {
	Sign      string `json:"sign" validate:"required"`
	WeekStart string `json:"week_start,omitempty"`
	Options   *shared.ReportOptions `json:"options,omitempty"`
}

type SignMonthlyParams struct {
	Sign    string `json:"sign" validate:"required"`
	Month   int    `json:"month,omitempty"`
	Year    int    `json:"year,omitempty"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type SignYearlyParams struct {
	Sign    string `json:"sign" validate:"required"`
	Year    int    `json:"year,omitempty"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type PersonalTextParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Format  string         `json:"format,omitempty"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type ChineseHoroscopeParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

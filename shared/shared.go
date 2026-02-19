// Package shared contains common request/response types used across multiple
// Astrology API categories.
package shared

// BirthData represents the birth information for a person.
type BirthData struct {
	Year        int     `json:"year"`
	Month       int     `json:"month,omitempty"`
	Day         int     `json:"day,omitempty"`
	Hour        int     `json:"hour,omitempty"`
	Minute      int     `json:"minute,omitempty"`
	Second      int     `json:"second,omitempty"`
	City        string  `json:"city,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
}

// Subject represents a person for whom calculations are performed.
type Subject struct {
	Name      string    `json:"name,omitempty"`
	BirthData BirthData `json:"birth_data"`
	Email     string    `json:"email,omitempty"`
	Notes     string    `json:"notes,omitempty"`
}

// DateTimeLocation represents a point in time at a geographic location.
type DateTimeLocation struct {
	Year        int     `json:"year"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	Hour        int     `json:"hour"`
	Minute      int     `json:"minute"`
	Second      int     `json:"second,omitempty"`
	City        string  `json:"city,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
}

// DateRange represents a start and end date in ISO 8601 format (YYYY-MM-DD).
type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// AstrologyOptions provides common configuration for chart calculations.
type AstrologyOptions struct {
	HouseSystem             string   `json:"house_system,omitempty"`
	ZodiacType              string   `json:"zodiac_type,omitempty"`
	ActivePoints            []string `json:"active_points,omitempty"`
	Precision               int      `json:"precision,omitempty"`
	Language                string   `json:"language,omitempty"`
	Tradition               string   `json:"tradition,omitempty"`
	Perspective             string   `json:"perspective,omitempty"`
	DetailLevel             string   `json:"detail_level,omitempty"`
	IncludeInterpretations  bool     `json:"include_interpretations,omitempty"`
	IncludeRawData          bool     `json:"include_raw_data,omitempty"`
}

// ReportOptions contains options for text interpretation reports.
type ReportOptions struct {
	Tradition string `json:"tradition,omitempty"`
	Language  string `json:"language,omitempty"`
}

// GenericResponse is a catch-all map for API responses when a specific
// typed response struct is not defined.
type GenericResponse map[string]any

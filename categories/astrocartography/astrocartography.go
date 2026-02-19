// Package astrocartography provides the AstrocartographyClient for locational astrology.
package astrocartography

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/astrocartography"

type GenericResponse map[string]any

type LinesParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type LocationAnalysisParams struct {
	Subject  shared.Subject          `json:"subject" validate:"required"`
	Location shared.DateTimeLocation `json:"location" validate:"required"`
	Options  *shared.AstrologyOptions `json:"options,omitempty"`
}

type CompareLocationsParams struct {
	Subject   shared.Subject            `json:"subject" validate:"required"`
	Locations []shared.DateTimeLocation `json:"locations" validate:"required"`
}

type PowerZonesParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
}

type RelocationChartParams struct {
	Subject  shared.Subject          `json:"subject" validate:"required"`
	Location shared.DateTimeLocation `json:"location" validate:"required"`
}

type SearchLocationsParams struct {
	Subject  shared.Subject `json:"subject" validate:"required"`
	Criteria map[string]any `json:"criteria,omitempty"`
}

type MapParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options map[string]any `json:"options,omitempty"`
}

// Client provides access to the /api/v3/astrocartography endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new AstrocartographyClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetLines(ctx context.Context, params LinesParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lines"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateMap(ctx context.Context, params MapParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "map"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateParanMap(ctx context.Context, params MapParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "paran-map"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) AnalyzeLocation(ctx context.Context, params LocationAnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "location-analysis"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CompareLocations(ctx context.Context, params CompareLocationsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "compare-locations"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindPowerZones(ctx context.Context, params PowerZonesParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "power-zones"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) SearchLocations(ctx context.Context, params SearchLocationsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "search-locations"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateRelocationChart(ctx context.Context, params RelocationChartParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "relocation-chart"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLineMeanings(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "line-meanings"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetReport(ctx context.Context, params LinesParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

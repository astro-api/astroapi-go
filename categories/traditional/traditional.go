// Package traditional provides the TraditionalClient for traditional astrology.
package traditional

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/traditional"

type GenericResponse map[string]any

type AnalysisParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type ProfectionParams struct {
	Subject    shared.Subject `json:"subject" validate:"required"`
	TargetDate string         `json:"target_date,omitempty"`
}

type ProfectionTimelineParams struct {
	Subject    shared.Subject `json:"subject" validate:"required"`
	StartAge   int            `json:"start_age,omitempty"`
	EndAge     int            `json:"end_age,omitempty"`
}

// Client provides access to the /api/v3/traditional endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new TraditionalClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetDignitiesReport(ctx context.Context, params AnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "dignities"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetAnalysis(ctx context.Context, params AnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "analysis"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLotsAnalysis(ctx context.Context, params AnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lots"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetProfections(ctx context.Context, params ProfectionParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "profections"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetAnnualProfection(ctx context.Context, params ProfectionParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "profections", "annual"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetProfectionTimeline(ctx context.Context, params ProfectionTimelineParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "profections", "timeline"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetHorary(ctx context.Context, params AnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "horary"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCapabilities(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "capabilities"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

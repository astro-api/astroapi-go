// Package enhanced provides the EnhancedClient for enhanced astrological analysis.
package enhanced

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/enhanced"
const chartsPrefix = "api/v3/enhanced_charts"

type GenericResponse map[string]any

type GlobalAnalysisParams struct {
	DatetimeLocation shared.DateTimeLocation `json:"datetime_location" validate:"required"`
	Options          *shared.AstrologyOptions `json:"options,omitempty"`
}

type PersonalAnalysisParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

// Client provides access to the /api/v3/enhanced endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new EnhancedClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetGlobalAnalysis(ctx context.Context, params GlobalAnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "global"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPersonalAnalysis(ctx context.Context, params PersonalAnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "personal"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetGlobalAnalysisChart(ctx context.Context, params GlobalAnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(chartsPrefix, "global"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPersonalAnalysisChart(ctx context.Context, params PersonalAnalysisParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(chartsPrefix, "personal"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

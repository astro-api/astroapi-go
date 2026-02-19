// Package fixedstars provides the FixedStarsClient for fixed star calculations.
package fixedstars

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/fixed-stars"

type GenericResponse map[string]any

type PositionsParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Preset  string         `json:"preset,omitempty"`
	Orb     float64        `json:"orb,omitempty"`
}

type ConjunctionsParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Orb     float64        `json:"orb,omitempty"`
	Stars   []string       `json:"stars,omitempty"`
}

type ReportParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

// Client provides access to the /api/v3/fixed-stars endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new FixedStarsClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetList(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "list"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPositions(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "positions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetConjunctions(ctx context.Context, params ConjunctionsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "conjunctions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateReport(ctx context.Context, params ReportParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPresets(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "presets"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

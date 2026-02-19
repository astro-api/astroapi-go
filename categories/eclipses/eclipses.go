// Package eclipses provides the EclipsesClient for eclipse data.
package eclipses

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/eclipses"

type GenericResponse map[string]any

type UpcomingParams struct {
	Type      string `json:"type,omitempty" url:"type,omitempty"`
	Limit     int    `json:"limit,omitempty" url:"limit,omitempty"`
	AfterDate string `json:"after_date,omitempty" url:"after_date,omitempty"`
}

type NatalCheckParams struct {
	Subject     shared.Subject `json:"subject" validate:"required"`
	EclipseDate string         `json:"eclipse_date" validate:"required"`
}

type InterpretationParams struct {
	EclipseDate string          `json:"eclipse_date" validate:"required"`
	Subject     *shared.Subject `json:"subject,omitempty"`
}

// Client provides access to the /api/v3/eclipses endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new EclipsesClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetUpcoming(ctx context.Context, params *UpcomingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "upcoming"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetList(ctx context.Context, params *UpcomingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	return c.GetUpcoming(ctx, params, opts...)
}

func (c *Client) CheckNatalImpact(ctx context.Context, params NatalCheckParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal-check"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetInterpretation(ctx context.Context, params InterpretationParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "interpretation"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

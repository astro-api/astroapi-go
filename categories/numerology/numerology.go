// Package numerology provides the NumerologyClient for numerology calculations.
package numerology

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/numerology"

type GenericResponse map[string]any

type SingleSubjectParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type CompatibilityParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
}

// Client provides access to the /api/v3/numerology endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new NumerologyClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetReport(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "core-numbers"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCoreNumbers(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	return c.GetReport(ctx, params, opts...)
}

func (c *Client) GetComprehensiveReport(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "comprehensive"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCompatibility(ctx context.Context, params CompatibilityParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

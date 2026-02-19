// Package svg provides the SVGClient for generating astrological chart SVGs.
package svg

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/svg"

type NatalChartSVGParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Theme   string         `json:"theme,omitempty"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type SynastryChartSVGParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
	Theme    string         `json:"theme,omitempty"`
}

type CompositeChartSVGParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
	Theme    string         `json:"theme,omitempty"`
}

type TransitChartSVGParams struct {
	NatalSubject    shared.Subject          `json:"natal_subject" validate:"required"`
	TransitDatetime shared.DateTimeLocation `json:"transit_datetime" validate:"required"`
	Theme           string                  `json:"theme,omitempty"`
}

// Client provides access to the /api/v3/svg endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new SVGClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetChart(ctx context.Context, params NatalChartSVGParams, opts ...option.RequestOption) (string, error) {
	var out string
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal"), params, &out, opts...); err != nil {
		return "", err
	}
	return out, nil
}

func (c *Client) GetNatalChart(ctx context.Context, params NatalChartSVGParams, opts ...option.RequestOption) (string, error) {
	return c.GetChart(ctx, params, opts...)
}

func (c *Client) GetSynastryChart(ctx context.Context, params SynastryChartSVGParams, opts ...option.RequestOption) (string, error) {
	var out string
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "synastry"), params, &out, opts...); err != nil {
		return "", err
	}
	return out, nil
}

func (c *Client) GetCompositeChart(ctx context.Context, params CompositeChartSVGParams, opts ...option.RequestOption) (string, error) {
	var out string
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "composite"), params, &out, opts...); err != nil {
		return "", err
	}
	return out, nil
}

func (c *Client) GetTransitChart(ctx context.Context, params TransitChartSVGParams, opts ...option.RequestOption) (string, error) {
	var out string
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "transit"), params, &out, opts...); err != nil {
		return "", err
	}
	return out, nil
}

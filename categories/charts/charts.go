// Package charts provides the ChartsClient for generating astrological charts.
package charts

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
)

const apiPrefix = "api/v3/charts"

// Client provides access to the /api/v3/charts endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new ChartsClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetNatal(ctx context.Context, params NatalChartParams, opts ...option.RequestOption) (*NatalChartResponse, error) {
	var out NatalChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSynastry(ctx context.Context, params SynastryChartParams, opts ...option.RequestOption) (*SynastryChartResponse, error) {
	var out SynastryChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "synastry"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetComposite(ctx context.Context, params CompositeChartParams, opts ...option.RequestOption) (*CompositeChartResponse, error) {
	var out CompositeChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "composite"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetTransit(ctx context.Context, params TransitChartParams, opts ...option.RequestOption) (*TransitChartResponse, error) {
	var out TransitChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "transit"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSolarReturn(ctx context.Context, params SolarReturnParams, opts ...option.RequestOption) (*SolarReturnChartResponse, error) {
	var out SolarReturnChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "solar-return"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLunarReturn(ctx context.Context, params LunarReturnParams, opts ...option.RequestOption) (*LunarReturnChartResponse, error) {
	var out LunarReturnChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-return"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSolarReturnTransits(ctx context.Context, params SolarReturnTransitsParams, opts ...option.RequestOption) (*SolarReturnTransitsResponse, error) {
	var out SolarReturnTransitsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "solar-return-transits"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLunarReturnTransits(ctx context.Context, params LunarReturnTransitsParams, opts ...option.RequestOption) (*LunarReturnTransitsResponse, error) {
	var out LunarReturnTransitsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-return-transits"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetNatalTransits(ctx context.Context, params NatalTransitsParams, opts ...option.RequestOption) (*NatalTransitsResponse, error) {
	var out NatalTransitsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal-transits"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetProgressions(ctx context.Context, params ProgressionParams, opts ...option.RequestOption) (*ProgressionChartResponse, error) {
	var out ProgressionChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "progressions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetDirections(ctx context.Context, params DirectionParams, opts ...option.RequestOption) (*DirectionChartResponse, error) {
	var out DirectionChartResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "directions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

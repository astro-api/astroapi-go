// Package data provides the DataClient for accessing planetary positions,
// aspects, house cusps, and lunar metrics.
package data

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
)

const apiPrefix = "api/v3/data"

// Client provides access to the /api/v3/data endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new DataClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

// GetNow returns the current moment's astrological data.
func (c *Client) GetNow(ctx context.Context, opts ...option.RequestOption) (*NowResponse, error) {
	var out NowResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "now"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPositions returns planetary positions for a given subject.
func (c *Client) GetPositions(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*PositionsResponse, error) {
	var out PositionsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "positions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEnhancedPositions returns enhanced planetary positions for a given subject.
func (c *Client) GetEnhancedPositions(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*PositionsResponse, error) {
	var out PositionsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "positions", "enhanced"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetGlobalPositions returns global planetary positions for a given datetime/location.
func (c *Client) GetGlobalPositions(ctx context.Context, params GlobalPositionsParams, opts ...option.RequestOption) (*GlobalPositionsResponse, error) {
	var out GlobalPositionsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "global-positions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAspects returns astrological aspects for a given subject.
func (c *Client) GetAspects(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*AspectsResponse, error) {
	var out AspectsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "aspects"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEnhancedAspects returns enhanced astrological aspects for a given subject.
func (c *Client) GetEnhancedAspects(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*AspectsResponse, error) {
	var out AspectsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "aspects", "enhanced"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHouseCusps returns house cusp data for a given subject.
func (c *Client) GetHouseCusps(ctx context.Context, params PositionsParams, opts ...option.RequestOption) (*HouseCuspsResponse, error) {
	var out HouseCuspsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "house-cusps"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetLunarMetrics returns lunar metrics for a given datetime/location.
func (c *Client) GetLunarMetrics(ctx context.Context, params LunarMetricsParams, opts ...option.RequestOption) (*LunarMetricsResponse, error) {
	var out LunarMetricsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-metrics"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetEnhancedLunarMetrics returns enhanced lunar metrics for a given datetime/location.
func (c *Client) GetEnhancedLunarMetrics(ctx context.Context, params LunarMetricsParams, opts ...option.RequestOption) (*LunarMetricsResponse, error) {
	var out LunarMetricsResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-metrics", "enhanced"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

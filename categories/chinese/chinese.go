// Package chinese provides the ChineseClient for Chinese astrology calculations.
package chinese

import (
	"context"
	"fmt"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/chinese"

type GenericResponse map[string]any

type BaZiParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type CompatibilityParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
}

type LuckPillarsParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Count   int            `json:"count,omitempty"`
}

type SingleSubjectParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
}

type YearlyForecastParams struct {
	Subject    shared.Subject `json:"subject" validate:"required"`
	ForecastYear int          `json:"forecast_year,omitempty"`
}

// Client provides access to the /api/v3/chinese endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new ChineseClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) CalculateBaZi(ctx context.Context, params BaZiParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "bazi"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CalculateCompatibility(ctx context.Context, params CompatibilityParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CalculateLuckPillars(ctx context.Context, params LuckPillarsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "luck-pillars"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CalculateMingGua(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "ming-gua"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetYearlyForecast(ctx context.Context, params YearlyForecastParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "yearly-forecast"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSolarTerms(ctx context.Context, year int, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "calendar", "solar-terms", fmt.Sprintf("%d", year)), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetZodiacAnimal(ctx context.Context, animal string, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "zodiac", animal), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Package horoscope provides the HoroscopeClient for generating horoscope forecasts.
package horoscope

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
)

const apiPrefix = "api/v3/horoscope"

// Client provides access to the /api/v3/horoscope endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new HoroscopeClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetPersonalDaily(ctx context.Context, params PersonalDailyParams, opts ...option.RequestOption) (*PersonalDailyResponse, error) {
	var out PersonalDailyResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "personal", "daily"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPersonalDailyText(ctx context.Context, params PersonalTextParams, opts ...option.RequestOption) (*HoroscopeTextResponse, error) {
	var out HoroscopeTextResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "personal", "daily", "text"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignDaily(ctx context.Context, params SignHoroscopeParams, opts ...option.RequestOption) (*PersonalDailyResponse, error) {
	var out PersonalDailyResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "daily"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignDailyText(ctx context.Context, params SignHoroscopeParams, opts ...option.RequestOption) (*HoroscopeTextResponse, error) {
	var out HoroscopeTextResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "daily", "text"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignWeekly(ctx context.Context, params SignWeeklyParams, opts ...option.RequestOption) (*WeeklyHoroscopeResponse, error) {
	var out WeeklyHoroscopeResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "weekly"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignWeeklyText(ctx context.Context, params SignWeeklyParams, opts ...option.RequestOption) (*HoroscopeTextResponse, error) {
	var out HoroscopeTextResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "weekly", "text"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignMonthly(ctx context.Context, params SignMonthlyParams, opts ...option.RequestOption) (*MonthlyHoroscopeResponse, error) {
	var out MonthlyHoroscopeResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "monthly"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignMonthlyText(ctx context.Context, params SignMonthlyParams, opts ...option.RequestOption) (*HoroscopeTextResponse, error) {
	var out HoroscopeTextResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "monthly", "text"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSignYearly(ctx context.Context, params SignYearlyParams, opts ...option.RequestOption) (*YearlyHoroscopeResponse, error) {
	var out YearlyHoroscopeResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "sign", "yearly"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetChinese(ctx context.Context, params ChineseHoroscopeParams, opts ...option.RequestOption) (*ChineseHoroscopeResponse, error) {
	var out ChineseHoroscopeResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "chinese", "bazi"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

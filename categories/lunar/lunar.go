// Package lunar provides the LunarClient for lunar data.
package lunar

import (
	"context"
	"fmt"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/lunar"

type GenericResponse map[string]any

type PhasesParams struct {
	DateRange shared.DateRange `json:"date_range" validate:"required"`
	Options   *shared.AstrologyOptions `json:"options,omitempty"`
}

type EventsParams struct {
	DateRange shared.DateRange `json:"date_range" validate:"required"`
	Options   *shared.AstrologyOptions `json:"options,omitempty"`
}

type MansionsParams struct {
	DatetimeLocation shared.DateTimeLocation `json:"datetime_location" validate:"required"`
	MansionSystem    string                  `json:"mansion_system,omitempty"`
}

type VoidOfCourseParams struct {
	DateRange shared.DateRange `json:"date_range" validate:"required"`
}

type CalendarParams struct {
	IncludePhases bool   `json:"include_phases,omitempty" url:"include_phases,omitempty"`
	Language      string `json:"language,omitempty" url:"language,omitempty"`
}

// Client provides access to the /api/v3/lunar endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new LunarClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetPhase(ctx context.Context, params PhasesParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "phases"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCalendar(ctx context.Context, year int, params *CalendarParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "calendar", fmt.Sprintf("%d", year)), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetVoidOfCourse(ctx context.Context, params VoidOfCourseParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "void-of-course"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetEvents(ctx context.Context, params EventsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "events"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetMansions(ctx context.Context, params MansionsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "mansions"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

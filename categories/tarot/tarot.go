// Package tarot provides the TarotClient for tarot card readings and reports.
package tarot

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/tarot"

type GenericResponse map[string]any

type DrawCardsParams struct {
	Count       int    `json:"count" validate:"required"`
	SpreadType  string `json:"spread_type,omitempty"`
	Tradition   string `json:"tradition,omitempty"`
}

type CardParams struct {
	CardID string `json:"card_id" validate:"required"`
}

type SpreadParams struct {
	SpreadType string `json:"spread_type" validate:"required"`
	Subject    *shared.Subject `json:"subject,omitempty"`
}

type TarotReportParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.ReportOptions `json:"options,omitempty"`
}

type TarotSynastryParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
}

type SearchParams struct {
	Query     string `json:"query,omitempty" url:"query,omitempty"`
	Tradition string `json:"tradition,omitempty" url:"tradition,omitempty"`
}

type GlossaryParams struct {
	Category string `json:"category,omitempty" url:"category,omitempty"`
}

type DailyCardParams struct {
	Date      string `json:"date,omitempty" url:"date,omitempty"`
	Tradition string `json:"tradition,omitempty" url:"tradition,omitempty"`
}

// Client provides access to the /api/v3/tarot endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new TarotClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetDraw(ctx context.Context, params DrawCardsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "draw"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCard(ctx context.Context, cardID string, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "cards", cardID), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSpread(ctx context.Context, params SpreadParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "spread"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCardsGlossary(ctx context.Context, params *GlossaryParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "glossary", "cards"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSpreadsGlossary(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "glossary", "spreads"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetDailyCard(ctx context.Context, params *DailyCardParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "daily-card"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateReport(ctx context.Context, params TarotReportParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GenerateSynastryReport(ctx context.Context, params TarotSynastryParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "synastry-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) SearchCards(ctx context.Context, params *SearchParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "cards", "search"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

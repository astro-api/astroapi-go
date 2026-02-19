// Package glossary provides the GlossaryClient for reference data lookups.
package glossary

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
)

const apiPrefix = "api/v3/glossary"

type CitySearchParams struct {
	Search      string `json:"search,omitempty" url:"search,omitempty"`
	CountryCode string `json:"country_code,omitempty" url:"country_code,omitempty"`
	Limit       int    `json:"limit,omitempty" url:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty" url:"offset,omitempty"`
}

type HousesParams struct {
	HouseSystem string `json:"house_system,omitempty" url:"house_system,omitempty"`
}

type KeywordsParams struct {
	Category string `json:"category,omitempty" url:"category,omitempty"`
}

type LifeAreasParams struct {
	Language string `json:"language,omitempty" url:"language,omitempty"`
}

type ActivePointsParams struct {
	Type string `json:"type,omitempty" url:"type,omitempty"`
}

type GenericResponse map[string]any

// Client provides access to the /api/v3/glossary endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new GlossaryClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetActivePoints(ctx context.Context, params *ActivePointsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "active-points"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetPrimaryActivePoints(ctx context.Context, params *ActivePointsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "active-points", "primary"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCities(ctx context.Context, params *CitySearchParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "cities"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCountries(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "countries"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetElements(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "elements"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFixedStars(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "fixed-stars"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetHouseSystems(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "house-systems"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetHouses(ctx context.Context, params *HousesParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "houses"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetKeywords(ctx context.Context, params *KeywordsParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "keywords"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLanguages(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "languages"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLifeAreas(ctx context.Context, params *LifeAreasParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "life-areas"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetThemes(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "themes"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetZodiacTypes(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix, "zodiac-types"), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

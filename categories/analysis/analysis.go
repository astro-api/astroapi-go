// Package analysis provides the AnalysisClient for astrological reports and interpretation.
package analysis

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
)

const apiPrefix = "api/v3/analysis"

// Client provides access to the /api/v3/analysis endpoints.
type Client struct {
	*categories.BaseCategoryClient
}

// NewClient creates a new AnalysisClient.
func NewClient(base *categories.BaseCategoryClient) *Client {
	return &Client{BaseCategoryClient: base}
}

func (c *Client) GetNatalReport(ctx context.Context, params NatalReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSynastryReport(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "synastry-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCompositeReport(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "composite-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCompatibility(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCompatibilityScore(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*CompatibilityScoreResponse, error) {
	var out CompatibilityScoreResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "compatibility-score"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetRelationship(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "relationship"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetRelationshipScore(ctx context.Context, params SynastryReportParams, opts ...option.RequestOption) (*CompatibilityScoreResponse, error) {
	var out CompatibilityScoreResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "relationship-score"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetTransitReport(ctx context.Context, params TransitReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "transit-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetNatalTransitReport(ctx context.Context, params TransitReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "natal-transit-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetProgressionReport(ctx context.Context, params ProgressionReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "progression-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetDirectionReport(ctx context.Context, params DirectionReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "direction-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLunarReturnReport(ctx context.Context, params LunarReturnReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-return-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSolarReturnReport(ctx context.Context, params SolarReturnReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "solar-return-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetCareerAnalysis(ctx context.Context, params NatalReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "career"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetHealthAnalysis(ctx context.Context, params NatalReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "health"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetKarmicAnalysis(ctx context.Context, params NatalReportParams, opts ...option.RequestOption) (*ReportResponse, error) {
	var out ReportResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "karmic"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetLunarAnalysis(ctx context.Context, params LunarAnalysisParams, opts ...option.RequestOption) (*LunarAnalysisResponse, error) {
	var out LunarAnalysisResponse
	if err := c.Post(ctx, c.BuildURL(apiPrefix, "lunar-analysis"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

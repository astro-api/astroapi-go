// Package insights provides the InsightsClient with nested sub-clients for
// specialized insights: relationship, pet, wellness, financial, and business.
package insights

import (
	"context"

	"github.com/astro-api/astroapi-go/categories"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
)

const apiPrefix = "api/v3/insights"

type GenericResponse map[string]any

// ---- Common param types ----

type SingleSubjectParams struct {
	Subject shared.Subject `json:"subject" validate:"required"`
	Options *shared.AstrologyOptions `json:"options,omitempty"`
}

type TwoSubjectParams struct {
	Subject1 shared.Subject `json:"subject1" validate:"required"`
	Subject2 shared.Subject `json:"subject2" validate:"required"`
	Options  *shared.AstrologyOptions `json:"options,omitempty"`
}

type MultiSubjectParams struct {
	Subjects []shared.Subject `json:"subjects" validate:"required"`
	Options  *shared.AstrologyOptions `json:"options,omitempty"`
}

type TimingParams struct {
	Subject   shared.Subject    `json:"subject" validate:"required"`
	DateRange *shared.DateRange `json:"date_range,omitempty"`
}

// ---- Client ----

// Client is the root Insights client. It exposes nested sub-clients for
// specialized insight categories.
type Client struct {
	*categories.BaseCategoryClient
	Relationship *RelationshipClient
	Pet          *PetClient
	Wellness     *WellnessClient
	Financial    *FinancialClient
	Business     *BusinessClient
}

// NewClient creates a new InsightsClient with all nested sub-clients.
func NewClient(base *categories.BaseCategoryClient) *Client {
	c := &Client{BaseCategoryClient: base}
	c.Relationship = &RelationshipClient{base}
	c.Pet = &PetClient{base}
	c.Wellness = &WellnessClient{base}
	c.Financial = &FinancialClient{base}
	c.Business = &BusinessClient{base}
	return c
}

// Discover returns the available insights endpoints.
func (c *Client) Discover(ctx context.Context, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Get(ctx, c.BuildURL(apiPrefix), nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ---- Relationship sub-client ----

type RelationshipClient struct {
	*categories.BaseCategoryClient
}

const relPrefix = "api/v3/insights/relationship"

func (c *RelationshipClient) GetCompatibility(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipClient) GetCompatibilityScore(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "compatibility-score"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipClient) GetLoveLanguages(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "love-languages"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipClient) GetDavisonReport(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "davison-report"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipClient) GetTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RelationshipClient) GetRedFlags(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(relPrefix, "red-flags"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ---- Pet sub-client ----

type PetClient struct {
	*categories.BaseCategoryClient
}

const petPrefix = "api/v3/insights/pet"

func (c *PetClient) GetPersonality(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(petPrefix, "personality"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PetClient) GetCompatibility(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(petPrefix, "compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PetClient) GetTrainingWindows(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(petPrefix, "training-windows"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PetClient) GetHealthSensitivities(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(petPrefix, "health-sensitivities"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PetClient) GetMultiPetDynamics(ctx context.Context, params MultiSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(petPrefix, "multi-pet-dynamics"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ---- Wellness sub-client ----

type WellnessClient struct {
	*categories.BaseCategoryClient
}

const wellnessPrefix = "api/v3/insights/wellness"

func (c *WellnessClient) GetBodyMapping(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "body-mapping"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WellnessClient) GetBiorhythms(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "biorhythms"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WellnessClient) GetWellnessTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WellnessClient) GetEnergyPatterns(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "energy-patterns"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WellnessClient) GetWellnessScore(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "score"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WellnessClient) GetMoonWellness(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(wellnessPrefix, "moon-wellness"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ---- Financial sub-client ----

type FinancialClient struct {
	*categories.BaseCategoryClient
}

const financialPrefix = "api/v3/insights/financial"

func (c *FinancialClient) GetMarketTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "market-timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FinancialClient) AnalyzePersonalTrading(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "personal-trading"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FinancialClient) GetGannAnalysis(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "gann-analysis"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FinancialClient) GetBradleySiderograph(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "bradley-siderograph"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FinancialClient) GetCryptoTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "crypto-timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FinancialClient) GetForexTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(financialPrefix, "forex-timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ---- Business sub-client ----

type BusinessClient struct {
	*categories.BaseCategoryClient
}

const businessPrefix = "api/v3/insights/business"

func (c *BusinessClient) GetTeamDynamics(ctx context.Context, params MultiSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "team-dynamics"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessClient) GetHiringCompatibility(ctx context.Context, params TwoSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "hiring-compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessClient) GetLeadershipStyle(ctx context.Context, params SingleSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "leadership-style"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessClient) GetBusinessTiming(ctx context.Context, params TimingParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "timing"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessClient) GetDepartmentCompatibility(ctx context.Context, params MultiSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "department-compatibility"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessClient) GetSuccessionPlanning(ctx context.Context, params MultiSubjectParams, opts ...option.RequestOption) (*GenericResponse, error) {
	var out GenericResponse
	if err := c.Post(ctx, c.BuildURL(businessPrefix, "succession-planning"), params, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

package insights_test

import (
	"context"
	"net/http"
	"testing"

	astroapi "github.com/astro-api/astroapi-go"
	"github.com/astro-api/astroapi-go/categories/insights"
	"github.com/astro-api/astroapi-go/internal/testutil"
	"github.com/astro-api/astroapi-go/option"
	"github.com/astro-api/astroapi-go/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestInsightsClient_Discover(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"categories": []any{}}))
	})
	defer cleanup()

	result, err := client.Insights.Discover(ctx)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"score": 90}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetCompatibility(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_ValidationError(t *testing.T) {
	client := astroapi.NewClient(option.WithAPIKey("test-key"), option.WithMaxRetries(0))
	_, err := client.Insights.Relationship.GetCompatibility(ctx, insights.TwoSubjectParams{})
	require.Error(t, err)
}

func TestInsightsClient_Pet_GetPersonality(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/pet/personality", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"traits": []any{}}))
	})
	defer cleanup()

	result, err := client.Insights.Pet.GetPersonality(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetBodyMapping(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/body-mapping", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"areas": []any{}}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetBodyMapping(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_GetMarketTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/market-timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"windows": []any{}}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.GetMarketTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetTeamDynamics(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/team-dynamics", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{"dynamics": []any{}}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetTeamDynamics(ctx, insights.MultiSubjectParams{
		Subjects: []shared.Subject{testutil.DefaultSubject()},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetCompatibilityScore(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/compatibility-score", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetCompatibilityScore(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetLoveLanguages(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/love-languages", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetLoveLanguages(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetDavisonReport(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/davison-report", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetDavisonReport(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Relationship_GetRedFlags(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/relationship/red-flags", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Relationship.GetRedFlags(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Pet_GetCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/pet/compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Pet.GetCompatibility(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Pet_GetTrainingWindows(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/pet/training-windows", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Pet.GetTrainingWindows(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Pet_GetHealthSensitivities(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/pet/health-sensitivities", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Pet.GetHealthSensitivities(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Pet_GetMultiPetDynamics(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/pet/multi-pet-dynamics", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Pet.GetMultiPetDynamics(ctx, insights.MultiSubjectParams{
		Subjects: []shared.Subject{testutil.DefaultSubject()},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetBiorhythms(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/biorhythms", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetBiorhythms(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetWellnessTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetWellnessTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetEnergyPatterns(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/energy-patterns", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetEnergyPatterns(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetWellnessScore(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/score", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetWellnessScore(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Wellness_GetMoonWellness(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/wellness/moon-wellness", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Wellness.GetMoonWellness(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_AnalyzePersonalTrading(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/personal-trading", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.AnalyzePersonalTrading(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_GetGannAnalysis(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/gann-analysis", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.GetGannAnalysis(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_GetBradleySiderograph(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/bradley-siderograph", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.GetBradleySiderograph(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_GetCryptoTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/crypto-timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.GetCryptoTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Financial_GetForexTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/financial/forex-timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Financial.GetForexTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetHiringCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/hiring-compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetHiringCompatibility(ctx, insights.TwoSubjectParams{
		Subject1: testutil.DefaultSubject(),
		Subject2: testutil.DefaultSubject2(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetLeadershipStyle(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/leadership-style", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetLeadershipStyle(ctx, insights.SingleSubjectParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetBusinessTiming(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/timing", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetBusinessTiming(ctx, insights.TimingParams{
		Subject: testutil.DefaultSubject(),
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetDepartmentCompatibility(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/department-compatibility", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetDepartmentCompatibility(ctx, insights.MultiSubjectParams{
		Subjects: []shared.Subject{testutil.DefaultSubject()},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Business_GetSuccessionPlanning(t *testing.T) {
	client, cleanup := testutil.NewClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v3/insights/business/succession-planning", r.URL.Path)
		testutil.JSON(w, testutil.DataEnvelope(map[string]any{}))
	})
	defer cleanup()

	result, err := client.Insights.Business.GetSuccessionPlanning(ctx, insights.MultiSubjectParams{
		Subjects: []shared.Subject{testutil.DefaultSubject()},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestInsightsClient_Integration(t *testing.T) {
	client := testutil.NewIntegrationClient(t)

	t.Run("Discover", func(t *testing.T) {
		result, err := client.Insights.Discover(ctx)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Relationship_GetCompatibility", func(t *testing.T) {
		result, err := client.Insights.Relationship.GetCompatibility(ctx, insights.TwoSubjectParams{
			Subject1: testutil.DefaultSubject(),
			Subject2: testutil.DefaultSubject2(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Wellness_GetBiorhythms", func(t *testing.T) {
		result, err := client.Insights.Wellness.GetBiorhythms(ctx, insights.SingleSubjectParams{
			Subject: testutil.DefaultSubject(),
		})
		require.NoError(t, err)
		assert.NotNil(t, result)
	})
}

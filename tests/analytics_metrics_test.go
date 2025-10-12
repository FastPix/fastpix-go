package tests

import (
	"context"
	"testing"

	"github.com/FastPix/fastpix-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyticsAndViews(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test List Video Views", func(t *testing.T) {
		// Test listing video views
		request := operations.ListVideoViewsRequest{
			Timespan: operations.ListVideoViewsTimespanTwentyFourhours,
			Limit:    int64Ptr(10),
			Offset:   int64Ptr(1),
		}

		response, err := config.Client.Views.ListVideoViews(ctx, request)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Get Video View Details", func(t *testing.T) {
		// First, get a list of views to find a valid ID
		listRequest := operations.ListVideoViewsRequest{
			Timespan: operations.ListVideoViewsTimespanTwentyFourhours,
			Limit:    int64Ptr(1),
		}

		listResponse, err := config.Client.Views.ListVideoViews(ctx, listRequest)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			viewID := listResponse.Object.Data[0].ViewID
			require.NotEmpty(t, viewID)

			// Test getting specific view details
			response, err := config.Client.Views.GetVideoViewDetails(ctx, viewID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No views found to test GetVideoViewDetails")
		}
	})

	t.Run("Test List by Top Content", func(t *testing.T) {
		// Test listing top content
		response, err := config.Client.Views.ListByTopContent(
			ctx,
			operations.ListByTopContentTimespanTwentyFourhours,
			nil, // filterby
			int64Ptr(10), // limit
		)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Get Timeseries Views", func(t *testing.T) {
		// Test getting timeseries views data
		response, err := config.Client.Views.GetDataViewlistCurrentViewsGetTimeseriesViews(ctx)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Filter Views", func(t *testing.T) {
		// Test filtering views
		dimension := operations.GetDataViewlistCurrentViewsFilterDimensionOsName
		response, err := config.Client.Views.GetDataViewlistCurrentViewsFilter(
			ctx,
			&dimension,
			int64Ptr(10),
		)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})
}

func TestMetrics(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test List Breakdown Values", func(t *testing.T) {
		// Test listing breakdown values
		measurement := "count"
		request := operations.ListBreakdownValuesRequest{
			MetricID:    operations.ListBreakdownValuesMetricIDViews,
			Timespan:    operations.ListBreakdownValuesTimespanTwentyFourhours,
			Limit:       int64Ptr(10),
			Offset:      int64Ptr(1),
			Measurement: &measurement,
		}

		response, err := config.Client.Metrics.ListBreakdownValues(ctx, request)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test List Overall Values", func(t *testing.T) {
		// Test listing overall values
		measurement := "count"
		response, err := config.Client.Metrics.ListOverallValues(
			ctx,
			operations.ListOverallValuesMetricIDViews,
			operations.ListOverallValuesTimespanTwentyFourhours,
			&measurement, // measurement
			nil, // filterby
		)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test Get Timeseries Data", func(t *testing.T) {
		// Test getting timeseries data
		groupBy := operations.GroupByHour
		measurement := "count"
		request := operations.GetTimeseriesDataRequest{
			MetricID:    operations.GetTimeseriesDataMetricIDViews,
			Timespan:    operations.GetTimeseriesDataTimespanTwentyFourhours,
			GroupBy:     &groupBy,
			Measurement: &measurement,
		}

		response, err := config.Client.Metrics.GetTimeseriesData(ctx, request)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test List Comparison Values", func(t *testing.T) {
		// Test listing comparison values
		dimension := operations.ListComparisonValuesDimensionBrowserName
		value := "Chrome"
		response, err := config.Client.Metrics.ListComparisonValues(
			ctx,
			operations.ListComparisonValuesTimespanTwentyFourhours,
			nil, // filterby
			&dimension, // dimension
			&value, // value
		)
		// This test may fail due to no data available, which is expected for trial accounts
		if err != nil {
			t.Logf("Comparison values test failed (expected for trial account): %v", err)
			t.Skip("No comparison data available for trial account")
		}
		assert.NotNil(t, response)
	})
}

func TestDimensions(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test List Dimensions", func(t *testing.T) {
		// Test listing dimensions
		response, err := config.Client.Dimensions.ListDimensions(ctx)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})

	t.Run("Test List Filter Values for Dimension", func(t *testing.T) {
		// Test listing filter values for a specific dimension
		response, err := config.Client.Dimensions.ListFilterValuesForDimension(
			ctx,
			operations.DimensionsIDBrowserName,
			operations.ListFilterValuesForDimensionTimespanTwentyFourhours,
			nil, // filterby
		)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})
}

func TestErrors(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test List Errors", func(t *testing.T) {
		// Test listing errors
		response, err := config.Client.Errors.ListErrors(
			ctx,
			operations.ListErrorsTimespanTwentyFourhours,
			nil, // filterby
			int64Ptr(10), // limit
		)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.Object)
	})
}

// Helper function to create int64 pointers
func int64Ptr(v int64) *int64 {
	return &v
}

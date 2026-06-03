//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
)

const dataPoints = "Data points: %d\n"

func main() {
	ctx := context.Background()

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	fmt.Println("=== Video Views Analytics ===")
	fmt.Println("\n--- Listing Video Views ---")

	viewsRequest := operations.ListVideoViewsRequest{
		Timespan: operations.ListVideoViewsTimespanTwentyFourhours,
		Limit:    fastpixgo.Int64(10),
		Offset:   fastpixgo.Int64(0),
	}

	viewsResponse, err := client.Views.ListVideoViews(ctx, viewsRequest)
	if err != nil {
		log.Printf("Error listing video views: %v", err)
	} else {
		printVideoViews(viewsResponse)
	}

	if viewsResponse.Object != nil && len(viewsResponse.Object.Data) > 0 {
		printViewDetails(ctx, client, viewsResponse.Object.Data[0].ViewID)
	}

	listTopContent(ctx, client)
	getTimeseriesViews(ctx, client)
	getFilteredViews(ctx, client)
	runMetricsAnalytics(ctx, client)
	runDimensionsAnalytics(ctx, client)
	runErrorAnalytics(ctx, client)
	runAdvancedAnalyticsQueries(ctx, client)
	runAnalyticsDashboardSummary(ctx, client)

	fmt.Println("\n=== Analytics Demo Complete ===")
}

func printVideoViews(viewsResponse *operations.ListVideoViewsResponse) {
	fmt.Printf("Found %d video views in the last 24 hours:\n", len(viewsResponse.Object.Data))
	for i, view := range viewsResponse.Object.Data {
		fmt.Printf("  %d. View ID: %s, Media ID: %s, Duration: %d seconds\n",
			i+1, view.ViewID, getStringValue(view.MediaID), getInt64Value(view.Duration))
	}
}

func printViewDetails(ctx context.Context, client *fastpixgo.FastPix, viewID string) {
	fmt.Printf("\n--- Getting View Details: %s ---\n", viewID)

	viewDetailsResponse, err := client.Views.GetVideoViewDetails(ctx, viewID)
	if err != nil {
		log.Printf("Error getting view details: %v", err)
		return
	}

	view := viewDetailsResponse.Object.Data
	fmt.Printf("View ID: %s\n", view.ViewID)
	fmt.Printf("Media ID: %s\n", getStringValue(view.MediaID))
	fmt.Printf("Duration: %d seconds\n", getInt64Value(view.Duration))
	fmt.Printf("Started: %s\n", getStringValue(view.StartedAt))
	fmt.Printf("Ended: %s\n", getStringValue(view.EndedAt))
}

func listTopContent(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Top Content ---")

	topContentResponse, err := client.Views.ListByTopContent(
		ctx,
		operations.ListByTopContentTimespanTwentyFourhours,
		nil,
		fastpixgo.Int64(10),
	)
	if err != nil {
		log.Printf("Error listing top content: %v", err)
		return
	}

	fmt.Printf("Found %d top content items:\n", len(topContentResponse.Object.Data))
	for i, content := range topContentResponse.Object.Data {
		fmt.Printf("  %d. Media ID: %s, Views: %d\n",
			i+1, getStringValue(content.MediaID), getInt64Value(content.Views))
	}
}

func getTimeseriesViews(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Timeseries Views ---")

	timeseriesResponse, err := client.Views.GetDataViewlistCurrentViewsGetTimeseriesViews(ctx)
	if err != nil {
		log.Printf("Error getting timeseries views: %v", err)
		return
	}

	fmt.Printf("Timeseries views data retrieved successfully!\n")
	fmt.Printf(dataPoints, len(timeseriesResponse.Object.Data))
}

func getFilteredViews(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Filtered Views by OS ---")

	dimension := operations.GetDataViewlistCurrentViewsFilterDimensionOsName
	filterResponse, err := client.Views.GetDataViewlistCurrentViewsFilter(
		ctx,
		&dimension,
		fastpixgo.Int64(10),
	)
	if err != nil {
		log.Printf("Error filtering views: %v", err)
		return
	}

	fmt.Printf("Filtered views data retrieved successfully!\n")
	fmt.Printf(dataPoints, len(filterResponse.Object.Data))
}

func runMetricsAnalytics(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Metrics Analytics ===")

	listBreakdownValues(ctx, client)
	listOverallValues(ctx, client)
	getTimeseriesData(ctx, client)
	listComparisonValues(ctx, client)
}

func listBreakdownValues(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Breakdown Values ---")

	breakdownRequest := operations.ListBreakdownValuesRequest{
		MetricID:    operations.ListBreakdownValuesMetricIDViews,
		Timespan:    operations.ListBreakdownValuesTimespanTwentyFourhours,
		Limit:       fastpixgo.Int64(10),
		Offset:      fastpixgo.Int64(0),
		Measurement: fastpixgo.Pointer("count"),
	}

	breakdownResponse, err := client.Metrics.ListBreakdownValues(ctx, breakdownRequest)
	if err != nil {
		log.Printf("Error listing breakdown values: %v", err)
		return
	}

	fmt.Printf("Found %d breakdown values:\n", len(breakdownResponse.Object.Data))
	for i, value := range breakdownResponse.Object.Data {
		fmt.Printf("  %d. Dimension: %s, Value: %s, Count: %d\n",
			i+1, getStringValue(value.Dimension), getStringValue(value.Value), getInt64Value(value.Count))
	}
}

func listOverallValues(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Overall Values ---")

	overallResponse, err := client.Metrics.ListOverallValues(
		ctx,
		operations.ListOverallValuesMetricIDViews,
		operations.ListOverallValuesTimespanTwentyFourhours,
		fastpixgo.Pointer("count"),
		nil,
	)
	if err != nil {
		log.Printf("Error listing overall values: %v", err)
		return
	}

	fmt.Printf("Overall values retrieved successfully!\n")
	fmt.Printf("Total views: %d\n", getInt64Value(overallResponse.Object.Data.Total))
}

func getTimeseriesData(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Timeseries Data ---")

	groupBy := operations.GroupByHour
	timeseriesRequest := operations.GetTimeseriesDataRequest{
		MetricID:    operations.GetTimeseriesDataMetricIDViews,
		Timespan:    operations.GetTimeseriesDataTimespanTwentyFourhours,
		GroupBy:     &groupBy,
		Measurement: fastpixgo.Pointer("count"),
	}

	timeseriesDataResponse, err := client.Metrics.GetTimeseriesData(ctx, timeseriesRequest)
	if err != nil {
		log.Printf("Error getting timeseries data: %v", err)
		return
	}

	fmt.Printf("Timeseries data retrieved successfully!\n")
	fmt.Printf(dataPoints, len(timeseriesDataResponse.Object.Data))
}

func listComparisonValues(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Comparison Values ---")

	dimensionComp := operations.ListComparisonValuesDimensionBrowserName
	value := "Chrome"
	comparisonResponse, err := client.Metrics.ListComparisonValues(
		ctx,
		operations.ListComparisonValuesTimespanTwentyFourhours,
		nil,
		&dimensionComp,
		&value,
	)
	if err != nil {
		log.Printf("Error listing comparison values: %v", err)
		return
	}

	fmt.Printf("Comparison values retrieved successfully!\n")
	fmt.Printf(dataPoints, len(comparisonResponse.Object.Data))
}

func runDimensionsAnalytics(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Dimensions ===")

	listDimensions(ctx, client)
	listFilterValuesForDimension(ctx, client)
}

func listDimensions(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Available Dimensions ---")

	dimensionsResponse, err := client.Dimensions.ListDimensions(ctx)
	if err != nil {
		log.Printf("Error listing dimensions: %v", err)
		return
	}

	fmt.Printf("Found %d dimensions:\n", len(dimensionsResponse.Object.Data))
	for i, dimension := range dimensionsResponse.Object.Data {
		fmt.Printf("  %d. ID: %s, Name: %s\n",
			i+1, getStringValue(dimension.ID), getStringValue(dimension.Name))
	}
}

func listFilterValuesForDimension(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n--- Filter Values for Browser Name ---")

	filterValuesResponse, err := client.Dimensions.ListFilterValuesForDimension(
		ctx,
		operations.DimensionsIDBrowserName,
		operations.ListFilterValuesForDimensionTimespanTwentyFourhours,
		nil,
	)
	if err != nil {
		log.Printf("Error listing filter values: %v", err)
		return
	}

	fmt.Printf("Found %d filter values for browser name:\n", len(filterValuesResponse.Object.Data))
	for i, value := range filterValuesResponse.Object.Data {
		fmt.Printf("  %d. Value: %s, Count: %d\n",
			i+1, getStringValue(value.Value), getInt64Value(value.Count))
	}
}

func runErrorAnalytics(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Error Analytics ===")
	fmt.Println("\n--- Recent Errors ---")

	errorsResponse, err := client.Errors.ListErrors(
		ctx,
		operations.ListErrorsTimespanTwentyFourhours,
		nil,
		fastpixgo.Int64(10),
	)
	if err != nil {
		log.Printf("Error listing errors: %v", err)
		return
	}

	fmt.Printf("Found %d errors in the last 24 hours:\n", len(errorsResponse.Object.Data))
	for i, error := range errorsResponse.Object.Data {
		fmt.Printf("  %d. Error ID: %s, Type: %s, Message: %s\n",
			i+1, error.ErrorID, getStringValue(error.Type), getStringValue(error.Message))
	}
}

func runAdvancedAnalyticsQueries(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Advanced Analytics Queries ===")

	timespans := []operations.ListVideoViewsTimespan{
		operations.ListVideoViewsTimespanOnehour,
		operations.ListVideoViewsTimespanTwentyFourhours,
		operations.ListVideoViewsTimespanSevendays,
		operations.ListVideoViewsTimespanThirtydays,
	}

	for _, timespan := range timespans {
		printViewsForTimespan(ctx, client, timespan)
	}
}

func printViewsForTimespan(ctx context.Context, client *fastpixgo.FastPix, timespan operations.ListVideoViewsTimespan) {
	fmt.Printf("\n--- Views for %s ---\n", timespan)

	response, err := client.Views.ListVideoViews(ctx, operations.ListVideoViewsRequest{
		Timespan: timespan,
		Limit:    fastpixgo.Int64(5),
	})
	if err != nil {
		log.Printf("Error getting views for %s: %v", timespan, err)
		return
	}

	fmt.Printf("Found %d views\n", len(response.Object.Data))
}

func runAnalyticsDashboardSummary(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Analytics Dashboard Summary ===")
	fmt.Println("Generating analytics summary...")

	overall, err := client.Metrics.ListOverallValues(
		ctx,
		operations.ListOverallValuesMetricIDViews,
		operations.ListOverallValuesTimespanTwentyFourhours,
		fastpixgo.Pointer("count"),
		nil,
	)
	if err == nil {
		fmt.Printf("Total views (24h): %d\n", getInt64Value(overall.Object.Data.Total))
	}

	topContent, err := client.Views.ListByTopContent(
		ctx,
		operations.ListByTopContentTimespanTwentyFourhours,
		nil,
		fastpixgo.Int64(5),
	)
	if err == nil {
		fmt.Printf("Top 5 content items: %d items\n", len(topContent.Object.Data))
	}

	errors, err := client.Errors.ListErrors(
		ctx,
		operations.ListErrorsTimespanTwentyFourhours,
		nil,
		fastpixgo.Int64(5),
	)
	if err == nil {
		fmt.Printf("Recent errors: %d errors\n", len(errors.Object.Data))
	}
}

func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func getInt64Value(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
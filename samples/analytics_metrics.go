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
	"github.com/FastPix/fastpix-go/optionalnullable"
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
		Timespan: operations.ListVideoViewsTimespanTwentyFourhours.ToPointer(),
		Limit:    fastpixgo.Int64(10),
		Offset:   fastpixgo.Int64(1),
	}

	viewsResponse, err := client.Views.ListVideoViews(ctx, viewsRequest)
	if err != nil {
		log.Printf("Error listing video views: %v", err)
	} else {
		printVideoViews(viewsResponse)
	}

	if viewsResponse != nil && viewsResponse.Object != nil && len(viewsResponse.Object.Data) > 0 {
		printViewDetails(ctx, client, getStringValue(viewsResponse.Object.Data[0].ViewID))
	}

	listTopContent(ctx, client)
	getFilteredViews(ctx, client)
	runMetricsAnalytics(ctx, client)
	runDimensionsAnalytics(ctx, client)
	runAdvancedAnalyticsQueries(ctx, client)
	runAnalyticsDashboardSummary(ctx, client)

	fmt.Println("\n=== Analytics Demo Complete ===")
}

func printVideoViews(viewsResponse *operations.ListVideoViewsResponse) {
	if viewsResponse.Object == nil {
		return
	}
	fmt.Printf("Found %d video views in the last 24 hours:\n", len(viewsResponse.Object.Data))
	for i, view := range viewsResponse.Object.Data {
		fmt.Printf("  %d. View ID: %s, Title: %s, Watch Time: %.0f\n",
			i+1, getStringValue(view.ViewID),
			getOptString(view.VideoTitle), getOptFloat(view.ViewWatchTime))
	}
}

func printViewDetails(ctx context.Context, client *fastpixgo.Fastpixgo, viewID string) {
	fmt.Printf("\n--- Getting View Details: %s ---\n", viewID)

	viewDetailsResponse, err := client.Views.GetDetails(ctx, viewID)
	if err != nil {
		log.Printf("Error getting view details: %v", err)
		return
	}
	if viewDetailsResponse.Object == nil || viewDetailsResponse.Object.Data == nil {
		return
	}

	view := viewDetailsResponse.Object.Data
	fmt.Printf("View ID: %s\n", getStringValue(view.ViewID))
	fmt.Printf("Media ID: %s\n", getOptString(view.MediaID))
	fmt.Printf("Video Title: %s\n", getOptString(view.VideoTitle))
	fmt.Printf("View Start: %s\n", getOptString(view.ViewStart))
	fmt.Printf("View End: %s\n", getOptString(view.ViewEnd))
}

func listTopContent(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Top Content ---")

	topContentResponse, err := client.Views.ListByTopContent(
		ctx,
		operations.ListByTopContentTimespanTwentyFourhours.ToPointer(),
		nil,
		fastpixgo.Int64(10),
	)
	if err != nil {
		log.Printf("Error listing top content: %v", err)
		return
	}
	if topContentResponse.Object == nil {
		return
	}

	fmt.Printf("Found %d top content items:\n", len(topContentResponse.Object.Data))
	for i, content := range topContentResponse.Object.Data {
		fmt.Printf("  %d. Title: %s, Views: %d, Unique Views: %d\n",
			i+1, getStringValue(content.VideoTitle),
			getInt64Value(content.Views), getInt64Value(content.UniqueViews))
	}
}

func getFilteredViews(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Filtered Views by OS ---")

	filterResponse, err := client.Views.ListVideoViews(ctx, operations.ListVideoViewsRequest{
		Timespan: operations.ListVideoViewsTimespanTwentyFourhours.ToPointer(),
		Filterby: fastpixgo.String("os_name:macOS"),
		Limit:    fastpixgo.Int64(10),
	})
	if err != nil {
		log.Printf("Error filtering views: %v", err)
		return
	}
	if filterResponse.Object == nil {
		return
	}

	fmt.Printf("Filtered views data retrieved successfully!\n")
	fmt.Printf(dataPoints, len(filterResponse.Object.Data))
}

func runMetricsAnalytics(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Metrics Analytics ===")

	listBreakdownValues(ctx, client)
	listOverallValues(ctx, client)
	getTimeseriesData(ctx, client)
	listComparisonValues(ctx, client)
}

func listBreakdownValues(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Breakdown Values ---")

	breakdownRequest := operations.ListBreakdownValuesRequest{
		MetricID:    operations.ListBreakdownValuesMetricIDViews,
		Timespan:    operations.ListBreakdownValuesTimespanTwentyFourhours.ToPointer(),
		GroupBy:     fastpixgo.String("browser_name"),
		Limit:       fastpixgo.Int64(10),
		Offset:      fastpixgo.Int64(1),
		Measurement: fastpixgo.String("count"),
	}

	breakdownResponse, err := client.Metrics.ListBreakdownValues(ctx, breakdownRequest)
	if err != nil {
		log.Printf("Error listing breakdown values: %v", err)
		return
	}
	if breakdownResponse.Object == nil {
		return
	}

	fmt.Printf("Found %d breakdown values:\n", len(breakdownResponse.Object.Data))
	for i, value := range breakdownResponse.Object.Data {
		fmt.Printf("  %d. Field: %s, Value: %.2f, Views: %d\n",
			i+1, getOptString(value.Field),
			getOptFloat(value.Value), getOptInt64(value.Views))
	}
}

func listOverallValues(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Overall Values ---")

	overallResponse, err := client.Metrics.ListOverallValues(
		ctx,
		operations.ListOverallValuesMetricIDViews,
		fastpixgo.String("count"),
		operations.ListOverallValuesTimespanTwentyFourhours.ToPointer(),
		nil,
	)
	if err != nil {
		log.Printf("Error listing overall values: %v", err)
		return
	}
	if overallResponse.Object == nil || overallResponse.Object.Data == nil {
		return
	}

	fmt.Printf("Overall values retrieved successfully!\n")
	fmt.Printf("Total views: %d\n", getOptInt64(overallResponse.Object.Data.TotalViews))
}

func getTimeseriesData(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Timeseries Data ---")

	timeseriesRequest := operations.GetTimeseriesDataRequest{
		MetricID:    operations.GetTimeseriesDataMetricIDViews,
		Timespan:    operations.GetTimeseriesDataTimespanTwentyFourhours.ToPointer(),
		GroupBy:     operations.GroupByHour.ToPointer(),
		Measurement: fastpixgo.String("count"),
	}

	timeseriesDataResponse, err := client.Metrics.GetTimeseriesData(ctx, timeseriesRequest)
	if err != nil {
		log.Printf("Error getting timeseries data: %v", err)
		return
	}
	if timeseriesDataResponse.Object == nil {
		return
	}

	fmt.Printf("Timeseries data retrieved successfully!\n")
	fmt.Printf(dataPoints, len(timeseriesDataResponse.Object.Data))
}

func listComparisonValues(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Comparison Values ---")

	dimensionComp := operations.DimensionBrowserName
	value := "Chrome"
	comparisonResponse, err := client.Metrics.ListComparisonValues(
		ctx,
		operations.ListComparisonValuesTimespanTwentyFourhours.ToPointer(),
		nil,
		&dimensionComp,
		&value,
	)
	if err != nil {
		log.Printf("Error listing comparison values: %v", err)
		return
	}
	if comparisonResponse.Object == nil {
		return
	}

	fmt.Printf("Comparison values retrieved successfully!\n")
	fmt.Printf(dataPoints, len(comparisonResponse.Object.Data))
}

func runDimensionsAnalytics(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Dimensions ===")

	listDimensions(ctx, client)
	listFilterValuesForDimension(ctx, client)
}

func listDimensions(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Available Dimensions ---")

	dimensionsResponse, err := client.Dimensions.List(ctx)
	if err != nil {
		log.Printf("Error listing dimensions: %v", err)
		return
	}
	if dimensionsResponse.Object == nil {
		return
	}

	fmt.Printf("Found %d dimensions:\n", len(dimensionsResponse.Object.Data))
	for i, dimension := range dimensionsResponse.Object.Data {
		fmt.Printf("  %d. %s\n", i+1, dimension)
	}
}

func listFilterValuesForDimension(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n--- Filter Values for Browser Name ---")

	filterValuesResponse, err := client.Dimensions.ListFilterValues(
		ctx,
		operations.DimensionsIDBrowserName,
		operations.ListFilterValuesForDimensionTimespanTwentyFourhours.ToPointer(),
		nil,
	)
	if err != nil {
		log.Printf("Error listing filter values: %v", err)
		return
	}
	if filterValuesResponse.Object == nil {
		return
	}

	fmt.Printf("Found %d filter values for browser name:\n", len(filterValuesResponse.Object.Data))
	for i, value := range filterValuesResponse.Object.Data {
		fmt.Printf("  %d. Value: %s, Count: %d\n",
			i+1, getStringValue(value.Value), getInt64Value(value.Count))
	}
}

func runAdvancedAnalyticsQueries(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Advanced Analytics Queries ===")

	timespans := []operations.ListVideoViewsTimespan{
		operations.ListVideoViewsTimespanSixtyminutes,
		operations.ListVideoViewsTimespanTwentyFourhours,
		operations.ListVideoViewsTimespanSevendays,
		operations.ListVideoViewsTimespanThirtydays,
	}

	for _, timespan := range timespans {
		printViewsForTimespan(ctx, client, timespan)
	}
}

func printViewsForTimespan(ctx context.Context, client *fastpixgo.Fastpixgo, timespan operations.ListVideoViewsTimespan) {
	fmt.Printf("\n--- Views for %s ---\n", timespan)

	response, err := client.Views.ListVideoViews(ctx, operations.ListVideoViewsRequest{
		Timespan: timespan.ToPointer(),
		Limit:    fastpixgo.Int64(5),
	})
	if err != nil {
		log.Printf("Error getting views for %s: %v", timespan, err)
		return
	}
	if response.Object == nil {
		return
	}

	fmt.Printf("Found %d views\n", len(response.Object.Data))
}

func runAnalyticsDashboardSummary(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Analytics Dashboard Summary ===")
	fmt.Println("Generating analytics summary...")

	overall, err := client.Metrics.ListOverallValues(
		ctx,
		operations.ListOverallValuesMetricIDViews,
		fastpixgo.String("count"),
		operations.ListOverallValuesTimespanTwentyFourhours.ToPointer(),
		nil,
	)
	if err == nil && overall.Object != nil && overall.Object.Data != nil {
		fmt.Printf("Total views (24h): %d\n", getOptInt64(overall.Object.Data.TotalViews))
	}

	topContent, err := client.Views.ListByTopContent(
		ctx,
		operations.ListByTopContentTimespanTwentyFourhours.ToPointer(),
		nil,
		fastpixgo.Int64(5),
	)
	if err == nil && topContent.Object != nil {
		fmt.Printf("Top 5 content items: %d items\n", len(topContent.Object.Data))
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

func getOptString(n optionalnullable.OptionalNullable[string]) string {
	v, _ := n.GetOrZero()
	return v
}

func getOptInt64(n optionalnullable.OptionalNullable[int64]) int64 {
	v, _ := n.GetOrZero()
	return v
}

func getOptFloat(n optionalnullable.OptionalNullable[float64]) float64 {
	v, _ := n.GetOrZero()
	return v
}

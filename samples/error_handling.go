//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/retry"
)

const invalidMediaID = "invalid-media-id"

func main() {
	ctx := context.Background()

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	testBasicErrorHandling(ctx, client)
	testSpecificErrorTypeHandling(ctx, client)
	testAuthErrorHandling(ctx)
	testTimeoutErrorHandling(ctx)
	testRetryErrorHandling(ctx)
	testComprehensiveErrorHandling(ctx, client)
	testErrorRecoveryStrategies(ctx, client)
	testErrorLoggingAndMonitoring(ctx, client)
	testCustomErrorHandling(ctx, client)

	fmt.Println("\n=== Error Handling Demo Complete ===")
}

func testBasicErrorHandling(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("=== Basic Error Handling ===")
	fmt.Println("Testing error handling with non-existent media summary...")

	_, err := client.ManageVideos.GetMediaSummary(ctx, "non-existent-media-id")
	handleError("GetMediaSummary", err)
}

func testSpecificErrorTypeHandling(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Specific Error Type Handling ===")
	fmt.Println("Testing with invalid media creation request...")

	invalidRequest := components.CreateMediaRequest{
		Inputs:       []components.Input{},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic.ToPointer(),
	}

	_, err := client.InputVideo.Create(ctx, invalidRequest)
	handleSpecificErrors("Create media", err)
}

func testAuthErrorHandling(ctx context.Context) {
	fmt.Println("\n=== Authentication Error Handling ===")

	invalidClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer("invalid-username"),
			Password: fastpixgo.Pointer("invalid-password"),
		}),
		fastpixgo.WithTimeout(10*time.Second),
	)

	fmt.Println("Testing with invalid credentials...")
	_, err := invalidClient.ManageVideos.List(ctx, nil, nil, nil)
	handleSpecificErrors("List with invalid auth", err)
}

func testTimeoutErrorHandling(ctx context.Context) {
	fmt.Println("\n=== Network and Timeout Error Handling ===")

	shortTimeoutClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(1*time.Millisecond),
	)

	fmt.Println("Testing with short timeout...")
	_, err := shortTimeoutClient.ManageVideos.List(ctx, nil, nil, nil)
	handleError("List with short timeout", err)
}

func testRetryErrorHandling(ctx context.Context) {
	fmt.Println("\n=== Retry Configuration and Error Handling ===")

	retryClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithRetryConfig(retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     5,
				Exponent:        2,
				MaxElapsedTime:  10,
			},
			RetryConnectionErrors: true,
		}),
	)

	fmt.Println("Testing with retry configuration...")
	_, err := retryClient.ManageVideos.List(ctx, nil, nil, nil)
	handleError("List with retry", err)
}

func testComprehensiveErrorHandling(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Comprehensive Error Handling ===")

	testOperations := []struct {
		name string
		fn   func() error
	}{
		{
			name: "GetMediaSummary with invalid ID",
			fn: func() error {
				_, err := client.ManageVideos.GetMediaSummary(ctx, "invalid-id")
				return err
			},
		},
		{
			name: "CreatePlaylist with invalid data",
			fn: func() error {
				_, err := client.Playlists.Create(ctx, components.CreatePlaylistRequest{})
				return err
			},
		},
		{
			name: "GetSigningKey with invalid ID",
			fn: func() error {
				_, err := client.SigningKeys.GetByID(ctx, "invalid-key-id")
				return err
			},
		},
		{
			name: "CreateLiveStream with invalid data",
			fn: func() error {
				_, err := client.StartLiveStream.Create(ctx, components.CreateLiveStreamRequest{})
				return err
			},
		},
	}

	for _, test := range testOperations {
		fmt.Printf("\nTesting: %s\n", test.name)
		handleSpecificErrors(test.name, test.fn())
	}
}

func testErrorRecoveryStrategies(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Error Recovery Strategies ===")
	fmt.Println("Demonstrating error recovery strategies...")

	fmt.Println("\nStrategy 1: Retry with exponential backoff")
	handleError("Retry with backoff", retryWithBackoff(ctx, client, 3))

	fmt.Println("\nStrategy 2: Fallback to alternative operation")
	handleError("Fallback operation", fallbackOperation(ctx, client))

	fmt.Println("\nStrategy 3: Graceful degradation")
	handleError("Graceful degradation", gracefulDegradation(ctx, client))
}

func testErrorLoggingAndMonitoring(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Error Logging and Monitoring ===")
	fmt.Println("Demonstrating error logging...")

	_, err := client.ManageVideos.GetMediaSummary(ctx, invalidMediaID)
	if err != nil {
		logError("GetMediaSummary", err, map[string]interface{}{
			"media_id": invalidMediaID,
			"user_id":  "demo-user",
			"action":   "get_media_summary",
		})
	}
}

func testCustomErrorHandling(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Custom Error Handling ===")
	fmt.Println("Demonstrating custom error handling wrapper...")

	result, err := safeExecute(func() (interface{}, error) {
		return client.ManageVideos.List(ctx, nil, nil, nil)
	})

	if err != nil {
		fmt.Printf("Custom error handling result: %v\n", err)
	} else {
		fmt.Printf("Custom error handling success: %+v\n", result)
	}
}

func handleError(operation string, err error) {
	if err != nil {
		fmt.Printf("[FAIL] %s failed: %v\n", operation, err)
	} else {
		fmt.Printf("[OK] %s succeeded\n", operation)
	}
}

func handleSpecificErrors(operation string, err error) {
	if err == nil {
		fmt.Printf("[OK] %s succeeded\n", operation)
		return
	}

	fmt.Printf("[FAIL] %s failed: ", operation)

	if msg, ok := matchAPIError(err); ok {
		fmt.Println(msg)
		return
	}

	fmt.Printf("Unknown Error: %v\n", err)
}

func matchAPIError(err error) (string, bool) {
	var apiErr *apierrors.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.StatusCode {
		case 400:
			return fmt.Sprintf("Bad Request Error (400): %s", apiErr.Error()), true
		case 401:
			return fmt.Sprintf("Unauthorized Error (401): %s", apiErr.Error()), true
		case 403:
			return fmt.Sprintf("Forbidden Error (403): %s", apiErr.Error()), true
		case 404:
			return fmt.Sprintf("Not Found Error (404): %s", apiErr.Error()), true
		case 422:
			return fmt.Sprintf("Validation Error (422): %s", apiErr.Error()), true
		default:
			return fmt.Sprintf("API Error (%d): %s", apiErr.StatusCode, apiErr.Error()), true
		}
	}

	return "", false
}

func retryWithBackoff(ctx context.Context, client *fastpixgo.Fastpixgo, maxRetries int) error {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		fmt.Printf("  Attempt %d/%d\n", i+1, maxRetries)

		_, err := client.ManageVideos.List(ctx, nil, nil, nil)
		if err == nil {
			fmt.Printf("  [OK] Success on attempt %d\n", i+1)
			return nil
		}

		lastErr = err

		if i < maxRetries-1 {
			backoffDuration := time.Duration(1<<uint(i)) * time.Second
			fmt.Printf("  Waiting %v before retry...\n", backoffDuration)
			time.Sleep(backoffDuration)
		}
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

func fallbackOperation(ctx context.Context, client *fastpixgo.Fastpixgo) error {
	_, err := client.ManageVideos.GetMediaSummary(ctx, invalidMediaID)
	if err == nil {
		fmt.Println("  [OK] Primary operation succeeded")
		return nil
	}

	fmt.Println("  Primary operation failed, trying fallback...")

	_, fallbackErr := client.ManageVideos.List(ctx, nil, nil, nil)
	if fallbackErr == nil {
		fmt.Println("  [OK] Fallback operation succeeded")
		return nil
	}

	return fmt.Errorf("both primary and fallback operations failed: primary=%v, fallback=%v", err, fallbackErr)
}

func gracefulDegradation(ctx context.Context, client *fastpixgo.Fastpixgo) error {
	_, err := client.ManageVideos.GetMediaSummary(ctx, invalidMediaID)
	if err == nil {
		fmt.Println("  [OK] Full functionality available")
		return nil
	}

	fmt.Println("  Full functionality unavailable, using degraded mode...")

	_, basicErr := client.ManageVideos.List(ctx, nil, nil, nil)
	if basicErr == nil {
		fmt.Println("  [OK] Degraded mode working")
		return nil
	}

	return fmt.Errorf("even degraded mode failed: %v", basicErr)
}

func logError(operation string, err error, context map[string]interface{}) {
	fmt.Printf("  Logging error for %s\n", operation)
	fmt.Printf("  Context: %+v\n", context)
	fmt.Printf("  Error: %v\n", err)
}

func safeExecute(fn func() (interface{}, error)) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Panic recovered: %v\n", r)
		}
	}()

	return fn()
}

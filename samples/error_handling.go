package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	fastpixgo "github.com/fastpix/fastpix-go"
	"github.com/fastpix/fastpix-go/models/apierrors"
	"github.com/fastpix/fastpix-go/models/components"
	"github.com/fastpix/fastpix-go/models/operations"
)

func main() {
	ctx := context.Background()

	// Initialize SDK
	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	// 1. Basic Error Handling
	fmt.Println("=== Basic Error Handling ===")
	
	// Try to get a non-existent media
	fmt.Println("Testing error handling with non-existent media...")
	_, err := client.ManageVideos.GetMedia(ctx, "non-existent-media-id")
	handleError("GetMedia", err)

	// 2. Specific Error Type Handling
	fmt.Println("\n=== Specific Error Type Handling ===")
	
	// Test with invalid media creation request
	fmt.Println("Testing with invalid media creation request...")
	invalidRequest := components.CreateMediaRequest{
		Inputs: []components.Input{}, // Empty inputs should cause validation error
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
	}
	
	_, err = client.InputVideo.CreateMedia(ctx, invalidRequest)
	handleSpecificErrors("CreateMedia", err)

	// 3. Authentication Error Handling
	fmt.Println("\n=== Authentication Error Handling ===")
	
	// Create client with invalid credentials
	invalidClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer("invalid-username"),
			Password: fastpixgo.Pointer("invalid-password"),
		}),
		fastpixgo.WithTimeout(10*time.Second),
	)
	
	fmt.Println("Testing with invalid credentials...")
	_, err = invalidClient.ManageVideos.ListMedia(ctx, nil, nil, nil)
	handleSpecificErrors("ListMedia with invalid auth", err)

	// 4. Network and Timeout Error Handling
	fmt.Println("\n=== Network and Timeout Error Handling ===")
	
	// Create client with very short timeout
	shortTimeoutClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(1*time.Millisecond), // Very short timeout
	)
	
	fmt.Println("Testing with short timeout...")
	_, err = shortTimeoutClient.ManageVideos.ListMedia(ctx, nil, nil, nil)
	handleError("ListMedia with short timeout", err)

	// 5. Retry Configuration and Error Handling
	fmt.Println("\n=== Retry Configuration and Error Handling ===")
	
	// Create client with retry configuration
	retryClient := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithRetryConfig(fastpixgo.RetryConfig{
			Strategy: "backoff",
			Backoff: &fastpixgo.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     5,
				Exponent:        2,
				MaxElapsedTime:  10,
			},
			RetryConnectionErrors: true,
		}),
	)
	
	fmt.Println("Testing with retry configuration...")
	_, err = retryClient.ManageVideos.ListMedia(ctx, nil, nil, nil)
	handleError("ListMedia with retry", err)

	// 6. Comprehensive Error Handling for Different Operations
	fmt.Println("\n=== Comprehensive Error Handling ===")
	
	// Test various operations with error handling
	testOperations := []struct {
		name string
		fn   func() error
	}{
		{
			name: "GetMedia with invalid ID",
			fn: func() error {
				_, err := client.ManageVideos.GetMedia(ctx, "invalid-id")
				return err
			},
		},
		{
			name: "CreatePlaylist with invalid data",
			fn: func() error {
				_, err := client.Playlist.CreateAPlaylist(ctx, components.CreatePlaylistRequest{})
				return err
			},
		},
		{
			name: "GetSigningKey with invalid ID",
			fn: func() error {
				_, err := client.SigningKeys.GetSigningKeyByID(ctx, "invalid-key-id")
				return err
			},
		},
		{
			name: "CreateLiveStream with invalid data",
			fn: func() error {
				_, err := client.StartLiveStream.CreateNewStream(ctx, components.CreateLiveStreamRequest{})
				return err
			},
		},
	}

	for _, test := range testOperations {
		fmt.Printf("\nTesting: %s\n", test.name)
		err := test.fn()
		handleSpecificErrors(test.name, err)
	}

	// 7. Error Recovery Strategies
	fmt.Println("\n=== Error Recovery Strategies ===")
	
	// Demonstrate fallback strategies
	fmt.Println("Demonstrating error recovery strategies...")
	
	// Strategy 1: Retry with exponential backoff
	fmt.Println("\nStrategy 1: Retry with exponential backoff")
	err = retryWithBackoff(ctx, client, 3)
	handleError("Retry with backoff", err)
	
	// Strategy 2: Fallback to alternative operation
	fmt.Println("\nStrategy 2: Fallback to alternative operation")
	err = fallbackOperation(ctx, client)
	handleError("Fallback operation", err)
	
	// Strategy 3: Graceful degradation
	fmt.Println("\nStrategy 3: Graceful degradation")
	err = gracefulDegradation(ctx, client)
	handleError("Graceful degradation", err)

	// 8. Error Logging and Monitoring
	fmt.Println("\n=== Error Logging and Monitoring ===")
	
	// Demonstrate proper error logging
	fmt.Println("Demonstrating error logging...")
	
	// Log errors with context
	_, err = client.ManageVideos.GetMedia(ctx, "invalid-media-id")
	if err != nil {
		logError("GetMedia", err, map[string]interface{}{
			"media_id": "invalid-media-id",
			"user_id":  "demo-user",
			"action":   "get_media",
		})
	}

	// 9. Custom Error Handling
	fmt.Println("\n=== Custom Error Handling ===")
	
	// Demonstrate custom error handling wrapper
	fmt.Println("Demonstrating custom error handling wrapper...")
	
	result, err := safeExecute(func() (interface{}, error) {
		return client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	})
	
	if err != nil {
		fmt.Printf("Custom error handling result: %v\n", err)
	} else {
		fmt.Printf("Custom error handling success: %+v\n", result)
	}

	fmt.Println("\n=== Error Handling Demo Complete ===")
}

// handleError demonstrates basic error handling
func handleError(operation string, err error) {
	if err != nil {
		fmt.Printf("❌ %s failed: %v\n", operation, err)
	} else {
		fmt.Printf("✅ %s succeeded\n", operation)
	}
}

// handleSpecificErrors demonstrates specific error type handling
func handleSpecificErrors(operation string, err error) {
	if err == nil {
		fmt.Printf("✅ %s succeeded\n", operation)
		return
	}

	fmt.Printf("❌ %s failed: ", operation)

	// Check for specific error types
	var badRequestErr *apierrors.BadRequestError
	if errors.As(err, &badRequestErr) {
		fmt.Printf("Bad Request Error (400): %s\n", badRequestErr.Error())
		return
	}

	var unauthorizedErr *apierrors.UnauthorizedError
	if errors.As(err, &unauthorizedErr) {
		fmt.Printf("Unauthorized Error (401): %s\n", unauthorizedErr.Error())
		return
	}

	var forbiddenErr *apierrors.ForbiddenError
	if errors.As(err, &forbiddenErr) {
		fmt.Printf("Forbidden Error (403): %s\n", forbiddenErr.Error())
		return
	}

	var notFoundErr *apierrors.NotFoundError
	if errors.As(err, &notFoundErr) {
		fmt.Printf("Not Found Error (404): %s\n", notFoundErr.Error())
		return
	}

	var validationErr *apierrors.ValidationErrorResponse
	if errors.As(err, &validationErr) {
		fmt.Printf("Validation Error (422): %s\n", validationErr.Error())
		return
	}

	var apiErr *apierrors.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API Error (%d): %s\n", apiErr.StatusCode, apiErr.Error())
		return
	}

	// Generic error
	fmt.Printf("Unknown Error: %v\n", err)
}

// retryWithBackoff demonstrates retry with exponential backoff
func retryWithBackoff(ctx context.Context, client *fastpixgo.Fastpixgo, maxRetries int) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		fmt.Printf("  Attempt %d/%d\n", i+1, maxRetries)
		
		_, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		if err == nil {
			fmt.Printf("  ✅ Success on attempt %d\n", i+1)
			return nil
		}
		
		lastErr = err
		
		if i < maxRetries-1 {
			backoffDuration := time.Duration(1<<uint(i)) * time.Second
			fmt.Printf("  ⏳ Waiting %v before retry...\n", backoffDuration)
			time.Sleep(backoffDuration)
		}
	}
	
	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

// fallbackOperation demonstrates fallback strategy
func fallbackOperation(ctx context.Context, client *fastpixgo.Fastpixgo) error {
	// Try primary operation
	_, err := client.ManageVideos.GetMedia(ctx, "invalid-media-id")
	if err == nil {
		fmt.Println("  ✅ Primary operation succeeded")
		return nil
	}
	
	fmt.Println("  ⚠️ Primary operation failed, trying fallback...")
	
	// Try fallback operation
	_, fallbackErr := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if fallbackErr == nil {
		fmt.Println("  ✅ Fallback operation succeeded")
		return nil
	}
	
	return fmt.Errorf("both primary and fallback operations failed: primary=%v, fallback=%v", err, fallbackErr)
}

// gracefulDegradation demonstrates graceful degradation
func gracefulDegradation(ctx context.Context, client *fastpixgo.Fastpixgo) error {
	// Try to get detailed information
	_, err := client.ManageVideos.GetMedia(ctx, "invalid-media-id")
	if err == nil {
		fmt.Println("  ✅ Full functionality available")
		return nil
	}
	
	fmt.Println("  ⚠️ Full functionality unavailable, using degraded mode...")
	
	// Fall back to basic functionality
	_, basicErr := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if basicErr == nil {
		fmt.Println("  ✅ Degraded mode working")
		return nil
	}
	
	return fmt.Errorf("even degraded mode failed: %v", basicErr)
}

// logError demonstrates proper error logging
func logError(operation string, err error, context map[string]interface{}) {
	fmt.Printf("  📝 Logging error for %s\n", operation)
	fmt.Printf("  Context: %+v\n", context)
	fmt.Printf("  Error: %v\n", err)
	
	// In a real application, you would send this to your logging system
	// log.WithFields(log.Fields(context)).WithError(err).Errorf("Operation %s failed", operation)
}

// safeExecute demonstrates custom error handling wrapper
func safeExecute(fn func() (interface{}, error)) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  🚨 Panic recovered: %v\n", r)
		}
	}()
	
	return fn()
}

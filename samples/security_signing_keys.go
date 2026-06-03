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
)

func main() {
	ctx := context.Background()

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	createSigningKey(ctx, client)
	listSigningKeys(ctx, client)
	manageFirstSigningKey(ctx, client)
	signingKeyManagementWorkflow(ctx, client)
	securityBestPractices(ctx, client)
	securityErrorHandling(ctx, client)
	paginationExample(ctx, client)

	fmt.Println("\n=== Security Demo Complete ===")
	fmt.Println("Remember to implement proper key management in production!")
}

func createSigningKey(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("=== Creating Signing Key ===")

	createKeyResponse, err := client.SigningKeys.CreateSigningKey(ctx)
	if err != nil {
		log.Printf("Error creating signing key: %v", err)
		return
	}

	if createKeyResponse.CreateResponse == nil || createKeyResponse.CreateResponse.Data == nil {
		fmt.Println("Signing key created but no data returned")
		return
	}

	data := createKeyResponse.CreateResponse.Data
	if data.ID != nil {
		fmt.Printf("Signing key created successfully! ID: %s\n", *data.ID)
	}
	fmt.Printf("Private Key: %s\n", getStringValue(data.PrivateKey))
	fmt.Printf("Public Key: %s\n", getStringValue(data.PublicKey))
}

func listSigningKeys(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Listing All Signing Keys ===")

	limit := 10.0
	offset := 0.0

	keysResponse, err := client.SigningKeys.ListSigningKeys(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing signing keys: %v", err)
		return
	}

	fmt.Printf("Found %d signing keys:\n", len(keysResponse.GetAllSigningKeyResponse.Data))
	for i, key := range keysResponse.GetAllSigningKeyResponse.Data {
		fmt.Printf("  %d. ID: %s, Created: %s\n",
			i+1, *key.ID, getStringValue(key.CreatedAt))
	}
}

func manageFirstSigningKey(ctx context.Context, client *fastpixgo.FastPix) {
	limit := 10.0
	offset := 0.0

	keysResponse, err := client.SigningKeys.ListSigningKeys(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing signing keys: %v", err)
		return
	}

	if keysResponse.GetAllSigningKeyResponse == nil || len(keysResponse.GetAllSigningKeyResponse.Data) == 0 {
		return
	}

	keyID := *keysResponse.GetAllSigningKeyResponse.Data[0].ID
	getSigningKeyDetails(ctx, client, keyID)
	deleteSigningKey(ctx, client, keyID)
}

func getSigningKeyDetails(ctx context.Context, client *fastpixgo.FastPix, keyID string) {
	fmt.Printf("\n=== Getting Signing Key Details: %s ===\n", keyID)

	keyDetailsResponse, err := client.SigningKeys.GetSigningKeyByID(ctx, keyID)
	if err != nil {
		log.Printf("Error getting signing key details: %v", err)
		return
	}

	key := keyDetailsResponse.GetPublicPemUsingSigningKeyIDResponseDTO.Data
	fmt.Printf("Key ID: %s\n", *key.ID)
	fmt.Printf("Public Key: %s\n", getStringValue(key.PublicKey))
	fmt.Printf("Created: %s\n", getStringValue(key.CreatedAt))
}

func deleteSigningKey(ctx context.Context, client *fastpixgo.FastPix, keyID string) {
	fmt.Printf("\n=== Deleting Signing Key: %s ===\n", keyID)

	deleteResponse, err := client.SigningKeys.DeleteSigningKey(ctx, keyID)
	if err != nil {
		log.Printf("Error deleting signing key: %v", err)
		return
	}

	fmt.Println("Signing key deleted successfully!")
	fmt.Printf("Response: %+v\n", deleteResponse)
}

func signingKeyManagementWorkflow(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Signing Key Management Workflow ===")
	fmt.Println("Creating multiple signing keys for rotation...")

	createdKeys := createMultipleSigningKeys(ctx, client, 3)

	fmt.Println("\nListing all signing keys after creation...")
	allKeysResponse, err := client.SigningKeys.ListSigningKeys(ctx, nil, nil)
	if err != nil {
		log.Printf("Error listing all keys: %v", err)
	} else {
		fmt.Printf("Total signing keys: %d\n", len(allKeysResponse.GetAllSigningKeyResponse.Data))
	}

	rotateSigningKeys(ctx, client, createdKeys)
}

func createMultipleSigningKeys(ctx context.Context, client *fastpixgo.FastPix, count int) []string {
	var createdKeys []string

	for i := 0; i < count; i++ {
		createResponse, err := client.SigningKeys.CreateSigningKey(ctx)
		if err != nil {
			log.Printf("Error creating signing key %d: %v", i+1, err)
			continue
		}

		if createResponse.CreateResponse != nil &&
			createResponse.CreateResponse.Data != nil &&
			createResponse.CreateResponse.Data.ID != nil {
			keyID := *createResponse.CreateResponse.Data.ID
			createdKeys = append(createdKeys, keyID)
			fmt.Printf("Created signing key %d: %s\n", i+1, keyID)
		} else {
			fmt.Printf("Created signing key %d but no ID returned\n", i+1)
		}
	}

	return createdKeys
}

func rotateSigningKeys(ctx context.Context, client *fastpixgo.FastPix, createdKeys []string) {
	fmt.Println("\n=== Key Rotation Example ===")

	if len(createdKeys) == 0 {
		return
	}

	fmt.Println("Performing key rotation...")

	for i, keyID := range createdKeys[:len(createdKeys)-1] {
		fmt.Printf("Deleting old key %d: %s\n", i+1, keyID)
		_, err := client.SigningKeys.DeleteSigningKey(ctx, keyID)
		if err != nil {
			log.Printf("Error deleting key %s: %v", keyID, err)
		} else {
			fmt.Printf("Successfully deleted key %s\n", keyID)
		}
	}

	activeKeyID := createdKeys[len(createdKeys)-1]
	fmt.Printf("Active key: %s\n", activeKeyID)

	activeKeyResponse, err := client.SigningKeys.GetSigningKeyByID(ctx, activeKeyID)
	if err != nil {
		log.Printf("Error getting active key details: %v", err)
		return
	}

	fmt.Printf("Active key public key: %s\n",
		getStringValue(activeKeyResponse.GetPublicPemUsingSigningKeyIDResponseDTO.Data.PublicKey))
}

func securityBestPractices(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Security Best Practices ===")
	fmt.Println("Creating a new signing key for security demonstration...")

	newKeyResponse, err := client.SigningKeys.CreateSigningKey(ctx)
	if err != nil {
		log.Printf("Error creating new signing key: %v", err)
		return
	}

	if newKeyResponse.CreateResponse == nil ||
		newKeyResponse.CreateResponse.Data == nil ||
		newKeyResponse.CreateResponse.Data.ID == nil {
		fmt.Println("New signing key created but no ID returned")
		return
	}

	newKeyID := *newKeyResponse.CreateResponse.Data.ID
	fmt.Printf("New signing key created: %s\n", newKeyID)

	printSecurityRecommendations()
	cleanupDemonstrationKey(ctx, client, newKeyID)
}

func printSecurityRecommendations() {
	fmt.Println("\nSecurity recommendations:")
	fmt.Println("1. Store private keys securely (not in code)")
	fmt.Println("2. Rotate keys regularly")
	fmt.Println("3. Use environment variables for sensitive data")
	fmt.Println("4. Implement proper access controls")
	fmt.Println("5. Monitor key usage and access")
}

func cleanupDemonstrationKey(ctx context.Context, client *fastpixgo.FastPix, keyID string) {
	fmt.Printf("\nCleaning up demonstration key: %s\n", keyID)

	_, err := client.SigningKeys.DeleteSigningKey(ctx, keyID)
	if err != nil {
		log.Printf("Error cleaning up key: %v", err)
		return
	}

	fmt.Println("Demonstration key cleaned up successfully")
}

func securityErrorHandling(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Security Error Handling ===")

	fakeKeyID := "non-existent-key-id"

	_, err := client.SigningKeys.GetSigningKeyByID(ctx, fakeKeyID)
	if err != nil {
		fmt.Printf("Expected error for non-existent key: %v\n", err)
	} else {
		fmt.Println("Unexpected success for non-existent key")
	}

	_, err = client.SigningKeys.DeleteSigningKey(ctx, fakeKeyID)
	if err != nil {
		fmt.Printf("Expected error deleting non-existent key: %v\n", err)
	} else {
		fmt.Println("Unexpected success deleting non-existent key")
	}
}

func paginationExample(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Pagination Example ===")

	pageSize := 5.0
	pageOffset := 0.0

	fmt.Printf("Fetching signing keys with pagination (limit: %.0f, offset: %.0f)\n", pageSize, pageOffset)

	paginatedResponse, err := client.SigningKeys.ListSigningKeys(ctx, &pageSize, &pageOffset)
	if err != nil {
		log.Printf("Error with paginated request: %v", err)
		return
	}

	fmt.Printf("Page 1: Found %d keys\n", len(paginatedResponse.GetAllSigningKeyResponse.Data))

	pageOffset = 5.0
	fmt.Printf("Fetching next page (limit: %.0f, offset: %.0f)\n", pageSize, pageOffset)

	nextPageResponse, err := client.SigningKeys.ListSigningKeys(ctx, &pageSize, &pageOffset)
	if err != nil {
		log.Printf("Error with next page request: %v", err)
		return
	}

	fmt.Printf("Page 2: Found %d keys\n", len(nextPageResponse.GetAllSigningKeyResponse.Data))
}

func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
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

	// Initialize SDK
	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

	// 1. Create Signing Key
	fmt.Println("=== Creating Signing Key ===")
	createKeyResponse, err := client.SigningKeys.CreateSigningKey(ctx)
	if err != nil {
		log.Printf("Error creating signing key: %v", err)
	} else {
		if createKeyResponse.CreateResponse != nil && createKeyResponse.CreateResponse.Data != nil {
			keyID := createKeyResponse.CreateResponse.Data.ID
			if keyID != nil {
				fmt.Printf("Signing key created successfully! ID: %s\n", *keyID)
			}
			fmt.Printf("Private Key: %s\n", getStringValue(createKeyResponse.CreateResponse.Data.PrivateKey))
			fmt.Printf("Public Key: %s\n", getStringValue(createKeyResponse.CreateResponse.Data.PublicKey))
		} else {
			fmt.Println("Signing key created but no data returned")
		}
	}

	// 2. List All Signing Keys
	fmt.Println("\n=== Listing All Signing Keys ===")
	limit := 10.0
	offset := 0.0

	keysResponse, err := client.SigningKeys.ListSigningKeys(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing signing keys: %v", err)
	} else {
		fmt.Printf("Found %d signing keys:\n", len(keysResponse.GetAllSigningKeyResponse.Data))
		for i, key := range keysResponse.GetAllSigningKeyResponse.Data {
			fmt.Printf("  %d. ID: %s, Created: %s\n",
				i+1, *key.ID, getStringValue(key.CreatedAt))
		}
	}

	// 3. Get Specific Signing Key Details
	if keysResponse.GetAllSigningKeyResponse != nil && len(keysResponse.GetAllSigningKeyResponse.Data) > 0 {
		keyID := keysResponse.GetAllSigningKeyResponse.Data[0].ID
		fmt.Printf("\n=== Getting Signing Key Details: %s ===\n", *keyID)

		keyDetailsResponse, err := client.SigningKeys.GetSigningKeyByID(ctx, *keyID)
		if err != nil {
			log.Printf("Error getting signing key details: %v", err)
		} else {
			key := keyDetailsResponse.GetPublicPemUsingSigningKeyIDResponseDTO.Data
			fmt.Printf("Key ID: %s\n", *key.ID)
			fmt.Printf("Public Key: %s\n", getStringValue(key.PublicKey))
			fmt.Printf("Created: %s\n", getStringValue(key.CreatedAt))
		}

		// 4. Delete Signing Key
		fmt.Printf("\n=== Deleting Signing Key: %s ===\n", *keyID)
		deleteResponse, err := client.SigningKeys.DeleteSigningKey(ctx, *keyID)
		if err != nil {
			log.Printf("Error deleting signing key: %v", err)
		} else {
			fmt.Println("Signing key deleted successfully!")
			fmt.Printf("Response: %+v\n", deleteResponse)
		}
	}

	// 5. Signing Key Management Workflow
	fmt.Println("\n=== Signing Key Management Workflow ===")

	// Create multiple signing keys for rotation
	fmt.Println("Creating multiple signing keys for rotation...")
	var createdKeys []string

	for i := 0; i < 3; i++ {
		createResponse, err := client.SigningKeys.CreateSigningKey(ctx)
		if err != nil {
			log.Printf("Error creating signing key %d: %v", i+1, err)
			continue
		}

		if createResponse.CreateResponse != nil && createResponse.CreateResponse.Data != nil && createResponse.CreateResponse.Data.ID != nil {
			keyID := createResponse.CreateResponse.Data.ID
			createdKeys = append(createdKeys, *keyID)
			fmt.Printf("Created signing key %d: %s\n", i+1, *keyID)
		} else {
			fmt.Printf("Created signing key %d but no ID returned\n", i+1)
		}
	}

	// List all keys after creation
	fmt.Println("\nListing all signing keys after creation...")
	allKeysResponse, err := client.SigningKeys.ListSigningKeys(ctx, nil, nil)
	if err != nil {
		log.Printf("Error listing all keys: %v", err)
	} else {
		fmt.Printf("Total signing keys: %d\n", len(allKeysResponse.GetAllSigningKeyResponse.Data))
	}

	// 6. Key Rotation Example
	fmt.Println("\n=== Key Rotation Example ===")

	if len(createdKeys) > 0 {
		// Simulate key rotation by deleting old keys
		fmt.Println("Performing key rotation...")

		// Keep the newest key, delete others
		for i, keyID := range createdKeys[:len(createdKeys)-1] {
			fmt.Printf("Deleting old key %d: %s\n", i+1, keyID)
			_, err := client.SigningKeys.DeleteSigningKey(ctx, keyID)
			if err != nil {
				log.Printf("Error deleting key %s: %v", keyID, err)
			} else {
				fmt.Printf("Successfully deleted key %s\n", keyID)
			}
		}

		// Keep the last key as active
		activeKeyID := createdKeys[len(createdKeys)-1]
		fmt.Printf("Active key: %s\n", activeKeyID)

		// Get details of active key
		activeKeyResponse, err := client.SigningKeys.GetSigningKeyByID(ctx, activeKeyID)
		if err != nil {
			log.Printf("Error getting active key details: %v", err)
		} else {
			fmt.Printf("Active key public key: %s\n",
				getStringValue(activeKeyResponse.GetPublicPemUsingSigningKeyIDResponseDTO.Data.PublicKey))
		}
	}

	// 7. Security Best Practices
	fmt.Println("\n=== Security Best Practices ===")

	// Create a new key for demonstration
	fmt.Println("Creating a new signing key for security demonstration...")
	newKeyResponse, err := client.SigningKeys.CreateSigningKey(ctx)
	if err != nil {
		log.Printf("Error creating new signing key: %v", err)
	} else {
		if newKeyResponse.CreateResponse != nil && newKeyResponse.CreateResponse.Data != nil && newKeyResponse.CreateResponse.Data.ID != nil {
			newKeyID := newKeyResponse.CreateResponse.Data.ID
			fmt.Printf("New signing key created: %s\n", *newKeyID)
		} else {
			fmt.Println("New signing key created but no ID returned")
		}

		// Demonstrate secure key handling
		fmt.Println("\nSecurity recommendations:")
		fmt.Println("1. Store private keys securely (not in code)")
		fmt.Println("2. Rotate keys regularly")
		fmt.Println("3. Use environment variables for sensitive data")
		fmt.Println("4. Implement proper access controls")
		fmt.Println("5. Monitor key usage and access")

		// Clean up the demonstration key
		fmt.Printf("\nCleaning up demonstration key: %s\n", *newKeyID)
		_, err = client.SigningKeys.DeleteSigningKey(ctx, *newKeyID)
		if err != nil {
			log.Printf("Error cleaning up key: %v", err)
		} else {
			fmt.Println("Demonstration key cleaned up successfully")
		}
	}

	// 8. Error Handling for Security Operations
	fmt.Println("\n=== Security Error Handling ===")

	// Try to get a non-existent key
	fakeKeyID := "non-existent-key-id"
	_, err = client.SigningKeys.GetSigningKeyByID(ctx, fakeKeyID)
	if err != nil {
		fmt.Printf("Expected error for non-existent key: %v\n", err)
	} else {
		fmt.Println("Unexpected success for non-existent key")
	}

	// Try to delete a non-existent key
	_, err = client.SigningKeys.DeleteSigningKey(ctx, fakeKeyID)
	if err != nil {
		fmt.Printf("Expected error deleting non-existent key: %v\n", err)
	} else {
		fmt.Println("Unexpected success deleting non-existent key")
	}

	// 9. Pagination Example
	fmt.Println("\n=== Pagination Example ===")

	// Demonstrate pagination for signing keys
	pageSize := 5.0
	pageOffset := 0.0

	fmt.Printf("Fetching signing keys with pagination (limit: %.0f, offset: %.0f)\n", pageSize, pageOffset)
	paginatedResponse, err := client.SigningKeys.ListSigningKeys(ctx, &pageSize, &pageOffset)
	if err != nil {
		log.Printf("Error with paginated request: %v", err)
	} else {
		fmt.Printf("Page 1: Found %d keys\n", len(paginatedResponse.GetAllSigningKeyResponse.Data))

		// Get next page
		pageOffset = 5.0
		fmt.Printf("Fetching next page (limit: %.0f, offset: %.0f)\n", pageSize, pageOffset)
		nextPageResponse, err := client.SigningKeys.ListSigningKeys(ctx, &pageSize, &pageOffset)
		if err != nil {
			log.Printf("Error with next page request: %v", err)
		} else {
			fmt.Printf("Page 2: Found %d keys\n", len(nextPageResponse.GetAllSigningKeyResponse.Data))
		}
	}

	fmt.Println("\n=== Security Demo Complete ===")
	fmt.Println("Remember to implement proper key management in production!")
}

// Helper functions to safely get values from pointers
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

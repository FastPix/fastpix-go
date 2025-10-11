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

	// 1. Create Playlist
	fmt.Println("=== Creating Playlist ===")
	createPlaylistRequest := components.CreatePlaylistRequest{
		Name:        "My Test Playlist",
		Description: fastpixgo.Pointer("A test playlist created by SDK"),
		ReferenceID: "testplaylistref123",
		Type:        components.CreatePlaylistRequestTypeManual,
		Metadata: map[string]string{
			"category": "demo",
			"public":   "true",
		},
	}

	playlistResponse, err := client.Playlist.CreateAPlaylist(ctx, createPlaylistRequest)
	if err != nil {
		log.Printf("Error creating playlist: %v", err)
	} else {
		if playlistResponse.PlaylistCreatedResponse != nil && playlistResponse.PlaylistCreatedResponse.Data != nil && playlistResponse.PlaylistCreatedResponse.Data.ID != nil {
			playlistID := playlistResponse.PlaylistCreatedResponse.Data.ID
			fmt.Printf("Playlist created successfully! ID: %s\n", *playlistID)
		} else {
			fmt.Println("Playlist created but no ID returned")
		}
	}

	// 2. List All Playlists
	fmt.Println("\n=== Listing All Playlists ===")
	limit := int64(10)
	offset := int64(0)

	playlistsResponse, err := client.Playlist.GetAllPlaylists(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing playlists: %v", err)
	} else {
		fmt.Printf("Found %d playlists:\n", len(playlistsResponse.GetAllPlaylistsResponse.Data))
		for i, playlist := range playlistsResponse.GetAllPlaylistsResponse.Data {
			fmt.Printf("  %d. ID: %s, Name: %s, Type: %s\n", 
				i+1, *playlist.ID, getStringValue(playlist.Name), getStringValue(playlist.Type))
		}
	}

	// 3. Get Specific Playlist
	if playlistsResponse.GetAllPlaylistsResponse != nil && len(playlistsResponse.GetAllPlaylistsResponse.Data) > 0 {
		playlistID := playlistsResponse.GetAllPlaylistsResponse.Data[0].ID
		fmt.Printf("\n=== Getting Playlist Details for ID: %s ===\n", *playlistID)

		playlistDetailsResponse, err := client.Playlist.GetPlaylistByID(ctx, *playlistID)
		if err != nil {
			log.Printf("Error getting playlist details: %v", err)
		} else {
			playlist := playlistDetailsResponse.PlaylistByIDResponse.Data
			fmt.Printf("Name: %s\n", getStringValue(playlist.Name))
			fmt.Printf("Description: %s\n", getStringValue(playlist.Description))
			fmt.Printf("Type: %s\n", getStringValue(playlist.Type))
			fmt.Printf("Media Count: %d\n", getInt64Value(playlist.MediaCount))
		}

		// 4. Update Playlist
		fmt.Printf("\n=== Updating Playlist: %s ===\n", *playlistID)
		updatePlaylistRequest := components.UpdatePlaylistRequest{
			Name:        "Updated Test Playlist",
			Description: "Updated description for testing",
			Metadata: map[string]string{
				"category": "updated-demo",
				"public":   "true",
				"updated":  "true",
			},
		}

		updatePlaylistResponse, err := client.Playlist.UpdateAPlaylist(ctx, *playlistID, updatePlaylistRequest)
		if err != nil {
			log.Printf("Error updating playlist: %v", err)
		} else {
			fmt.Println("Playlist updated successfully!")
		}

		// 5. Add Media to Playlist
		fmt.Printf("\n=== Adding Media to Playlist: %s ===\n", *playlistID)
		
		// First, get some media to add to the playlist
		mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		if err != nil {
			log.Printf("Error listing media: %v", err)
		} else if mediaResponse.Object != nil && len(mediaResponse.Object.Data) > 0 {
			mediaID := mediaResponse.Object.Data[0].ID
			fmt.Printf("Adding media %s to playlist\n", *mediaID)

			addMediaRequest := components.MediaIdsRequest{
				MediaIds: []string{*mediaID},
			}

			addMediaResponse, err := client.Playlist.AddMediaToPlaylist(ctx, *playlistID, addMediaRequest)
			if err != nil {
				log.Printf("Error adding media to playlist: %v", err)
			} else {
				fmt.Println("Media added to playlist successfully!")
				fmt.Printf("Response: %+v\n", addMediaResponse)
			}

			// 6. Change Media Order in Playlist
			fmt.Printf("\n=== Changing Media Order in Playlist: %s ===\n", *playlistID)
			changeOrderRequest := components.ChangeMediaOrderRequest{
				MediaIds: []string{*mediaID},
			}

			changeOrderResponse, err := client.Playlist.ChangeMediaOrderInPlaylist(ctx, *playlistID, changeOrderRequest)
			if err != nil {
				log.Printf("Error changing media order: %v", err)
			} else {
				fmt.Println("Media order changed successfully!")
			}

			// 7. Delete Media from Playlist
			fmt.Printf("\n=== Deleting Media from Playlist: %s ===\n", *playlistID)
			deleteMediaRequest := components.MediaIdsRequest{
				MediaIds: []string{*mediaID},
			}

			deleteMediaResponse, err := client.Playlist.DeleteMediaFromPlaylist(ctx, *playlistID, deleteMediaRequest)
			if err != nil {
				log.Printf("Error deleting media from playlist: %v", err)
			} else {
				fmt.Println("Media deleted from playlist successfully!")
			}
		} else {
			fmt.Println("No media available to add to playlist")
		}

		// 8. Delete Playlist
		fmt.Printf("\n=== Deleting Playlist: %s ===\n", *playlistID)
		deletePlaylistResponse, err := client.Playlist.DeleteAPlaylist(ctx, *playlistID)
		if err != nil {
			log.Printf("Error deleting playlist: %v", err)
		} else {
			fmt.Println("Playlist deleted successfully!")
		}
	}

	// 9. Media Playback Management
	fmt.Println("\n=== Media Playback Management ===")
	
	// Get media for playback operations
	mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
	} else if mediaResponse.Object != nil && len(mediaResponse.Object.Data) > 0 {
		mediaID := mediaResponse.Object.Data[0].ID
		fmt.Printf("Working with media: %s\n", *mediaID)

		// Create Playback ID for Media
		fmt.Printf("\n=== Creating Playback ID for Media: %s ===\n", *mediaID)
		createPlaybackRequest := operations.CreateMediaPlaybackIDRequestBody{
			AccessPolicy: components.AccessPolicyPublic,
		}

		playbackResponse, err := client.Playback.CreateMediaPlaybackID(ctx, *mediaID, &createPlaybackRequest)
		if err != nil {
			log.Printf("Error creating playback ID: %v", err)
		} else {
			playbackID := playbackResponse.Object.Data.PlaybackIds[0].ID
			fmt.Printf("Playback ID created successfully! ID: %s\n", *playbackID)

			// Get Playback ID Details
			fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", *playbackID)
			playbackDetailsResponse, err := client.Playback.GetPlaybackID(ctx, *mediaID, *playbackID)
			if err != nil {
				log.Printf("Error getting playback ID details: %v", err)
			} else {
				fmt.Printf("Playback ID details retrieved successfully!\n")
				fmt.Printf("Access Policy: %s\n", getStringValue(playbackDetailsResponse.Object.Data.AccessPolicy))
				fmt.Printf("Created: %s\n", getStringValue(playbackDetailsResponse.Object.Data.CreatedAt))
			}

			// Delete Playback ID
			fmt.Printf("\n=== Deleting Playback ID: %s ===\n", *playbackID)
			deletePlaybackResponse, err := client.Playback.DeleteMediaPlaybackID(ctx, *mediaID, *playbackID)
			if err != nil {
				log.Printf("Error deleting playback ID: %v", err)
			} else {
				fmt.Println("Playback ID deleted successfully!")
			}
		}
	} else {
		fmt.Println("No media available for playback operations")
	}

	// 10. DRM Configuration
	fmt.Println("\n=== DRM Configuration ===")
	
	// List DRM Configurations
	fmt.Println("Listing DRM Configurations...")
	drmResponse, err := client.DRMConfigurations.GetDrmConfiguration(ctx)
	if err != nil {
		log.Printf("Error listing DRM configurations: %v", err)
	} else {
		fmt.Printf("Found %d DRM configurations\n", len(drmResponse.Object.Data))
		for i, config := range drmResponse.Object.Data {
			fmt.Printf("  %d. ID: %s, Name: %s\n", i+1, *config.ID, getStringValue(config.Name))
		}
	}

	// Get Specific DRM Configuration
	if drmResponse.Object != nil && len(drmResponse.Object.Data) > 0 {
		drmConfigID := drmResponse.Object.Data[0].ID
		fmt.Printf("\n=== Getting DRM Configuration Details: %s ===\n", *drmConfigID)

		drmDetailsResponse, err := client.DRMConfigurations.GetDrmConfigurationByID(ctx, *drmConfigID)
		if err != nil {
			log.Printf("Error getting DRM configuration details: %v", err)
		} else {
			fmt.Printf("DRM configuration details retrieved successfully!\n")
			fmt.Printf("Name: %s\n", getStringValue(drmDetailsResponse.Object.Data.Name))
			fmt.Printf("Type: %s\n", getStringValue(drmDetailsResponse.Object.Data.Type))
		}
	}
}

// Helper functions to safely get values from pointers
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

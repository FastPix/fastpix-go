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

func main() {
	ctx := context.Background()

	client := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
		fastpixgo.WithTimeout(30*time.Second),
	)

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
		printCreatePlaylistResult(playlistResponse)
	}

	fmt.Println("\n=== Listing All Playlists ===")
	limit := int64(10)
	offset := int64(0)

	playlistsResponse, err := client.Playlist.GetAllPlaylists(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing playlists: %v", err)
	} else {
		printPlaylistList(playlistsResponse)
	}

	if playlistsResponse.GetAllPlaylistsResponse != nil && len(playlistsResponse.GetAllPlaylistsResponse.Data) > 0 {
		managePlaylist(ctx, client, *playlistsResponse.GetAllPlaylistsResponse.Data[0].ID)
	}

	manageMediaPlayback(ctx, client)
	manageDRMConfigurations(ctx, client)
}

func printCreatePlaylistResult(playlistResponse *operations.CreateAPlaylistResponse) {
	if playlistResponse.PlaylistCreatedResponse != nil &&
		playlistResponse.PlaylistCreatedResponse.Data != nil &&
		playlistResponse.PlaylistCreatedResponse.Data.ID != nil {
		fmt.Printf("Playlist created successfully! ID: %s\n", *playlistResponse.PlaylistCreatedResponse.Data.ID)
		return
	}
	fmt.Println("Playlist created but no ID returned")
}

func printPlaylistList(playlistsResponse *operations.GetAllPlaylistsResponse) {
	fmt.Printf("Found %d playlists:\n", len(playlistsResponse.GetAllPlaylistsResponse.Data))
	for i, playlist := range playlistsResponse.GetAllPlaylistsResponse.Data {
		fmt.Printf("  %d. ID: %s, Name: %s, Type: %s\n",
			i+1, *playlist.ID, getStringValue(playlist.Name), getStringValue(playlist.Type))
	}
}

func managePlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID string) {
	getPlaylistDetails(ctx, client, playlistID)
	updatePlaylist(ctx, client, playlistID)
	managePlaylistMedia(ctx, client, playlistID)
	deletePlaylist(ctx, client, playlistID)
}

func getPlaylistDetails(ctx context.Context, client *fastpixgo.FastPix, playlistID string) {
	fmt.Printf("\n=== Getting Playlist Details for ID: %s ===\n", playlistID)

	playlistDetailsResponse, err := client.Playlist.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		log.Printf("Error getting playlist details: %v", err)
		return
	}

	playlist := playlistDetailsResponse.PlaylistByIDResponse.Data
	fmt.Printf("Name: %s\n", getStringValue(playlist.Name))
	fmt.Printf("Description: %s\n", getStringValue(playlist.Description))
	fmt.Printf("Type: %s\n", getStringValue(playlist.Type))
	fmt.Printf("Media Count: %d\n", getInt64Value(playlist.MediaCount))
}

func updatePlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID string) {
	fmt.Printf("\n=== Updating Playlist: %s ===\n", playlistID)

	updatePlaylistRequest := components.UpdatePlaylistRequest{
		Name:        "Updated Test Playlist",
		Description: "Updated description for testing",
		Metadata: map[string]string{
			"category": "updated-demo",
			"public":   "true",
			"updated":  "true",
		},
	}

	_, err := client.Playlist.UpdateAPlaylist(ctx, playlistID, updatePlaylistRequest)
	if err != nil {
		log.Printf("Error updating playlist: %v", err)
		return
	}

	fmt.Println("Playlist updated successfully!")
}

func managePlaylistMedia(ctx context.Context, client *fastpixgo.FastPix, playlistID string) {
	fmt.Printf("\n=== Adding Media to Playlist: %s ===\n", playlistID)

	mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
		return
	}

	if mediaResponse.Object == nil || len(mediaResponse.Object.Data) == 0 {
		fmt.Println("No media available to add to playlist")
		return
	}

	mediaID := *mediaResponse.Object.Data[0].ID
	fmt.Printf("Adding media %s to playlist\n", mediaID)

	addMediaToPlaylist(ctx, client, playlistID, mediaID)
	changeMediaOrderInPlaylist(ctx, client, playlistID, mediaID)
	deleteMediaFromPlaylist(ctx, client, playlistID, mediaID)
}

func addMediaToPlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID, mediaID string) {
	addMediaRequest := components.MediaIdsRequest{
		MediaIds: []string{mediaID},
	}

	addMediaResponse, err := client.Playlist.AddMediaToPlaylist(ctx, playlistID, addMediaRequest)
	if err != nil {
		log.Printf("Error adding media to playlist: %v", err)
		return
	}

	fmt.Println("Media added to playlist successfully!")
	fmt.Printf("Response: %+v\n", addMediaResponse)
}

func changeMediaOrderInPlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID, mediaID string) {
	fmt.Printf("\n=== Changing Media Order in Playlist: %s ===\n", playlistID)

	_, err := client.Playlist.ChangeMediaOrderInPlaylist(ctx, playlistID, components.ChangeMediaOrderRequest{
		MediaIds: []string{mediaID},
	})
	if err != nil {
		log.Printf("Error changing media order: %v", err)
		return
	}

	fmt.Println("Media order changed successfully!")
}

func deleteMediaFromPlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID, mediaID string) {
	fmt.Printf("\n=== Deleting Media from Playlist: %s ===\n", playlistID)

	_, err := client.Playlist.DeleteMediaFromPlaylist(ctx, playlistID, components.MediaIdsRequest{
		MediaIds: []string{mediaID},
	})
	if err != nil {
		log.Printf("Error deleting media from playlist: %v", err)
		return
	}

	fmt.Println("Media deleted from playlist successfully!")
}

func deletePlaylist(ctx context.Context, client *fastpixgo.FastPix, playlistID string) {
	fmt.Printf("\n=== Deleting Playlist: %s ===\n", playlistID)

	_, err := client.Playlist.DeleteAPlaylist(ctx, playlistID)
	if err != nil {
		log.Printf("Error deleting playlist: %v", err)
		return
	}

	fmt.Println("Playlist deleted successfully!")
}

func manageMediaPlayback(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== Media Playback Management ===")

	mediaResponse, err := client.ManageVideos.ListMedia(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
		return
	}

	if mediaResponse.Object == nil || len(mediaResponse.Object.Data) == 0 {
		fmt.Println("No media available for playback operations")
		return
	}

	mediaID := *mediaResponse.Object.Data[0].ID
	fmt.Printf("Working with media: %s\n", mediaID)
	managePlaybackID(ctx, client, mediaID)
}

func managePlaybackID(ctx context.Context, client *fastpixgo.FastPix, mediaID string) {
	fmt.Printf("\n=== Creating Playback ID for Media: %s ===\n", mediaID)

	createPlaybackRequest := operations.CreateMediaPlaybackIDRequestBody{
		AccessPolicy: components.AccessPolicyPublic,
	}

	playbackResponse, err := client.Playback.CreateMediaPlaybackID(ctx, mediaID, &createPlaybackRequest)
	if err != nil {
		log.Printf("Error creating playback ID: %v", err)
		return
	}

	playbackID := *playbackResponse.Object.Data.PlaybackIds[0].ID
	fmt.Printf("Playback ID created successfully! ID: %s\n", playbackID)

	getPlaybackIDDetails(ctx, client, mediaID, playbackID)
	deleteMediaPlaybackID(ctx, client, mediaID, playbackID)
}

func getPlaybackIDDetails(ctx context.Context, client *fastpixgo.FastPix, mediaID, playbackID string) {
	fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", playbackID)

	playbackDetailsResponse, err := client.Playback.GetPlaybackID(ctx, mediaID, playbackID)
	if err != nil {
		log.Printf("Error getting playback ID details: %v", err)
		return
	}

	fmt.Printf("Playback ID details retrieved successfully!\n")
	fmt.Printf("Access Policy: %s\n", getStringValue(playbackDetailsResponse.Object.Data.AccessPolicy))
	fmt.Printf("Created: %s\n", getStringValue(playbackDetailsResponse.Object.Data.CreatedAt))
}

func deleteMediaPlaybackID(ctx context.Context, client *fastpixgo.FastPix, mediaID, playbackID string) {
	fmt.Printf("\n=== Deleting Playback ID: %s ===\n", playbackID)

	_, err := client.Playback.DeleteMediaPlaybackID(ctx, mediaID, playbackID)
	if err != nil {
		log.Printf("Error deleting playback ID: %v", err)
		return
	}

	fmt.Println("Playback ID deleted successfully!")
}

func manageDRMConfigurations(ctx context.Context, client *fastpixgo.FastPix) {
	fmt.Println("\n=== DRM Configuration ===")
	fmt.Println("Listing DRM Configurations...")

	drmResponse, err := client.DRMConfigurations.GetDrmConfiguration(ctx)
	if err != nil {
		log.Printf("Error listing DRM configurations: %v", err)
		return
	}

	if drmResponse.Object == nil || len(drmResponse.Object.Data) == 0 {
		fmt.Println("No DRM configurations found")
		return
	}

	fmt.Printf("Found %d DRM configurations\n", len(drmResponse.Object.Data))
	for i, config := range drmResponse.Object.Data {
		fmt.Printf("  %d. ID: %s, Name: %s\n", i+1, *config.ID, getStringValue(config.Name))
	}

	getDrmConfigurationByID(ctx, client, *drmResponse.Object.Data[0].ID)
}

func getDrmConfigurationByID(ctx context.Context, client *fastpixgo.FastPix, drmConfigID string) {
	fmt.Printf("\n=== Getting DRM Configuration Details: %s ===\n", drmConfigID)

	drmDetailsResponse, err := client.DRMConfigurations.GetDrmConfigurationByID(ctx, drmConfigID)
	if err != nil {
		log.Printf("Error getting DRM configuration details: %v", err)
		return
	}

	fmt.Printf("DRM configuration details retrieved successfully!\n")
	fmt.Printf("Name: %s\n", getStringValue(drmDetailsResponse.Object.Data.Name))
	fmt.Printf("Type: %s\n", getStringValue(drmDetailsResponse.Object.Data.Type))
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
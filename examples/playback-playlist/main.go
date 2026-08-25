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
	// Reference IDs must be alphanumeric and unique per workspace, so suffix a
	// timestamp to keep re-runs from colliding.
	createPlaylistRequest := components.CreateCreatePlaylistRequestManual(
		components.CreatePlaylistRequestManual{
			Name:        "My Test Playlist",
			ReferenceID: fmt.Sprintf("sdkexample%d", time.Now().Unix()),
			Description: fastpixgo.Pointer("A test playlist created by SDK"),
		},
	)

	var createdPlaylistID string
	playlistResponse, err := client.Playlists.Create(ctx, createPlaylistRequest)
	if err != nil {
		log.Printf("Error creating playlist: %v", err)
	} else {
		printCreatePlaylistResult(playlistResponse)
		if playlistResponse.PlaylistCreatedResponse != nil {
			if manual := playlistResponse.PlaylistCreatedResponse.GetDataManual(); manual != nil && manual.ID != nil {
				createdPlaylistID = *manual.ID
			}
		}
	}

	fmt.Println("\n=== Listing All Playlists ===")
	limit := int64(10)
	offset := int64(1)

	playlistsResponse, err := client.Playlists.List(ctx, &limit, &offset)
	if err != nil {
		log.Printf("Error listing playlists: %v", err)
	} else {
		printPlaylistList(playlistsResponse)
	}

	// Manage (and finally delete) only the playlist this example created above —
	// never an existing playlist from the workspace.
	if createdPlaylistID != "" {
		managePlaylist(ctx, client, createdPlaylistID)
	} else {
		fmt.Println("\nSkipping playlist management steps (no playlist was created).")
	}

	manageMediaPlayback(ctx, client)
}

func printCreatePlaylistResult(playlistResponse *operations.CreateAPlaylistResponse) {
	if playlistResponse.PlaylistCreatedResponse != nil {
		manual := playlistResponse.PlaylistCreatedResponse.GetDataManual()
		if manual != nil && manual.ID != nil {
			fmt.Printf("Playlist created successfully! ID: %s\n", *manual.ID)
			return
		}
	}
	fmt.Println("Playlist created but no ID returned")
}

func printPlaylistList(playlistsResponse *operations.GetAllPlaylistsResponse) {
	if playlistsResponse.GetAllPlaylistsResponse == nil {
		fmt.Println("No playlists returned")
		return
	}
	playlists := playlistsResponse.GetAllPlaylistsResponse.Data
	fmt.Printf("Found %d playlists:\n", len(playlists))
	for i, playlist := range playlists {
		typ := ""
		if playlist.Type != nil {
			typ = string(*playlist.Type)
		}
		fmt.Printf("  %d. ID: %s, Name: %s, Type: %s\n",
			i+1, getStringValue(playlist.ID), getStringValue(playlist.Name), typ)
	}
}

func managePlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID string) {
	getPlaylistDetails(ctx, client, playlistID)
	updatePlaylist(ctx, client, playlistID)
	managePlaylistMedia(ctx, client, playlistID)
	deletePlaylist(ctx, client, playlistID)
}

func getPlaylistDetails(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID string) {
	fmt.Printf("\n=== Getting Playlist Details for ID: %s ===\n", playlistID)

	playlistDetailsResponse, err := client.Playlist.Get(ctx, playlistID)
	if err != nil {
		log.Printf("Error getting playlist details: %v", err)
		return
	}

	if playlistDetailsResponse.PlaylistByIDResponse == nil {
		fmt.Println("No playlist details returned")
		return
	}

	manual := playlistDetailsResponse.PlaylistByIDResponse.GetDataManual()
	if manual == nil {
		fmt.Println("Playlist is not a manual playlist")
		return
	}
	fmt.Printf("Name: %s\n", getStringValue(manual.Name))
	fmt.Printf("Description: %s\n", getStringValue(manual.Description))
	fmt.Printf("Type: %s\n", string(manual.Type))
	fmt.Printf("Media Count: %d\n", getInt64Value(manual.MediaCount))
}

func updatePlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID string) {
	fmt.Printf("\n=== Updating Playlist: %s ===\n", playlistID)

	updatePlaylistRequest := components.UpdatePlaylistRequest{
		Name:        "Updated Test Playlist",
		Description: "Updated description for testing",
	}

	_, err := client.Playlists.Update(ctx, playlistID, updatePlaylistRequest)
	if err != nil {
		log.Printf("Error updating playlist: %v", err)
		return
	}

	fmt.Println("Playlist updated successfully!")
}

func managePlaylistMedia(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID string) {
	fmt.Printf("\n=== Adding Media to Playlist: %s ===\n", playlistID)

	mediaResponse, err := client.ManageVideos.List(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
		return
	}

	// Only media in the "Ready" state can be added to a playlist.
	mediaID := firstReadyMediaID(mediaResponse)
	if mediaID == "" {
		fmt.Println("No ready media available to add to playlist")
		return
	}
	fmt.Printf("Adding media %s to playlist\n", mediaID)

	addMediaToPlaylist(ctx, client, playlistID, mediaID)
	changeMediaOrderInPlaylist(ctx, client, playlistID, mediaID)
	deleteMediaFromPlaylist(ctx, client, playlistID, mediaID)
}

func addMediaToPlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID, mediaID string) {
	addMediaRequest := components.MediaIdsRequest{
		MediaIds: []string{mediaID},
	}

	addMediaResponse, err := client.Playlist.AddMedia(ctx, playlistID, addMediaRequest)
	if err != nil {
		log.Printf("Error adding media to playlist: %v", err)
		return
	}

	fmt.Println("Media added to playlist successfully!")
	fmt.Printf("Response: %+v\n", addMediaResponse)
}

func changeMediaOrderInPlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID, mediaID string) {
	fmt.Printf("\n=== Changing Media Order in Playlist: %s ===\n", playlistID)

	_, err := client.Playlist.ChangeMediaOrder(ctx, playlistID, components.MediaIdsRequest{
		MediaIds: []string{mediaID},
	})
	if err != nil {
		log.Printf("Error changing media order: %v", err)
		return
	}

	fmt.Println("Media order changed successfully!")
}

func deleteMediaFromPlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID, mediaID string) {
	fmt.Printf("\n=== Deleting Media from Playlist: %s ===\n", playlistID)

	_, err := client.Playlists.DeleteMedia(ctx, playlistID, &components.MediaIdsRequest{
		MediaIds: []string{mediaID},
	})
	if err != nil {
		log.Printf("Error deleting media from playlist: %v", err)
		return
	}

	fmt.Println("Media deleted from playlist successfully!")
}

func deletePlaylist(ctx context.Context, client *fastpixgo.Fastpixgo, playlistID string) {
	fmt.Printf("\n=== Deleting Playlist: %s ===\n", playlistID)

	_, err := client.Playlists.Delete(ctx, playlistID)
	if err != nil {
		log.Printf("Error deleting playlist: %v", err)
		return
	}

	fmt.Println("Playlist deleted successfully!")
}

func manageMediaPlayback(ctx context.Context, client *fastpixgo.Fastpixgo) {
	fmt.Println("\n=== Media Playback Management ===")

	mediaResponse, err := client.ManageVideos.List(ctx, nil, nil, nil)
	if err != nil {
		log.Printf("Error listing media: %v", err)
		return
	}

	// Playback IDs can only be created once a media has finished processing,
	// so pick the first media that is in the "Ready" state.
	mediaID := firstReadyMediaID(mediaResponse)
	if mediaID == "" {
		fmt.Println("No ready media available for playback operations")
		return
	}

	fmt.Printf("Working with media: %s\n", mediaID)
	managePlaybackID(ctx, client, mediaID)
}

func managePlaybackID(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID string) {
	fmt.Printf("\n=== Creating Playback ID for Media: %s ===\n", mediaID)

	createPlaybackRequest := operations.CreateMediaPlaybackIDRequestBody{
		AccessPolicy: components.AccessPolicyPublic,
	}

	playbackResponse, err := client.Playback.Create(ctx, mediaID, &createPlaybackRequest)
	if err != nil {
		log.Printf("Error creating playback ID: %v", err)
		return
	}

	if playbackResponse.Object == nil || playbackResponse.Object.Data == nil ||
		playbackResponse.Object.Data.ID == nil {
		fmt.Println("Playback ID created but no ID returned")
		return
	}

	playbackID := *playbackResponse.Object.Data.ID
	fmt.Printf("Playback ID created successfully! ID: %s\n", playbackID)

	getPlaybackIDDetails(ctx, client, mediaID, playbackID)
	deleteMediaPlaybackID(ctx, client, mediaID, playbackID)
}

func getPlaybackIDDetails(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID, playbackID string) {
	fmt.Printf("\n=== Getting Playback ID Details: %s ===\n", playbackID)

	playbackDetailsResponse, err := client.Playback.GetByID(ctx, mediaID, playbackID)
	if err != nil {
		log.Printf("Error getting playback ID details: %v", err)
		return
	}

	if playbackDetailsResponse.Object == nil || playbackDetailsResponse.Object.Data == nil {
		fmt.Println("No playback ID details returned")
		return
	}

	fmt.Printf("Playback ID details retrieved successfully!\n")
	accessPolicy := ""
	if playbackDetailsResponse.Object.Data.AccessPolicy != nil {
		accessPolicy = string(*playbackDetailsResponse.Object.Data.AccessPolicy)
	}
	fmt.Printf("Access Policy: %s\n", accessPolicy)
	fmt.Printf("Playback ID: %s\n", getStringValue(playbackDetailsResponse.Object.Data.ID))
}

func deleteMediaPlaybackID(ctx context.Context, client *fastpixgo.Fastpixgo, mediaID, playbackID string) {
	fmt.Printf("\n=== Deleting Playback ID: %s ===\n", playbackID)

	_, err := client.Playback.Delete(ctx, mediaID, playbackID)
	if err != nil {
		log.Printf("Error deleting playback ID: %v", err)
		return
	}

	fmt.Println("Playback ID deleted successfully!")
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

// firstReadyMediaID returns the ID of the first media in the "Ready" state, or
// "" if none. Most media operations require processing to have finished.
func firstReadyMediaID(mediaResponse *operations.ListMediaResponse) string {
	if mediaResponse == nil || mediaResponse.Object == nil {
		return ""
	}
	for _, media := range mediaResponse.Object.Data {
		if media.ID != nil && media.Status != nil && *media.Status == components.GetAllMediaResponseStatusReady {
			return *media.ID
		}
	}
	return ""
}

package tests

import (
	"context"
	"testing"

	"github.com/fastpix/fastpix-go/models/components"
	"github.com/fastpix/fastpix-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlaybackManagement(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Create Media Playback ID", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Test creating playback ID
			playbackRequest := operations.CreateMediaPlaybackIDRequestBody{
				AccessPolicy: components.AccessPolicyPublic,
			}

			response, err := config.Client.Playback.CreateMediaPlaybackID(ctx, *mediaID, &playbackRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No media found to test CreateMediaPlaybackID")
		}
	})

	t.Run("Test Get Playback ID", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Get playback IDs for the media
			playbackResponse, err := config.Client.Playback.CreateMediaPlaybackID(ctx, *mediaID, nil)
			if err == nil && playbackResponse.Object != nil && playbackResponse.Object.Data != nil && len(playbackResponse.Object.Data.PlaybackIds) > 0 {
				playbackID := playbackResponse.Object.Data.PlaybackIds[0].ID
				require.NotNil(t, playbackID)

				// Test getting specific playback ID
				response, err := config.Client.Playback.GetPlaybackID(ctx, *mediaID, *playbackID)
				require.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				t.Skip("Could not create playback ID for testing")
			}
		} else {
			t.Skip("No media found to test GetPlaybackID")
		}
	})

	t.Run("Test Delete Media Playback ID", func(t *testing.T) {
		// First, get a list of media to find a valid ID
		listResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if listResponse.Object != nil && len(listResponse.Object.Data) > 0 {
			mediaID := listResponse.Object.Data[0].ID
			require.NotNil(t, mediaID)

			// Create a playback ID first
			playbackResponse, err := config.Client.Playback.CreateMediaPlaybackID(ctx, *mediaID, nil)
			if err == nil && playbackResponse.Object != nil && playbackResponse.Object.Data != nil && len(playbackResponse.Object.Data.PlaybackIds) > 0 {
				playbackID := playbackResponse.Object.Data.PlaybackIds[0].ID
				require.NotNil(t, playbackID)

				// Test deleting playback ID
				response, err := config.Client.Playback.DeleteMediaPlaybackID(ctx, *mediaID, *playbackID)
				require.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				t.Skip("Could not create playback ID for testing")
			}
		} else {
			t.Skip("No media found to test DeleteMediaPlaybackID")
		}
	})
}

func TestPlaylistManagement(t *testing.T) {
	config := GetTestConfig(t)
	ctx := context.Background()

	t.Run("Test Create Playlist", func(t *testing.T) {
		// Test creating a new playlist
		createRequest := components.CreatePlaylistRequest{
			Name:        "Test Playlist",
			Description: stringPtr("Test playlist created by SDK"),
			ReferenceID: "testplaylistref123",
			Type:        components.CreatePlaylistRequestTypeManual,
		}

		response, err := config.Client.Playlist.CreateAPlaylist(ctx, createRequest)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.PlaylistCreatedResponse)
		
		// Store the playlist ID for cleanup
		if response.PlaylistCreatedResponse != nil && response.PlaylistCreatedResponse.Data != nil {
			t.Logf("Created playlist with ID: %s", *response.PlaylistCreatedResponse.Data.ID)
		}
	})

	t.Run("Test Get All Playlists", func(t *testing.T) {
		// Test listing all playlists
		response, err := config.Client.Playlist.GetAllPlaylists(ctx, nil, nil)
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, response.GetAllPlaylistsResponse)
	})

	t.Run("Test Get Playlist by ID", func(t *testing.T) {
		// First, get a list of playlists to find a valid ID
		listResponse, err := config.Client.Playlist.GetAllPlaylists(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.GetAllPlaylistsResponse != nil && len(listResponse.GetAllPlaylistsResponse.Data) > 0 {
			playlistID := listResponse.GetAllPlaylistsResponse.Data[0].ID
			require.NotNil(t, playlistID)

			// Test getting specific playlist
			response, err := config.Client.Playlist.GetPlaylistByID(ctx, *playlistID)
			require.NoError(t, err)
			assert.NotNil(t, response)
			assert.NotNil(t, response.PlaylistByIDResponse)
		} else {
			t.Skip("No playlists found to test GetPlaylistByID")
		}
	})

	t.Run("Test Update Playlist", func(t *testing.T) {
		// First, get a list of playlists to find a valid ID
		listResponse, err := config.Client.Playlist.GetAllPlaylists(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.GetAllPlaylistsResponse != nil && len(listResponse.GetAllPlaylistsResponse.Data) > 0 {
			playlistID := listResponse.GetAllPlaylistsResponse.Data[0].ID
			require.NotNil(t, playlistID)

			// Test updating playlist
			updateRequest := components.UpdatePlaylistRequest{
				Name:        "Updated Test Playlist",
				Description: "Updated by SDK test",
			}

			response, err := config.Client.Playlist.UpdateAPlaylist(ctx, *playlistID, updateRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No playlists found to test UpdateAPlaylist")
		}
	})

	t.Run("Test Add Media to Playlist", func(t *testing.T) {
		// First, get a list of playlists to find a valid ID
		playlistListResponse, err := config.Client.Playlist.GetAllPlaylists(ctx, nil, nil)
		require.NoError(t, err)

		// Get a list of media to add to playlist
		mediaListResponse, err := config.Client.ManageVideos.ListMedia(ctx, nil, nil, nil)
		require.NoError(t, err)

		if playlistListResponse.GetAllPlaylistsResponse != nil && len(playlistListResponse.GetAllPlaylistsResponse.Data) > 0 &&
			mediaListResponse.Object != nil && len(mediaListResponse.Object.Data) > 0 {
			
			playlistID := playlistListResponse.GetAllPlaylistsResponse.Data[0].ID
			mediaID := mediaListResponse.Object.Data[0].ID
			require.NotNil(t, playlistID)
			require.NotNil(t, mediaID)

			// Test adding media to playlist
			mediaRequest := components.MediaIdsRequest{
				MediaIds: []string{*mediaID},
			}

			response, err := config.Client.Playlist.AddMediaToPlaylist(ctx, *playlistID, mediaRequest)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No playlists or media found to test AddMediaToPlaylist")
		}
	})

	t.Run("Test Delete Playlist", func(t *testing.T) {
		// First, get a list of playlists to find a valid ID
		listResponse, err := config.Client.Playlist.GetAllPlaylists(ctx, nil, nil)
		require.NoError(t, err)

		if listResponse.GetAllPlaylistsResponse != nil && len(listResponse.GetAllPlaylistsResponse.Data) > 0 {
			playlistID := listResponse.GetAllPlaylistsResponse.Data[0].ID
			require.NotNil(t, playlistID)

			// Test deleting playlist
			response, err := config.Client.Playlist.DeleteAPlaylist(ctx, *playlistID)
			require.NoError(t, err)
			assert.NotNil(t, response)
		} else {
			t.Skip("No playlists found to test DeleteAPlaylist")
		}
	})
}

// Helper function to create string pointers
func stringPtr(v string) *string {
	return &v
}

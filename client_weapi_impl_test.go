// +build integration

package netease

import (
	"fmt"
	"github.com/jixunmoe-go/netease/search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSongInfoUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Song(38019092, 5308028)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp.Songs), "Does not receive data for 2 songs")

	assert.Equal(t, "Lifeline", resp.Songs[0].Name, "First song should be 'Life Line'")
	assert.Equal(t, "My Soul", resp.Songs[1].Name, "Second song should be 'My Soul'")
}

func TestPlaylistUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Playlist(3094646186)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Playlist.Tracks[0].Name)
	fmt.Printf("total tracks: %d\n", len(resp.Playlist.Tracks))
}

func TestAlbumUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Album(83351855)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "蔡徐坤", resp.Album.Artist.Name, "The album should belong to cxk")
}

func TestSongURLUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.SongURL(320000, 38019092, 5308028)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 38019092, resp.Data[0].ID)
	assert.Equal(t, 5308028, resp.Data[1].ID)
}

func TestSongURLv1UsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.SongURLv1("standard", "aac", 38019092, 5308028)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, 38019092, resp.Data[0].ID)
	assert.Equal(t, 5308028, resp.Data[1].ID)
}

func TestArtistUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Artist(33184, 50)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Ed Sheeran", resp.Artist.Name, "The artist should be Ed Sheeran")
}

func TestLyricUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Lyric(443277477)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, byte('['), resp.Lyrics.LRC[0])
}

func TestSearchUsingWEAPI(t *testing.T) {
	api := New()
	api.UseWEAPIClient()
	resp, err := api.Search("蔡徐坤", search.Artist, 0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, resp.Result.ArtistCount, uint(0), "Should be able to find at least one artist")
	fmt.Printf("First artist: %s\n", resp.Result.Artists[0].Name)
}

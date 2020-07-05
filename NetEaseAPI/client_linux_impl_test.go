// +build integration

package NetEaseAPI

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSongInfo_Linux(t *testing.T) {
	api := New()
	resp, err := api.Song(38019092, 5308028)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp.Songs), "Does not receive data for 2 songs")

	assert.Equal(t, "Lifeline", resp.Songs[0].Name, "First song should be 'Life Line'")
	assert.Equal(t, "My Soul", resp.Songs[1].Name, "Second song should be 'My Soul'")
}

func TestPlaylist_Linux(t *testing.T) {
	api := New()
	resp, err := api.Playlist(3094646186)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Playlist.Tracks[0].Name)
	fmt.Printf("total tracks: %d\n", len(resp.Playlist.Tracks))
}

func TestAlbum_Linux(t *testing.T) {
	api := New()
	resp, err := api.Album(83351855)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "蔡徐坤", resp.Album.Artist.Name, "The album should belong to cxk")
}

func TestSongURL_Linux(t *testing.T) {
	api := New()
	resp, err := api.SongURL(320000, 38019092, 5308028)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 38019092, resp.Data[0].ID)
	assert.Equal(t, 5308028, resp.Data[1].ID)
}

func TestArtist_Linux(t *testing.T) {
	api := New()
	resp, err := api.Artist(33184, 50)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Ed Sheeran", resp.Artist.Name, "The artist should be Ed Sheeran")
}

func TestLyric_Linux(t *testing.T) {
	api := New()
	resp, err := api.Lyric(443277477)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, byte('['), resp.Lyrics.LRC[0])
}

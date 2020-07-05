// +build integration

package NetEaseAPI

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSongInfo(t *testing.T) {
	api := New()
	resp, err := api.Song(38019092, 5308028)
	if err == nil && resp != nil {
		if len(resp.Songs) != 2 {
			t.Error("Does not receive data for 2 songs")
		}
		assert.Equal(t, "Lifeline", resp.Songs[0].Name, "First song should be 'Life Line'")
		assert.Equal(t, "My Soul", resp.Songs[1].Name, "Second song should be 'My Soul'")
	} else {
		t.Error(err)
	}
}

func TestPlaylist(t *testing.T) {
	api := New()
	resp, err := api.Playlist(3094646186)
	if err == nil && resp != nil {
		fmt.Println(resp.Playlist.Tracks[0].Name)
		fmt.Printf("total tracks: %d\n", len(resp.Playlist.Tracks))
	}
}

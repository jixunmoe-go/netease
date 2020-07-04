package api

import (
	"fmt"
	"testing"
)

func TestSongInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping song test in this mode.")
	}

	api := New()
	resp, err := api.Song(5308028, 5308029)
	if err == nil && resp != nil {
		fmt.Println("First song (name)")
		fmt.Println(resp.Songs[0].Name)
		fmt.Println("Second song (name, translation)")
		fmt.Printf("%s, %s\n", resp.Songs[1].Name, resp.Songs[1].TranslatedName[0])
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

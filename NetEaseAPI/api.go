package NetEaseAPI

import (
	"errors"
	"fmt"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/search"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/types"
	"strings"
)

func makeSongID(id uint64) string {
	return fmt.Sprintf("{id:%d}", id)
}

func makeSongIDs(ids []uint64) string {
	sb := strings.Builder{}
	sb.WriteByte('[')
	sb.WriteString(makeSongID(ids[0]))
	for i := 1; i < len(ids); i++ {
		sb.WriteByte(',')
		sb.WriteString(makeSongID(ids[i]))
	}
	sb.WriteByte(']')
	return sb.String()
}

func (n *NetEase) Song(ids ...uint64) (*types.SongResp, error) {
	if len(ids) < 1 {
		return nil, errors.New("NetEase::Song: no id supplied")
	}

	var result types.SongResp
	err := n.Client.Request(n, &result, "POST", "/v3/song/detail",
		map[string]interface{}{
			"c": makeSongIDs(ids),
		},
	)

	return &result, err
}

// SongURL requests for the MP3 url of a list of song ids.
// bitrate can be one of the following: 128000, 192000, 320000.
// Other undocumented bitrate *might* be available (999000 = flac format?).
func (n *NetEase) SongURL(bitrate int, ids ...uint64) (*types.SongURLResp, error) {
	var result types.SongURLResp
	err := n.Client.Request(n, &result, "POST", "/song/enhance/player/url",
		map[string]interface{}{
			"ids": ids,
			"br":  bitrate,
		},
	)
	return &result, err
}

// SongURL requests for the url of a list of song ids.
// level: "standard" (128k?), "exhigh" (320k?)
// encodeType: "aac", "mp3"
func (n *NetEase) SongURLV1(level, encodeType string, ids ...uint64) (*types.SongURLResp, error) {
	var result types.SongURLResp
	err := n.Client.Request(n, &result, "POST", "/song/enhance/player/url/v1",
		map[string]interface{}{
			"ids":        ids,
			"level":      level,
			"encodeType": encodeType,
		},
	)
	return &result, err
}

func (n *NetEase) Playlist(id uint64) (*types.PlayListResp, error) {
	var result types.PlayListResp
	err := n.Client.Request(n, &result, "POST", "/v3/playlist/detail",
		map[string]interface{}{
			"id": id,
			"n":  1000,
		},
	)
	return &result, err
}

func (n *NetEase) Album(id uint64) (*types.AlbumResp, error) {
	var result types.AlbumResp
	err := n.Client.Request(n, &result, "POST", fmt.Sprintf("/v1/album/%d", id),
		map[string]interface{}{
			"id": id,
		},
	)
	return &result, err
}

// Artist can get information about a specific artist.
// limit: the number of "hot" songs to return.
// a reasonable limit would be 50.
func (n *NetEase) Artist(id, limit uint64) (*types.ArtistResp, error) {
	var result types.ArtistResp
	err := n.Client.Request(n, &result, "POST", fmt.Sprintf("/v1/artist/%d", id),
		map[string]interface{}{
			"id":  id,
			"ext": true,
			"top": limit,
		},
	)
	return &result, err
}

func (n *NetEase) Lyric(id uint64) (*types.LyricResp, error) {
	var result types.LyricResp
	err := n.Client.Request(n, &result, "POST", "/song/lyric",
		map[string]interface{}{
			"id": id,
			"os": "linux",
			"lv": -1,
			"kv": -1,
			"tv": -1,
		},
	)
	return &result, err
}

// Search for a song.
// keyword: the search phrase
// category: search category
// offset: the page to search (starting from 0)
// limit: the number of items per page
func (n *NetEase) Search(keyword string, category search.Type, offset, limit uint) (*types.SearchResp, error) {
	var result types.SearchResp
	err := n.Client.Request(n, &result, "POST", "/cloudsearch/pc",
		map[string]interface{}{
			"s":      keyword,
			"type":   category,
			"limit":  limit,
			"total":  true,
			"offset": offset,
		},
	)
	return &result, err
}

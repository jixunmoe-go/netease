package NetEaseAPI

import (
	"errors"
	"fmt"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/types"
	"strings"
)

func (n *NetEase) Song(ids ...uint64) (*types.SongResp, error) {
	if len(ids) < 1 {
		return nil, errors.New("NetEase::Song: no id supplied")
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[{id:%d}", ids[0]))
	for i := 1; i < len(ids); i++ {
		sb.WriteString(fmt.Sprintf(",{id:%d}", ids[i]))
	}
	sb.WriteByte(']')

	var result types.SongResp
	err := n.Client.Request(n, &result, "POST", "/v3/song/detail",
		map[string]interface{}{
			"c": sb.String(),
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

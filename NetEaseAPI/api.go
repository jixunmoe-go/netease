package NetEaseAPI

import (
	"encoding/json"
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

	param, err := encodeParam(
		"POST",
		"http://music.163.com/api/v3/song/detail",
		map[string]interface{}{
			"c": sb.String(),
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.SongResp
	resp, err := n.linuxRequest(param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &result)
	result.Response = resp
	return &result, err
}

func (n *NetEase) Playlist(id uint64) (*types.PlayListResp, error) {
	param, err := encodeParam(
		"POST",
		"http://music.163.com/api/v3/playlist/detail",
		map[string]uint64{
			"id": id,
			"n":  1000,
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.PlayListResp
	resp, err := n.linuxRequest(param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &result)
	result.Response = resp
	return &result, err
}

package types

import "encoding/json"

func (r *AlbumResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	if err == nil {
		r.Response = resp
	}
	return err
}

func (r *PlayListResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	if err == nil {
		r.Response = resp
	}
	return err
}

func (r *SongResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	if err == nil {
		r.Response = resp
	}
	return err
}

func (r *SongURLResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	if err == nil {
		r.Response = resp
	}
	return err
}

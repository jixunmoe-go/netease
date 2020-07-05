package types

import "encoding/json"

func (r *AlbumResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *PlayListResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *SongResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *SongURLResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *ArtistResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *LyricResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

func (r *SearchResp) Deserialize(resp string) error {
	err := json.Unmarshal([]byte(resp), r)
	r.Response = resp
	return err
}

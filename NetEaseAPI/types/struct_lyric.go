package types

type LRCData struct {
	Version int `json:"version"`

	// LRC formatted Lyrics
	LRC string `json:"lyric,omitempty"`
}

type LyricResp struct {
	UnknownSgc bool `json:"sgc"`
	UnknownSfy bool `json:"sfy"`
	UnknownQfy bool `json:"qfy"`

	Lyrics          LRCData `json:"lrc,omitempty"`
	LocalisedLyrics LRCData `json:"tlyric,omitempty"`

	Code     int    `json:"code"`
	Response string `json:"-"`
}

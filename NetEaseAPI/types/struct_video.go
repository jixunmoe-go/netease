package types

type VideoInfo struct {
	CoverURL string `json:"coverUrl"`
	Title    string `json:"title"`
	// Duration in MS
	DurationMS int             `json:"durationms"`
	PlayTime   int             `json:"playTime"`
	Type       int             `json:"type"`
	Creator    []BasicUserInfo `json:"creator"`
	AliasName  interface{}     `json:"aliaName"`
	TransName  interface{}     `json:"transName"`
	VideoID    string          `json:"vid"`
	MarkTypes  []interface{}   `json:"markTypes"`
	Algorithm  string          `json:"alg"`
}

package types

type FreeTrialInfo struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type SongURLInfo struct {
	ID int `json:"id"`

	// 200: OK
	Code int `json:"code"`

	URL     string `json:"url"`
	Bitrate int    `json:"br"`
	Bytes   int    `json:"size"`
	MD5     string `json:"md5"`

	// Usually 1200, 20 minutes?
	ExpireIn      int            `json:"expi"`
	Type          string         `json:"type"`
	Gain          float64        `json:"gain"`
	Fee           int            `json:"fee"`
	Payed         int            `json:"payed"`
	UnknownUf     interface{}    `json:"uf"`
	Flag          int            `json:"flag"`
	CanExtend     bool           `json:"canExtend"`
	FreeTrialInfo *FreeTrialInfo `json:"freeTrialInfo"`

	// e.g. "standard"
	Level string `json:"level"`

	// e.g. "mp3"
	EncodeType string `json:"encodeType"`
}

type SongURLResp struct {
	Data     []SongURLInfo `json:"data"`
	Code     int           `json:"code"`
	Response string        `json:"-"`
}

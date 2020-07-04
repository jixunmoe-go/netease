package types

type Artist struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Tns   []interface{} `json:"tns"`
	Alias []interface{} `json:"alias"`
}

type Album struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	ImageURL   string        `json:"picUrl"`
	UnknownTns []interface{} `json:"tns"`
	ImageID    int64         `json:"pic"`
}

type ResourceInfo struct {
	BitRate   int     `json:"br"`
	FileID    int     `json:"fid"`
	Size      int     `json:"size"`
	UnknownVd float64 `json:"vd"`
}

type BasicSongInfo struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	TranslatedName []string `json:"tns,omitempty"`
	UnknownPst     int      `json:"pst"`
	UnknownT       int      `json:"t"`
	Artists        []Artist `json:"ar"`
	// unknown
	Aliases              []string      `json:"alia,omitempty"`
	UnknownPop           float64       `json:"pop"`
	UnknownSt            int           `json:"st"`
	UnknownRt            string        `json:"rt"`
	UnknownFee           int           `json:"fee"`
	UnknownV             int           `json:"v"`
	UnknownCrbt          interface{}   `json:"crbt"`
	UnknownCf            string        `json:"cf"`
	Album                Album         `json:"al"`
	UnknownDt            int           `json:"dt"`
	HighQuality          ResourceInfo  `json:"h"`
	MediumQuality        ResourceInfo  `json:"m"`
	LowQuality           ResourceInfo  `json:"l"`
	UnknownA             interface{}   `json:"a"`
	UnknownCd            string        `json:"cd"`
	UnknownNo            int           `json:"no"`
	UnknownRtURL         interface{}   `json:"rtUrl"`
	UnknownFtype         int           `json:"ftype"`
	UnknownRtUrls        []interface{} `json:"rtUrls"`
	UnknownDjID          int           `json:"djId"`
	Copyright            int           `json:"copyright"`
	UnknownSID           int           `json:"s_id"`
	Mark                 int           `json:"mark"`
	OriginalCoverType    int           `json:"originCoverType"`
	Single               int           `json:"single"`
	NoCopyrightRecommend interface{}   `json:"noCopyrightRcmd"`
	HasMV                int           `json:"mv"`
	ResourceType         int           `json:"rtype"`
	ResourceURL          string        `json:"rurl"`
	UnknownMst           int           `json:"mst"`
	UnknownCp            int           `json:"cp"`
	PublishTime          int64         `json:"publishTime"`
}

type ChargeInfo struct {
	BitRate       int         `json:"rate"`
	ChargeURL     interface{} `json:"chargeUrl"`
	ChargeMessage interface{} `json:"chargeMessage"`
	ChargeType    int         `json:"chargeType"`
}

type PrivilegeInfo struct {
	ID                 int          `json:"id"`
	Fee                int          `json:"fee"`
	Payed              int          `json:"payed"`
	UnknownSt          int          `json:"st"`
	UnknownPl          int          `json:"pl"`
	UnknownDl          int          `json:"dl"`
	UnknownSp          int          `json:"sp"`
	UnknownCp          int          `json:"cp"`
	UnknownSubp        int          `json:"subp"`
	UnknownCs          bool         `json:"cs"`
	MaxBitrate         int          `json:"maxbr"`
	UnknownFl          int          `json:"fl"`
	UnknownToast       bool         `json:"toast"`
	UnknownFlag        int          `json:"flag"`
	PreOrder           bool         `json:"preSell"`
	PlayMaxBitrate     int          `json:"playMaxbr"`
	DownloadMaxBitrate int          `json:"downloadMaxbr"`
	ChargeInfoList     []ChargeInfo `json:"chargeInfoList"`
}

type SongResp struct {
	Songs      []BasicSongInfo `json:"songs"`
	Privileges []PrivilegeInfo `json:"privileges"`
	Code       int             `json:"code"`
}

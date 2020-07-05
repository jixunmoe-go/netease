package types

type TrackID struct {
	ID  int         `json:"id"`
	V   int         `json:"v"`
	At  int64       `json:"at"`
	Alg interface{} `json:"alg"`
}

type BasicPlaylistInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CoverImgURL string `json:"coverImgUrl"`

	Creator BasicPlaylistCreatorInfo `json:"creator"`

	Subscribed   bool        `json:"subscribed"`
	TrackCount   int         `json:"trackCount"`
	UserID       int         `json:"userId"`
	PlayCount    int         `json:"playCount"`
	BookCount    int         `json:"bookCount"`
	OfficialTags interface{} `json:"officialTags,omitempty"`
	Description  string      `json:"description"`
	HighQuality  bool        `json:"highQuality"`
}

type Playlist struct {
	BasicPlaylistInfo
	Creator UserInfo `json:"creator"`

	Subscribers           []interface{} `json:"subscribers"`
	Tracks                []SongInfo    `json:"tracks"`
	TrackIds              []TrackID     `json:"trackIds"`
	UpdateFrequency       interface{}   `json:"updateFrequency"`
	BackgroundCoverID     int           `json:"backgroundCoverId"`
	BackgroundCoverURL    interface{}   `json:"backgroundCoverUrl"`
	TitleImage            int           `json:"titleImage"`
	TitleImageURL         interface{}   `json:"titleImageUrl"`
	EnglishTitle          interface{}   `json:"englishTitle"`
	OpRecommend           bool          `json:"opRecommend"`
	AdType                int           `json:"adType"`
	TrackNumberUpdateTime int64         `json:"trackNumberUpdateTime"`
	SubscribedCount       int           `json:"subscribedCount"`
	CloudTrackCount       int           `json:"cloudTrackCount"`
	CreateTime            int64         `json:"createTime"`
	SpecialType           int           `json:"specialType"`
	UpdateTime            int64         `json:"updateTime"`
	CoverImgID            int64         `json:"coverImgId"`
	NewImported           bool          `json:"newImported"`
	CommentThreadID       string        `json:"commentThreadId"`
	Privacy               int           `json:"privacy"`
	TrackUpdateTime       int64         `json:"trackUpdateTime"`
	TrackCount            int           `json:"trackCount"`
	Ordered               bool          `json:"ordered"`
	Tags                  []interface{} `json:"tags"`
	Status                int           `json:"status"`
	ShareCount            int           `json:"shareCount"`
	CommentCount          int           `json:"commentCount"`
}

type PlayListResp struct {
	Code          int             `json:"code"`
	RelatedVideos interface{}     `json:"relatedVideos"`
	Playlist      Playlist        `json:"playlist"`
	Urls          interface{}     `json:"urls"`
	Privileges    []PrivilegeInfo `json:"privileges"`
	Response      string          `json:"-"`
}

package types

type UserInfo struct {
	DefaultAvatar      bool        `json:"defaultAvatar"`
	Province           int         `json:"province"`
	AuthStatus         int         `json:"authStatus"`
	Followed           bool        `json:"followed"`
	AvatarURL          string      `json:"avatarUrl"`
	AccountStatus      int         `json:"accountStatus"`
	Gender             int         `json:"gender"`
	City               int         `json:"city"`
	Birthday           int64       `json:"birthday"`
	UserID             int         `json:"userId"`
	UserType           int         `json:"userType"`
	Nickname           string      `json:"nickname"`
	Signature          string      `json:"signature"`
	Description        string      `json:"description"`
	DetailDescription  string      `json:"detailDescription"`
	AvatarImgID        int64       `json:"avatarImgId"`
	BackgroundImgID    int64       `json:"backgroundImgId"`
	BackgroundURL      string      `json:"backgroundUrl"`
	Authority          int         `json:"authority"`
	Mutual             bool        `json:"mutual"`
	ExpertTags         interface{} `json:"expertTags"`
	Experts            interface{} `json:"experts"`
	DjStatus           int         `json:"djStatus"`
	VipType            int         `json:"vipType"`
	RemarkName         interface{} `json:"remarkName"`
	AvatarImgIDStr     string      `json:"avatarImgIdStr"`
	BackgroundImgIDStr string      `json:"backgroundImgIdStr"`
}

type TrackID struct {
	ID  int         `json:"id"`
	V   int         `json:"v"`
	At  int64       `json:"at"`
	Alg interface{} `json:"alg"`
}

type Playlist struct {
	Subscribers           []interface{}   `json:"subscribers"`
	Subscribed            bool            `json:"subscribed"`
	Creator               UserInfo        `json:"creator"`
	Tracks                []BasicSongInfo `json:"tracks"`
	TrackIds              []TrackID       `json:"trackIds"`
	UpdateFrequency       interface{}     `json:"updateFrequency"`
	BackgroundCoverID     int             `json:"backgroundCoverId"`
	BackgroundCoverURL    interface{}     `json:"backgroundCoverUrl"`
	TitleImage            int             `json:"titleImage"`
	TitleImageURL         interface{}     `json:"titleImageUrl"`
	EnglishTitle          interface{}     `json:"englishTitle"`
	OpRecommend           bool            `json:"opRecommend"`
	AdType                int             `json:"adType"`
	TrackNumberUpdateTime int64           `json:"trackNumberUpdateTime"`
	SubscribedCount       int             `json:"subscribedCount"`
	CloudTrackCount       int             `json:"cloudTrackCount"`
	UserID                int             `json:"userId"`
	CreateTime            int64           `json:"createTime"`
	HighQuality           bool            `json:"highQuality"`
	SpecialType           int             `json:"specialType"`
	UpdateTime            int64           `json:"updateTime"`
	CoverImgID            int64           `json:"coverImgId"`
	NewImported           bool            `json:"newImported"`
	CommentThreadID       string          `json:"commentThreadId"`
	CoverImgURL           string          `json:"coverImgUrl"`
	Privacy               int             `json:"privacy"`
	TrackUpdateTime       int64           `json:"trackUpdateTime"`
	TrackCount            int             `json:"trackCount"`
	PlayCount             int             `json:"playCount"`
	Ordered               bool            `json:"ordered"`
	Tags                  []interface{}   `json:"tags"`
	Description           interface{}     `json:"description"`
	Status                int             `json:"status"`
	Name                  string          `json:"name"`
	ID                    int             `json:"id"`
	ShareCount            int             `json:"shareCount"`
	CommentCount          int             `json:"commentCount"`
}

type PlayListResp struct {
	Code          int             `json:"code"`
	RelatedVideos interface{}     `json:"relatedVideos"`
	Playlist      Playlist        `json:"playlist"`
	Urls          interface{}     `json:"urls"`
	Privileges    []PrivilegeInfo `json:"privileges"`
}

package types

type BasicUserInfo struct {
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
}

type BasicPlaylistCreatorInfo struct {
	UserID     int         `json:"userId"`
	Nickname   string      `json:"nickname"`
	UserType   int         `json:"userType"`
	AuthStatus int         `json:"authStatus"`
	ExpertTags interface{} `json:"expertTags"`
	Experts    interface{} `json:"experts"`
}

type UserInfo struct {
	BasicPlaylistCreatorInfo

	DefaultAvatar      bool        `json:"defaultAvatar"`
	Province           int         `json:"province"`
	Followed           bool        `json:"followed"`
	AvatarURL          string      `json:"avatarUrl"`
	AccountStatus      int         `json:"accountStatus"`
	Gender             int         `json:"gender"`
	City               int         `json:"city"`
	Birthday           int64       `json:"birthday"`
	Signature          string      `json:"signature"`
	Description        string      `json:"description"`
	DetailDescription  string      `json:"detailDescription"`
	AvatarImgID        int64       `json:"avatarImgId"`
	BackgroundImgID    int64       `json:"backgroundImgId"`
	BackgroundURL      string      `json:"backgroundUrl"`
	Authority          int         `json:"authority"`
	Mutual             bool        `json:"mutual"`
	DjStatus           int         `json:"djStatus"`
	VipType            int         `json:"vipType"`
	RemarkName         interface{} `json:"remarkName"`
	AvatarImgIDStr     string      `json:"avatarImgIdStr"`
	BackgroundImgIDStr string      `json:"backgroundImgIdStr"`
}

type DJInfo struct {
	UserInfo

	AuthenticationTypes int  `json:"authenticationTypes"`
	Anchor              bool `json:"anchor"`
}

type ExtendedUserInfoForSearch struct {
	UserInfo

	AuthenticationTypes int  `json:"authenticationTypes"`
	Anchor              bool `json:"anchor"`

	FollowedByCount int `json:"followeds"`
	FollowsCount    int `json:"follows"`

	Algorithm                 string `json:"alg"`
	PlaylistCount             int    `json:"playlistCount"`
	PlaylistBeSubscribedCount int    `json:"playlistBeSubscribedCount"`
	AvatarImgIDStr            string `json:"avatarImgId_str,omitempty"`
}

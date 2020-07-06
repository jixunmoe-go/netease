package types

type AlbumArtist struct {
	Img1V1ID    int64         `json:"img1v1Id"`
	TopicPerson int           `json:"topicPerson"`
	Followed    bool          `json:"followed"`
	Alias       []interface{} `json:"alias"`
	PicID       int           `json:"picId"`
	BriefDesc   string        `json:"briefDesc"`
	MusicSize   int           `json:"musicSize"`
	AlbumSize   int           `json:"albumSize"`
	Img1V1URL   string        `json:"img1v1Url"`
	Trans       string        `json:"trans"`
	PicURL      string        `json:"picUrl"`
	Name        string        `json:"name"`
	ID          int           `json:"id"`
	Img1V1IDStr string        `json:"img1v1Id_str"`
}

type AlbumCommentResourceInfo struct {
	ID        int         `json:"id"`
	UserID    int         `json:"userId"`
	Name      string      `json:"name"`
	ImgURL    string      `json:"imgUrl"`
	Creator   interface{} `json:"creator"`
	EncodedID interface{} `json:"encodedId"`
	SubTitle  interface{} `json:"subTitle"`
	WebURL    interface{} `json:"webUrl"`
}

type AlbumCommentThread struct {
	ID               string                   `json:"id"`
	ResourceInfo     AlbumCommentResourceInfo `json:"resourceInfo"`
	ResourceType     int                      `json:"resourceType"`
	CommentCount     int                      `json:"commentCount"`
	LikedCount       int                      `json:"likedCount"`
	ShareCount       int                      `json:"shareCount"`
	HotCount         int                      `json:"hotCount"`
	LatestLikedUsers interface{}              `json:"latestLikedUsers"`
	ResourceID       int                      `json:"resourceId"`
	ResourceTitle    string                   `json:"resourceTitle"`
	ResourceOwnerID  int                      `json:"resourceOwnerId"`
}

type AlbumSocialData struct {
	CommentThread    AlbumCommentThread `json:"commentThread"`
	LatestLikedUsers interface{}        `json:"latestLikedUsers"`
	Liked            bool               `json:"liked"`
	Comments         interface{}        `json:"comments"`
	ResourceType     int                `json:"resourceType"`
	ResourceID       int                `json:"resourceId"`
	CommentCount     int                `json:"commentCount"`
	LikedCount       int                `json:"likedCount"`
	ShareCount       int                `json:"shareCount"`
	ThreadID         string             `json:"threadId"`
}

type AlbumInfo struct {
	ID                   int             `json:"id"`
	Type                 string          `json:"type"`
	Size                 int             `json:"size"`
	Songs                []interface{}   `json:"songs"`
	Paid                 bool            `json:"paid"`
	OnSale               bool            `json:"onSale"`
	Mark                 int             `json:"mark"`
	Alias                []interface{}   `json:"alias"`
	Artists              []AlbumArtist   `json:"artists"`
	CopyrightID          int             `json:"copyrightId"`
	PicID                uint64          `json:"picId"`
	Artist               AlbumArtist     `json:"artist"`
	PublishTime          uint64          `json:"publishTime"`
	Company              string          `json:"company"`
	BriefDescription     string          `json:"briefDesc"`
	CommentThreadID      string          `json:"commentThreadId"`
	CoverImageURL        string          `json:"picUrl"`
	CoverImageID         uint64          `json:"pic"`
	CoverImageStringID   string          `json:"picId_str"`
	BlurredCoverImageURL string          `json:"blurPicUrl"`
	CompanyID            int             `json:"companyId"`
	Tags                 string          `json:"tags"`
	Description          string          `json:"description"`
	Status               int             `json:"status"`
	SubType              string          `json:"subType"`
	Name                 string          `json:"name"`
	SocialInfo           AlbumSocialData `json:"info"`
}

type AlbumResp struct {
	Songs    []SongInfo `json:"songs"`
	Code     int        `json:"code"`
	Album    AlbumInfo  `json:"album"`
	Response string     `json:"-"`
}

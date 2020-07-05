package types

type ArtistInfo struct {
	ID int `json:"id"`

	// Original Name
	Name string `json:"name"`
	// Localised name (Chinese)
	LocalisedName string `json:"trans"`

	// A brief description of this artist.
	Description string `json:"briefDesc"`

	Img1V1ID    int64  `json:"img1v1Id"`
	Img1V1IDStr string `json:"img1v1Id_str"`
	Img1V1URL   string `json:"img1v1Url"`

	PicID    int64  `json:"picId"`
	PicIDStr string `json:"picId_str"`
	PicURL   string `json:"picUrl"`

	TopicPerson int           `json:"topicPerson"`
	Alias       []interface{} `json:"alias"`

	// Number of songs
	SongCount int `json:"musicSize"`
	// Number of albums
	AlbumCount int `json:"albumSize"`
	// Number of MV
	MVCount int `json:"mvSize"`

	Followed bool `json:"followed"`

	PublishTime uint64 `json:"publishTime"`
	AccountID   uint64 `json:"accountId"`
}

type ArtistResp struct {
	Artist ArtistInfo `json:"artist"`

	// Popular songs
	Songs []SongInfo `json:"hotSongs"`

	// There's more songs flag?
	More bool `json:"more"`

	Code     int    `json:"code"`
	Response string `json:"-"`
}

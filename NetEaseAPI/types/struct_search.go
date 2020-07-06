package types

type BasicSongSearchResult struct {
	Songs     []SongInfo `json:"songs,omitempty"`
	SongCount uint       `json:"songCount,omitempty"`
}

type SongSearchResult struct {
	BasicSongSearchResult
	RecQuery   interface{} `json:"rec_query,omitempty"`
	RecType    interface{} `json:"rec_type,omitempty"`
	Highlights []string    `json:"highlights,omitempty"`
	HasMore    bool        `json:"hasMore,omitempty"`
}

type ArtistSearchResult struct {
	ArtistCount uint              `json:"artistCount,omitempty"`
	Artists     []BasicArtistInfo `json:"artists,omitempty"`
}

type AlbumSearchResult struct {
	Albums     []Album `json:"albums,omitempty"`
	AlbumCount uint    `json:"albumCount,omitempty"`
}

type VideoSearchResult struct {
	VideoCount uint        `json:"videoCount,omitempty"`
	Videos     []VideoInfo `json:"videos,omitempty"`
}

// a subset of `SongSearchResult`
type LyricSearchResult BasicSongSearchResult

type PlaylistSearchResult struct {
	Playlists     []BasicPlaylistInfo `json:"playlists,omitempty"`
	PlaylistCount uint                `json:"playlistCount,omitempty"`
}

type RadioSearchResult struct {
	Radios      []BasicRadioInfo `json:"djRadios,omitempty"`
	RadiosCount uint             `json:"djRadiosCount,omitempty"`
}

type UserSearchResult struct {
	Users      []ExtendedUserInfoForSearch `json:"userprofiles"`
	UsersCount uint                        `json:"userprofileCount"`
}

type SearchResultUnion struct {
	SongSearchResult
	ArtistSearchResult
	AlbumSearchResult
	VideoSearchResult
	LyricSearchResult
	PlaylistSearchResult
	RadioSearchResult
	UserSearchResult
}

type SearchResp struct {
	Result SearchResultUnion `json:"result"`

	Code     int    `json:"code"`
	Response string `json:"-"`
}

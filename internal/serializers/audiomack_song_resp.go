package serializers

type AudiomackSongResp struct {
	Id       int                    `json:"id"`
	Type     string                 `json:"type"`
	Released string                 `json:"released"`
	Artist   string                 `json:"artist"`
	Cover    string                 `json:"image_base"`
	Title    string                 `json:"title"`
	Stats    AudiomackSongRespStats `json:"stats"`
}

type AudiomackSongRespStats struct {
	Plays     int `json:"plays-raw"`
	Favorites int `json:"favorites-raw"`
}

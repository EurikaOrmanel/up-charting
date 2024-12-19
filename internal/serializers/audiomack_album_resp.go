package serializers

type AudiomackAlbumResp struct {
	Id       int                    `json:"id"`
	Type     string                 `json:"type"`
	Released string                 `json:"released"`
	Artist   string                 `json:"artist"`
	Cover    string                 `json:"image_base"`
	Title    string                 `json:"title"`
	Stats    AudiomackSongRespStats `json:"stats"`
}

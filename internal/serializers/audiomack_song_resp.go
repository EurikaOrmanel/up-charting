package serializers

import (
	"strconv"
	"time"
)

type AudiomackSongResp struct {
	Id       int                    `json:"id"`
	Type     string                 `json:"type"`
	Released string                 `json:"released"`
	Artist   string                 `json:"artist"`
	Cover    string                 `json:"image_base"`
	Title    string                 `json:"title"`
	Stats    AudiomackSongRespStats `json:"stats"`
}

func (songRsp AudiomackSongResp) GetReleasedTime() *time.Time {
	releasedInt, err := strconv.Atoi(songRsp.Released)
	if err != nil {
		return nil
	}
	released := time.UnixMilli(int64(releasedInt))
	return &released
}

type AudiomackSongRespStats struct {
	Plays     int `json:"plays-raw"`
	Favorites int `json:"favorites-raw"`
}

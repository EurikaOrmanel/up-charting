package services

import (
	"EurikaOrmanel/up-charter/internal/serializers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func AudiomackSongByURL(url string) {}

func AudiomackAlbumInfo(albumUrl string) (serializers.AudiomackAlbumResp, error) {
	client := http.Client{}
	AUDIOMACK_API_BASE_URL := os.Getenv("AUDIOMACK_API_BASE_URL")

	searchUrl, err := url.Parse(albumUrl)
	if err != nil {
		return serializers.AudiomackAlbumResp{}, err
	}
	finalUrl := fmt.Sprintf("%s/audiomack/album?url=%s", AUDIOMACK_API_BASE_URL, searchUrl.String())
	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header = genApiHeader()
	if err != nil {
		return serializers.AudiomackAlbumResp{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return serializers.AudiomackAlbumResp{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return serializers.AudiomackAlbumResp{}, err
	}

	albumResp := new(serializers.AudiomackAlbumResp)
	err = json.Unmarshal(body, albumResp)
	if err != nil {
		return serializers.AudiomackAlbumResp{}, err
	}
	return *albumResp, err

}

func AudiomackSongInfo(songUrl string) (serializers.AudiomackSongResp, error) {
	client := http.Client{}
	AUDIOMACK_API_BASE_URL := os.Getenv("AUDIOMACK_API_BASE_URL")

	searchsongUrl, err := url.Parse(songUrl)
	finalUrl := fmt.Sprintf("%s/audiomack/song?url=%s", AUDIOMACK_API_BASE_URL, searchsongUrl.String())
	if err != nil {

		return serializers.AudiomackSongResp{}, err
	}
	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header = genApiHeader()
	if err != nil {
		return serializers.AudiomackSongResp{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return serializers.AudiomackSongResp{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return serializers.AudiomackSongResp{}, err
	}

	songResp := new(serializers.AudiomackSongResp)
	err = json.Unmarshal(body, songResp)
	if err != nil {
		return serializers.AudiomackSongResp{}, err
	}
	return *songResp, err
}

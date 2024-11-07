package schemas

import (
	"EurikaOrmanel/up-charter/internal/models"

	"github.com/google/uuid"
)

type AddSongInput struct {
	Title     string             `validate:"string|required" json:"title"`
	Platforms SongPlatformsInput `validate:"required|list" json:"platforms"`
	ArtistId  string             `validate:"uuid4|required" json:"artist_id"`
	GenreId   string             `validate:"required|uuid4" json:"genre_id"`
}

type SongPlatformsInput []struct {
	PlatformId string `validate:"uuid4|required" json:"platform_id"`
	Url        string `validate:"fullurl|required" json:"url"`
}

func (platformINpus SongPlatformsInput) ToSongPlatform() []models.SongPlatform {
	songPlatforms := []models.SongPlatform{}
	for _, snPlt := range platformINpus {
		songPlatforms = append(songPlatforms, models.SongPlatform{
			// SongID:     uuid.MustParse(snPlt.SongId),
			PlatformID: uuid.MustParse(snPlt.PlatformId),
			Url:        snPlt.Url,
		})
	}
	return songPlatforms
}
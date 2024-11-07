package schemas

import (
	"EurikaOrmanel/up-charter/internal/models"

	"github.com/google/uuid"
)

type ArtistPlatformInput struct {
	ArtistId string `validate:"required|uuid4" json:"artist_id"`
	ArtistPlatformWOArtInput
}

type ArtistPlatformWOArtInput struct {
	PlatformId string `validate:"required|uuid4" json:"platform_id"`
	Url        string `validate:"required|fullUrl" json:"url"`
}
type ArtistPlatformsInput struct {
	ArtistPlatforms []ArtistPlatformInput `json:"artist_platforms"`
}

type ArtistPlatformInputs []ArtistPlatformWOArtInput

func (artPlats ArtistPlatformInputs) ToArtistPlatform() []models.ArtistPlatform {
	artistPlatforms := []models.ArtistPlatform{}
	for _, artistPlatformInput := range artPlats {
		artistPlatforms = append(artistPlatforms, models.ArtistPlatform{
			Url:        artistPlatformInput.Url,
			PlatformID: uuid.MustParse(artistPlatformInput.PlatformId)})
	}

	return artistPlatforms
}

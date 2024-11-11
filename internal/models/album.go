package models

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	Base

	Title string `json:"title"`

	Cover string `json:"cover"`

	ArtistID uuid.UUID
	Artist   *Artist `json:"artist"`

	Songs []Song `json:"songs"`

	Platforms []AlbumPlatform `json:"platforms"`

	PlayCounts []AlbumPlayCount
}

func (adM *Album) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

type AlbumPlatforms []AlbumPlatform

func (alPlats AlbumPlatforms) FindAlbumPlatformByName(name string) *AlbumPlatform {
	for _, snPlat := range alPlats {
		if strings.Contains(strings.ToLower(snPlat.Url), strings.ToLower(name)) {
			return &snPlat
		}
	}
	return nil
}

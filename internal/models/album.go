package models

import (
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
}

func (adM *Album) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

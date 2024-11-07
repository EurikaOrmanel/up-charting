package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	Base

	Title string

	Cover string

	ArtistID uuid.UUID
	Artist *Artist

	Songs []Song

	Platforms []AlbumPlatform `json:"platforms"`

}

func (adM *Album) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

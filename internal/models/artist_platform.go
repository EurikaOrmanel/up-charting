package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArtistPlatform struct {
	Base

	ArtistID uuid.UUID

	Url        string
	PlatformID uuid.UUID
	Platform   *Platform
}

func (adM *ArtistPlatform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

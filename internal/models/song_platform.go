package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SongPlatform struct {
	Base

	SongID uuid.UUID

	PlatformID uuid.UUID
	Platform   Platform
	Url        string //the unique ID the current album bears on the specified platform
}

func (adM *SongPlatform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

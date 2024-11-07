package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumPlatform struct {
	Base
	AlbumID uuid.UUID

	PlatformID uuid.UUID
	Platform   *Platform

	Uid string //the unique ID the current album bears on the specified platform
}

func (adM *AlbumPlatform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

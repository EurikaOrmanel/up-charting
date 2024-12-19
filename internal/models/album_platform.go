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

	Url string 
}

func (adM *AlbumPlatform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

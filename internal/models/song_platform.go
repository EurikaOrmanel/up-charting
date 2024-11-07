package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type SongPlatform struct {
	Base

	SongID int
	Song   Song

	PlatformID int
	Platform   Platform
	Uid        string //the unique ID the current album bears on the specified platform
}


func (adM *SongPlatform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}
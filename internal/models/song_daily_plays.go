package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SongDailyPlay struct {
	Base
	Count int

	SongID uuid.UUID

	PlatformID uuid.UUID

	Platform *Platform
}

func (adM *SongDailyPlay) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

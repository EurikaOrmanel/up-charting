package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SongDailyPlay struct {
	Base
	Count int

	SongID uuid.UUID

	SongPlatformID uuid.UUID `json:"song_platform_id"`

	SongPlatform *SongPlatform `json:"platform,omitempty"`
}

func (adM *SongDailyPlay) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

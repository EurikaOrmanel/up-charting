package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SongDailyPlay struct {
	Base
	Count int

	SongID uuid.UUID

	PlatformID uuid.UUID `json:"platform_id"`

	Platform *Platform `json:"platform,omitempty"`
}

func (adM *SongDailyPlay) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

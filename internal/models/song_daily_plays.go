package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type SongDailyPlay struct {
	Base
	Count int

	SongID uuid.UUID
	Song   *Song
}



func (adM *SongDailyPlay) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}
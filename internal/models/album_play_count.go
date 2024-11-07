package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumPlayCount struct {
	Base

	Count   int
	AlbumID uuid.UUID
	Album   *Album
}

func (adM *AlbumPlayCount) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()
	return nil
}

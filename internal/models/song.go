package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	Base

	ArtistID int
	Artist   *Artist

	Title string

	ReleasedDate time.Time

	GenreID int
	Genre   Genre

	AlbumID int
}

func (adM *Song) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

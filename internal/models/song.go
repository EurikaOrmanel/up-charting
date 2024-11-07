package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	Base

	ArtistID uuid.UUID
	Artist   *Artist

	Title string

	ReleasedDate time.Time

	GenreID uuid.UUID
	Genre   Genre

	AlbumID int

	Platforms []SongPlatform `json:"platforms"`
}

func (adM *Song) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

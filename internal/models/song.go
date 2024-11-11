package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	Base

	ArtistID uuid.UUID `json:"artist_id"`
	Artist   *Artist `json:"artist,omitempty"`

	Title string

	Cover string

	ReleasedDate time.Time

	GenreID uuid.UUID
	Genre   Genre

	AlbumID int

	Platforms SongPlatforms `json:"platforms"`

	PlayCounts []SongDailyPlay
}

type SongPlatforms []SongPlatform

func (snPlats SongPlatforms) FindSongPlatformByName(name string) *SongPlatform {
	for _, snPlat := range snPlats {
		if strings.Contains(strings.ToLower(snPlat.Url), strings.ToLower(name)) {
			return &snPlat
		}
	}
	return nil
}

func (adM *Song) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

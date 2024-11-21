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
	Artist   *Artist   `json:"artist,omitempty"`

	Title string `json:"title"`

	Cover string `json:"cover"`

	ReleasedDate time.Time `json:"released_date,omitempty"`

	GenreID uuid.UUID `json:"genre_id,omitempty"`
	Genre   *Genre    `json:"genre,omitempty"`

	AlbumID int `json:"album_id"`

	Platforms *SongPlatforms `json:"platforms,omitempty"`

	PlayCounts []SongDailyPlay `json:"play_counts"`
}
type SongNCount struct {
	Song
	TotalCount int `json:"total_count"`
}
type SongPlatforms []*SongPlatform

func (snPlats SongPlatforms) FindSongPlatformByName(name string) *SongPlatform {
	for _, snPlat := range snPlats {
		if strings.Contains(strings.ToLower(snPlat.Url), strings.ToLower(name)) {
			return snPlat
		}
	}
	return nil
}

func (adM *Song) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}

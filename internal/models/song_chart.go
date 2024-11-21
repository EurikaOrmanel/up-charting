package models

import "github.com/google/uuid"

type Top100Chart struct {
	Base
	Position         int       `gorm:"unique" json:"position,omitempty"`
	PreviousPosition int       `json:"previous_position,omitempty"`
	SongID           uuid.UUID `gorm:"unique" json:"song_id,omitempty"`
	Song             *Song     `json:"song,omitempty"`
	GenreID          uuid.UUID `json:"genre_id,omitempty"`
	Genre            *Genre    `json:"genre,omitempty"`
}
type Top100WPlayCount struct {
	Top100Chart
	TotalCount int
}

type Top100Charts []Top100Chart

func (topCharts Top100Charts) GetIdStrings() []string {
	ids := []string{}
	for _, topChart := range topCharts {
		ids = append(ids, topChart.ID.String())
	}
	return ids
}

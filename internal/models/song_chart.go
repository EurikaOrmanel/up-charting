package models

import "github.com/google/uuid"

type Top100Chart struct {
	Base
	Position         int `gorm:"unique"`
	PreviousPosition int
	SongID           uuid.UUID `gorm:"unique"`
	Song             *Song
	GenreID          uuid.UUID
	Genre            *Genre
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

package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (db DB) GetLastItemInChart() models.Top100WPlayCount {
	songChart := new(models.Top100WPlayCount)
	db.
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("INNER JOIN song_platforms ON song_platforms.id = song_daily_plays.song_platform_id").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Group("top100_charts.song_id").
		Order("position DESC").
		Limit(1).
		Find(&songChart)
	return *songChart
}
func (db DB) GetChartNSongPlayCountGrtrThanCurrentPosition(
	playCount int,
	query schemas.PaginationQuery,
) []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Scopes(paginate(query, db.DB)).
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Having("total_count > ?", playCount).
		Joins("INNER JOIN song_platforms ON song_platforms.id = song_daily_plays.song_platform_id").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Group("top100_charts.song_id").
		Order("position DESC").
		Find(&songsFromPosition)
	return songsFromPosition
}
func (db DB) GetSongChart(query schemas.SongChartQuery) []*models.Top100Chart {
	charts := []*models.Top100Chart{}
	db.Preload("song").
		Scopes(paginate(query.PaginationQuery, db.DB)).Model(charts)
	return charts
}

func (db DB) UpdateChartPosition(songChart *models.Top100Chart) error {
	currentSong := db.SongNChartByID(songChart.SongID.String())
	if currentSong.ID == uuid.Nil {
		return errors.New("song not found")
	}
	top100Counts := db.GetChartNSongPlayCountGrtrThanCurrentPosition(currentSong.TotalCount, schemas.PaginationQuery{Count: 5, Page: 1})

	if len(top100Counts) == 0 {
		return nil
	}
	fmt.Println("currentSong.TotalPlay:", currentSong.TotalCount)
	fmt.Println("top100Counts[0].TotalPlay:", top100Counts[0].TotalCount)

	if currentSong.TotalCount > top100Counts[0].TotalCount {
		db.ShiftDownPosition(top100Counts[0].Position, top100Counts[0].Position+1)
		songChart.PreviousPosition = songChart.Position
		songChart.Position = top100Counts[0].Position
		return db.UpdateTop100Chart(songChart)
	}
	return nil
}

func (db DB) AddSongToChart(songChart *models.Top100Chart) error {
	currentSong := db.SongNChartByID(songChart.SongID.String())
	if currentSong.ID == uuid.Nil {
		return errors.New("song not found")
	}
	top100Counts := db.GetChartNSongPlayCountGrtrThanCurrentPosition(currentSong.TotalCount, schemas.PaginationQuery{Count: 5, Page: 1})
	if len(top100Counts) == 0 {
		lastSongInChart := db.GetLastItemInChart()
		fmt.Println("lastSongInChart.TotalCount:", lastSongInChart.TotalCount, " lastSongInChart.Position:", lastSongInChart.Position)
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaadddddddddddddddddddddd new chart with position as 1")
		if lastSongInChart.ID == uuid.Nil {
			songChart.Position = 1

		} else {
			if lastSongInChart.TotalCount < currentSong.TotalCount {
				db.ShiftDownPosition(lastSongInChart.Position, lastSongInChart.Position+1)
				songChart.Position = lastSongInChart.Position
			} else {
				songChart.Position = lastSongInChart.Position + 1
			}
		}
		return db.Create(songChart).Error
	}

	db.ShiftDownPosition(top100Counts[0].Position, top100Counts[0].Position+1)
	songChart.Position = top100Counts[0].Position

	return db.Create(songChart).Error
}
func (db DB) GetCharts(query schemas.PaginationQuery) models.Top100Charts {
	chartsNSongs := models.Top100Charts{}

	db.
		Scopes(paginate(query, db.DB)).
		Preload("song").Find(&chartsNSongs)
	return chartsNSongs

}
func (db DB) GetFromChartBySongID(songId string) models.Top100Chart {
	chart := new(models.Top100Chart)
	db.First(chart, "song_id = ?", songId)
	return *chart
}

func (db DB) ShiftUpPosition(currentPosition int, newPosition int) error {
	return db.Exec("UPDATE top100_charts SET position = position - 1,previous_position = position WHERE position < ? AND position >= ?", currentPosition, newPosition).Error

}

func (db DB) ShiftDownPosition(currentPosition int, newPosition int) error {
	return db.Exec("UPDATE top100_charts SET position = position + 1,previous_position = position WHERE  position >= ? AND position < ? ", currentPosition, newPosition).Error

}

func (db DB) UpdateTop100Chart(chart *models.Top100Chart) error {
	return db.Model(models.Top100Chart{}).Where("id = ?", chart.ID).Updates(*chart).Error
}

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
		Select("top100_charts.*, COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("position DESC").
		Limit(1).
		Find(songChart)
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
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("total_count DESC").
		Find(&songsFromPosition)
	return songsFromPosition
}

func (db DB) GetAllInChartInPositionOrder(
	playCount int,
	query schemas.PaginationQuery,
) []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Scopes(paginate(query, db.DB)).
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Having("total_count > ?", playCount).
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("total_count DESC").
		Find(&songsFromPosition)
	return songsFromPosition
}
func (db DB) GetChartNSongPlayCountCurrentPosition(
	playCount int,
	query schemas.PaginationQuery,
) []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Scopes(paginate(query, db.DB)).
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("position DESC").
		Find(&songsFromPosition)
	return songsFromPosition
}

func (db DB) GetFullChartSortedByPosition() []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("position DESC").
		Find(&songsFromPosition)
	return songsFromPosition
}

func (db DB) GetChartNSongPlayCountLTCurrentPosition(
	playCount int,
	query schemas.PaginationQuery,
) []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Scopes(paginate(query, db.DB)).
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Having("total_count < ?", playCount).
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("total_count DESC").
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
	//TODO:updates will be done differently where all songs
	currentSong := db.SongNChartByID(songChart.SongID.String())
	if currentSong.ID == uuid.Nil {
		return errors.New("song not found")
	}
	top100Counts := db.GetChartNSongPlayCountLTCurrentPosition(currentSong.TotalCount, schemas.PaginationQuery{Count: 5, Page: 1})

	if len(top100Counts) == 0 {
		return nil
	}
	fmt.Println("currentSong.TotalPlay:", currentSong.TotalCount)
	fmt.Println("top100Counts[0].TotalPlay:", top100Counts[0].TotalCount)
	if len(top100Counts) > 0 {
		if currentSong.TotalCount > top100Counts[0].TotalCount {

			currentPosition := db.GetFromChartBySongID(currentSong.ID.String())
			db.ShiftDownPosition(currentPosition.Position, top100Counts[0].Position)
			songChart.PreviousPosition = songChart.Position
			songChart.Position = top100Counts[0].Position
			return db.UpdateTop100Chart(songChart)
		}
	}
	return nil
}

func (db DB) AddSongToChart(songChart *models.Top100Chart) error {
	// Fetch the song's current play count
	currentSong := db.SongNChartByID(songChart.SongID.String())
	if currentSong.ID == uuid.Nil {
		return errors.New("song not found")
	}

	// Start a transaction for atomic updates
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the entire chart sorted by position
	chart := db.GetFullChartSortedByPosition()
	newPosition := -1

	// Determine the new position based on TotalCount
	for _, chartPos := range chart {
		if currentSong.TotalCount > chartPos.TotalCount {
			newPosition = chartPos.Position
			break
		}
	}

	if newPosition == -1 {
		// If the song has the lowest count or the chart is empty
		lastChart := db.GetLastItemInChart()
		if lastChart.ID == uuid.Nil {
			// Empty chart, set position to 1
			newPosition = 1
		} else {
			// Append to the end of the chart
			newPosition = lastChart.Position + 1
		}
	} else {
		// Shift down all songs from the new position
		if err := db.ShiftDownPosition(newPosition, len(chart)+1); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Assign the calculated position to the new song
	songChart.Position = newPosition
	if err := tx.Create(songChart).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

// func (db DB) AddSongToChart(songChart *models.Top100Chart) error {
// 	currentSong := db.SongNChartByID(songChart.SongID.String())
// 	if currentSong.ID == uuid.Nil {
// 		return errors.New("song not found")
// 	}
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 		  tx.Rollback()
// 		}
// 	  }()
// 	chart := db.GetFullChartSortedByPosition()
// 	newPosition := -1
// 	for _, chartPos := range chart {
// 		if chartPos.TotalCount < currentSong.TotalCount {
// 			newPosition = chartPos.Position
// 			break
// 		}
// 	}

// 	if newPosition == -1 {
// 		lastChart := db.GetLastItemInChart()
// 		newPosition = lastChart.Position
// 		if lastChart.ID == uuid.Nil {
// 			newPosition = 1
// 		}else{

// 		}
// 	} else {
// 		db.ShiftDownPosition(newPosition, len(chart)+1)

// 	}
// 	songChart.Position = newPosition
// 	if err := tx.Create(songChart).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit().Error
// }

// func (db DB) AddSongToChart(songChart *models.Top100Chart) error {
// 	// Validate the song exists
// 	currentSong := db.SongNChartByID(songChart.SongID.String())
// 	if currentSong.ID == uuid.Nil {
// 		return errors.New("song not found")
// 	}

// 	// Fetch last song and top100 counts
// 	lastSongInChart := db.GetLastItemInChart()
// 	top100Counts := db.GetChartNSongPlayCountGrtrThanCurrentPosition(currentSong.TotalCount, schemas.PaginationQuery{Count: 5, Page: 1})

// 	var newPosition int
// 	if len(top100Counts) == 0 {
// 		if lastSongInChart.ID == uuid.Nil {
// 			newPosition = 1 // Empty chart
// 		} else if lastSongInChart.TotalCount > currentSong.TotalCount {
// 			newPosition = lastSongInChart.Position + 1 // Add at the end
// 		} else {
// 			db.ShiftDownPosition(1, lastSongInChart.Position+1)
// 			newPosition = 1 // Add at the top
// 		}
// 	} else {
// 		db.ShiftDownPosition(top100Counts[0].Position, lastSongInChart.Position+1)
// 		newPosition = top100Counts[0].Position // Add in between
// 	}

// 	songChart.Position = newPosition
// 	return db.Create(songChart).Error
// }

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

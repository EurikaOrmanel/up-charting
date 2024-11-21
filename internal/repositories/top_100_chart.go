package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (db DB) GetChartyBySongID(songID string) models.Top100WPlayCount {
	songChart := new(models.Top100WPlayCount)
	db.
		Table("top100_charts").
		Select("top100_charts.*, COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Where("top100_charts.song_id = ?", songID).
		Scan(songChart)
	return *songChart
}

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

func (db DB) GetFullChartSortedByTotalCount() []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("total_count DESC").
		Scan(&songsFromPosition)
	return songsFromPosition
}
func (db DB) GetTop100Chart() []models.Top100Chart {
	chart := []models.Top100Chart{}
	db.Find(&chart).Order("position ASC")
	return chart
}
func (db DB) GetFullChartSortedByPostion() []models.Top100WPlayCount {
	songsFromPosition := []models.Top100WPlayCount{}
	db.
		Table("top100_charts").
		Select("top100_charts.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = top100_charts.song_id").
		Group("top100_charts.song_id").
		Order("position ASC").
		Find(&songsFromPosition)
	return songsFromPosition
}

func (db DB) GetChart100NSong() []models.Top100Chart {
	top100 := []models.Top100Chart{}
	err := db.Preload("Song").
		Order("position ASC").
		Limit(100).Find(&top100).Error
	if err != nil {
		fmt.Println(err)
	}
	return top100
}

func (db DB) GetChartNSongPlayCountLTCurrentPosition(playCount int,
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
	// Fetch the current song from the chart
	currentPosition := db.GetChartyBySongID(songChart.SongID.String())
	if currentPosition.ID == uuid.Nil {
		return errors.New("song not found in chart")
	}

	// Start a transaction for atomic updates
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the target position (sorted by TotalCount in descending order)
	chart := db.GetFullChartSortedByTotalCount()

	// Identify the song at the position where the swap needs to occur
	targetSong := models.Top100Chart{}
	for _, chartEntry := range chart {
		if currentPosition.TotalCount > chartEntry.TotalCount &&
			currentPosition.Position > chartEntry.Position {
			targetSong = chartEntry.Top100Chart
			break
		}
	}

	if targetSong.ID == uuid.Nil {
		// No swapping needed; the song is already in the correct position
		return nil
	}

	// Swap positions
	currentPosition.PreviousPosition = currentPosition.Position
	targetSong.PreviousPosition = targetSong.Position

	// Perform the position switch
	tempPosition := currentPosition.Position
	currentPosition.Position = targetSong.Position
	targetSong.Position = tempPosition

	// Update the database with the swapped positions
	if err := tx.Save(&currentPosition).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&targetSong).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

// func (db DB) UpdateChartPosition(songChart *models.Top100Chart) error {
// 	//TODO:updates will be done differently where all songs
// 	currentSong := db.SongNChartByID(songChart.SongID.String())
// 	if currentSong.ID == uuid.Nil {
// 		return errors.New("song not found")
// 	}
// 	currentPosition := db.GetChartyBySongID(songChart.SongID.String())
// 	chart := db.GetFullChartSortedByTotalCount()
// 	toUpdate := models.Top100Charts{}
// 	for i := currentPosition.Position - 1; i <= len(chart); i++ {
// 		if currentPosition.TotalCount > chart[i].TotalCount &&
// 			currentPosition.Position > chart[i].Position {

// 			chart[i].Position = currentPosition.Position
// 			currentPosition.Position = chart[i].Position
// 		}
// 	}
// 	db.UpdateTop100Charts(&toUpdate)

//		return nil
//	}
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
	chart := db.GetFullChartSortedByTotalCount()
	newPosition := -1

	// Determine the new position based on TotalCount
	for _, chartPos := range chart {
		if currentSong.TotalCount > chartPos.TotalCount {
			newPosition = chartPos.Position
			fmt.Println("Chart position assigned to existing one:", newPosition)
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
		chartsByPosition := db.GetFullChartSortedByPostion()
		if err := db.ShiftDownPosition(newPosition, chartsByPosition[len(chartsByPosition)-1].Position+1); err != nil {
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
func (db DB) GetChartsByPositionGTCurrent(currentPosition int) models.Top100Charts {
	chartsNSongs := models.Top100Charts{}
	db.
		Preload("top100_charts").Where("position >= ?", currentPosition).
		Order("position ASC").
		Find(&chartsNSongs)
	return chartsNSongs

}

func (db DB) ShiftUpPosition(currentPosition int, newPosition int) error {
	return db.Exec("UPDATE top100_charts SET position = position - 1,previous_position = position WHERE position < ? AND position >= ?", currentPosition, newPosition).Error

}

// func (db DB) ShiftDownPosition(currentPosition int, newPosition int) error {
// 	chartsGTOrE := db.GetChartsByPositionGTCurrent(currentPosition)
// 	for _, currentChart := range chartsGTOrE {
// 		if currentChart.Position >= currentPosition && currentChart.Position < newPosition {
// 			currentChart.Position = currentChart.Position + 1
// 			currentChart.PreviousPosition = currentChart.Position
// 			db.UpdateTop100Chart(&currentChart)
// 		}
// 	}
// 	// return db.Exec("UPDATE top100_charts SET position = position + 1,previous_position = position WHERE  position >= ? AND position < ? ", currentPosition, newPosition).Error
// 	return nil
// }

// func (db DB) ShiftDownPosition(currentPosition int, newPosition int) error {
// 	// Ensure newPosition is greater than currentPosition
// 	if newPosition <= currentPosition {
// 		return fmt.Errorf("newPosition must be greater than currentPosition")
// 	}

// 	// Perform the update in descending order to avoid conflicts
// 	query := `
//         UPDATE top100_charts
//         SET position = position + 1,
//             previous_position = position
//         WHERE position >= ? AND position < ?
//         ORDER BY position DESC;`

// 	if err := db.Exec(query, currentPosition, newPosition).Error; err != nil {
// 		return fmt.Errorf("failed to shift down positions: %w", err)
// 	}
// 	return nil
// }

func (db DB) ShiftDownPosition(currentPosition int, newPosition int) error {
	// Ensure newPosition is greater than currentPosition
	if newPosition <= currentPosition {
		return fmt.Errorf("newPosition must be greater than currentPosition")
	}

	// Fetch rows to update in descending order of position
	rows := []models.Top100Chart{}
	if err := db.Where("position >= ? AND position < ?", currentPosition, newPosition).
		Order("position DESC").
		Find(&rows).Error; err != nil {
		return fmt.Errorf("failed to fetch rows: %w", err)
	}

	// Update each row one at a time
	for _, row := range rows {
		row.PreviousPosition = row.Position
		row.Position = row.Position + 1
		if err := db.Save(&row).Error; err != nil {
			return fmt.Errorf("failed to update row ID %d: %w", row.ID, err)
		}
	}

	return nil
}
func (db DB) UpdateChart(songChart *models.Top100Chart) error {
	return db.Model(models.Top100Chart{}).Where("id = ?", songChart.ID).Updates(songChart).Error
}

func (db DB) UpdateTop100Chart(chart *models.Top100Chart) error {
	return db.Model(models.Top100Chart{}).Where("id = ?", chart.ID).Updates(*chart).Error
}

func (db DB) UpdateTop100Charts(chart *models.Top100Charts) error {
	return db.Model(models.Top100Chart{}).Where("id IN (?)", chart.GetIdStrings()).Updates(*chart).Error
}

package repositories

import "EurikaOrmanel/up-charter/internal/models"

// func (db DB) AddSongPlayCount(songPlay *models.SongDailyPlay)error{

// }

func (db DB) AddSongDailyPlays(songPlays []*models.SongDailyPlay) error {
	return db.Create(songPlays).Error
}

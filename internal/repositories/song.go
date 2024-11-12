package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
)

func (db DB) AddSong(song *models.Song) error {
	return db.Create(song).Error
}

func (db DB) FindSongByArtistIdNTitle(artistId string, title string) models.Song {
	song := new(models.Song)

	db.First(song, "artist_id = ? AND title = ?", artistId, title)
	return *song

}

// func (db DB) SongPlatformNChart(query schemas.PaginationQuery) {
// 	songs := []models.Song{}
// 	db.Preload("platforms").Preload("play_counts").Table("song")
// }

func (db DB) SongsTotalChartNLastChartDate(query schemas.PaginationQuery) {
	//as the name reads, that's what we're gon do.
	
}

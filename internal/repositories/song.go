package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
)

func (db DB) AddSong(song *models.Song) error {
	return db.Create(song).Error
}

func (db DB) FindSongByArtistIdNTitle(artistId string, title string) models.Song {
	song := new(models.Song)

	db.First(song, "artist_id = ? AND title = ?", artistId, title)
	return *song

}

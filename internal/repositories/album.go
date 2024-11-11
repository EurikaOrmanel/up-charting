package repositories

import "EurikaOrmanel/up-charter/internal/models"

func (db DB) AddAlbum(song *models.Album) error {
	return db.Create(song).Error
}

func (db DB) FindAlbumByArtistIdNTitle(artistId string, title string) models.Album {
	song := new(models.Album)

	db.First(song, "artist_id = ? AND title = ?", artistId, title)
	return *song

}

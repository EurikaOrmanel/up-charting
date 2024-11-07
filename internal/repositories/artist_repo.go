package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
)

func (db DB) FindArtistByStagename(name string) *models.Artist {
	artist := new(models.Artist)
	db.First(artist, "stagename = ?", name)
	return artist
}

func (db DB) CreateArtist(artist *models.Artist) error {
	return db.Create(artist).Error
}

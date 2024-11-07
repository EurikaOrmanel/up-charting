package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"

	"strings"
)

func (db DB) FindGenreByName(name string) models.Genre {
	genre := new(models.Genre)

	db.First(genre, "LOWER(name) = ?", strings.ToLower(name))
	return *genre
}

func (db DB) CreateGenre(genre *models.Genre) error {
	return db.Create(genre).Error
}

func (db DB) Genres(query schemas.PaginationQuery) ([]models.Genre, error) {
	genres := make([]models.Genre, 0)
	err := db.Scopes(paginate(query, db.DB)).Find(genres).Error
	return genres, err
}



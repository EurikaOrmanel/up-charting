package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
)

func (db DB) FindPlatformByName(name string) models.Platform {
	platform := new(models.Platform)

	db.First(platform, "name = ?", name)
	return *platform
}

func (db DB) CreatePlatform(platform *models.Platform) error {
	return db.Create(platform).Error
}

func (db DB) FindPlatforms(query schemas.PlatformQuery) ([]models.Platform, error) {
	platforms := make([]models.Platform, 0)
	err := db.Scopes(paginate(query.PaginationQuery, db.DB)).Find(&platforms, "name LIKE @q OR url LIKE @q ", map[string]any{"q": "%" + query.Q + "%"}).Error
	return platforms, err
}

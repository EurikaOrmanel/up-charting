package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
)

func (db DB) FindPlatformByName(name string) models.Platform {
	platform := new(models.Platform)

	db.First(platform, "name = ?", name)
	return *platform
}

func (db DB) CreatePlatform(platform *models.Platform) error {
	return db.Create(platform).Error
}

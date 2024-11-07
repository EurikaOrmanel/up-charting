package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"fmt"
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {
	DB_NAME := os.Getenv("DB_NAME")
	DB_PATH := os.Getenv("DB_DIR")
	ensureBaseDir(DB_PATH)
	// Correct the DSN format for PostgreSQL
	dbUrl := fmt.Sprintf("%s/%s", DB_PATH, DB_NAME)
	fmt.Println(dbUrl)
	// Connect to the database using GORM
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	return db

}

func (repoDb DB) MigrateAll() {
	repoDb.AutoMigrate(&models.Album{}, &models.AlbumPlatform{}, &models.Platform{}, &models.AlbumPlayCount{})
	repoDb.AutoMigrate(&models.Genre{}, &models.Song{}, &models.SongPlatform{})
	repoDb.AutoMigrate(&models.Admin{})
	repoDb.AutoMigrate(&models.Artist{}, &models.ArtistPlatform{}, &models.SongDailyPlay{})
}

func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

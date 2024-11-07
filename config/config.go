package config

import (
	"EurikaOrmanel/up-charter/internal/cache"
	"EurikaOrmanel/up-charter/internal/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gookit/validate"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	RepoDb      repositories.DB
	CacheConfig cache.CacheConfig
}

// func validator() {
// 	validate.AddValidator("ghPhone", validators.GHPhoneValidator)
// }

func Config() AppConfig {
	// validator()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	gormDB := repositories.InitDb()

	repoDb := repositories.DB{DB: gormDB}

	cacheConfig := cache.Config()
	repoDb.MigrateAll()

	return AppConfig{
		RepoDb:      repoDb,
		CacheConfig: cacheConfig,
	}

}

func InjectAppConfig(appConfig AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("appConfig", appConfig)
		return c.Next()
	}

}

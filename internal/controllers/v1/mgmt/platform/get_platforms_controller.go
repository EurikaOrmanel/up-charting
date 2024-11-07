package platform

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
)

func GetPlatformsController(c *fiber.Ctx) error {

	query := new(schemas.PlatformQuery)
	c.QueryParser(query)
	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	genres, err := repoDb.FindPlatforms(*query)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	return c.Status(fiber.StatusOK).JSON(genres)

}

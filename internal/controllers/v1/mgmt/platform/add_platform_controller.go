package platform

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddPlatformController(c *fiber.Ctx) error {
	body := new(schemas.AddPlatformInput)
	c.BodyParser(body)

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb

	foundPlatform := repoDb.FindPlatformByName(body.Name)
	if foundPlatform.ID != uuid.Nil {

		errResp.Message = "A platform with provided name already exists"
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	platform := models.Platform{Name: body.Name, Url: body.Url}
	err := repoDb.CreatePlatform(&platform)
	if err != nil {
		errResp.Message = "An error occured internally during the write"
		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}
	return c.Status(fiber.StatusOK).JSON(platform)
}

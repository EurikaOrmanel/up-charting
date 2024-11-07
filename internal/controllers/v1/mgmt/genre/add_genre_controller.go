package genre

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddGenreController(c *fiber.Ctx) error {
	body := new(schemas.AddGenreInput)
	c.BodyParser(body)

	// admin := c.Locals("adminUser").(models.Admin)

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	genre := models.Genre{Name: body.Name}
	genreByName := repoDb.FindGenreByName(body.Name)
	if genreByName.ID != uuid.Nil {
		errResp.Message = "A genre with current name already exists"
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	err := repoDb.CreateGenre(&genre)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}

	return c.JSON(genre)

}

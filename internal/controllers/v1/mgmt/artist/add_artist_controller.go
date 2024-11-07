package artist

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddArtistController(c *fiber.Ctx) error {

	body := new(schemas.AddArtistInput)
	c.BodyParser(body)

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	artist := models.Artist{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Stagename: body.Stagename,
		Platforms: body.Platforms.ToArtistPlatform(),
	}
	artistFound := repoDb.FindArtistByStagename(body.Stagename)
	if artistFound.ID != uuid.Nil {
		errResp.Message = "An artist with current stage name already exists"
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	err := repoDb.CreateArtist(&artist)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}
	return c.JSON(artist)

}

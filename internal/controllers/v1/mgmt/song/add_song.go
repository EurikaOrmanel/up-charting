package song

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddSongController(c *fiber.Ctx) error {
	body := new(schemas.AddSongInput)
	c.BodyParser(body)

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	/*TODO:Add a song to database, make first query to the provided urls
	for the respective data for the counts..*/

	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	foundSong := repoDb.FindSongByArtistIdNTitle(body.ArtistId, body.Title)
	if foundSong.ID != uuid.Nil {
		errResp.Message = "Song already exists"
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	song := models.Song{
		Title:     body.Title,
		ArtistID:  uuid.MustParse(body.ArtistId),
		GenreID:   uuid.MustParse(body.GenreId),
		Platforms: body.Platforms.ToSongPlatform(),
	}
	err := repoDb.AddSong(&song)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}

	return c.Status(fiber.StatusOK).JSON(errResp)
}

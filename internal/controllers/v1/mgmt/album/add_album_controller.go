package album

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	audiomackServices "EurikaOrmanel/up-charter/internal/services/audiomack"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddAlbumController(c *fiber.Ctx) error {
	body := new(schemas.AddAlbumInput)
	c.BodyParser(body)

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}

	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	foundSong := repoDb.FindAlbumByArtistIdNTitle(body.ArtistId, body.Title)
	if foundSong.ID != uuid.Nil {
		errResp.Message = "Album already exists"
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	audiomackLink := body.Platforms.FindLinkByPart("audiomack.com")
	albumPlatforms := body.Platforms.ToAlbumPlatform()
	audiomackData, err := audiomackServices.AudiomackAlbumInfo(audiomackLink)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}

	foundPlat := albumPlatforms.FindAlbumPlatformByName("audiomack.com")

	album := models.Album{
		Cover:    audiomackData.Cover,
		Title:    body.Title,
		ArtistID: uuid.MustParse(body.ArtistId),
		// GenreID:  uuid.MustParse(body.GenreId),
		Platforms: albumPlatforms,
		PlayCounts: []models.AlbumPlayCount{{Count: audiomackData.Stats.Plays,
			PlatformID: foundPlat.ID}},
	}
	err = repoDb.AddAlbum(&album)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}

	return c.Status(fiber.StatusOK).JSON(album)
}

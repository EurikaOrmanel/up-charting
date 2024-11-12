package song

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	audiomackServices "EurikaOrmanel/up-charter/internal/services/audiomack"
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
	audiomackLink := body.Platforms.FindLinkByPart("audiomack.com")
	songPlatforms := body.Platforms.ToSongPlatform()
	audiomackData, err := audiomackServices.AudiomackSongInfo(audiomackLink)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}

	foundPlat := songPlatforms.FindSongPlatformByName("audiomack.com")
	song := models.Song{
		Cover:     audiomackData.Cover,
		Title:     body.Title,
		ArtistID:  uuid.MustParse(body.ArtistId),
		GenreID:   uuid.MustParse(body.GenreId),
		Platforms: songPlatforms,
	}

	err = repoDb.AddSong(&song)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(errResp)
	}
	songPlays := []*models.SongDailyPlay{{
		Count:          audiomackData.Stats.Plays,
		SongID:         song.ID,
		SongPlatformID: foundPlat.ID,
	}}
	err = repoDb.AddSongDailyPlays(songPlays)

	return c.Status(fiber.StatusOK).JSON(song)
}

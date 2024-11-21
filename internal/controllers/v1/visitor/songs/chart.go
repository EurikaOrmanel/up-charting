package songs

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
)

// gets all charts, from general,by genre,by artist.
func GetChart(c *fiber.Ctx) error {
	query := new(schemas.SongChartQuery)
	c.QueryParser(query)

	/*todo:check queries and fetch accordingly, fetch only total stream
	counts including song details
	*/

	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	cacheConfig := appConfig.CacheConfig

	charts := cacheConfig.GetTop100Chart()
	if charts == nil {
		charts = repoDb.GetChart100NSong()
	}

	return c.
		Status(fiber.StatusOK).
		JSON(charts)
}

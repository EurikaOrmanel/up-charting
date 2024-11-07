package platform

// import (
// 	"EurikaOrmanel/up-charter/config"
// 	"EurikaOrmanel/up-charter/internal/schemas"

// 	"github.com/gofiber/fiber/v2"
// )

// func AddArtistPlatformsController(c *fiber.Ctx) error {
// 	body := new(schemas.ArtistPlatformsInput)
// 	c.BodyParser(body)
// 	errResp := schemas.ErrorResponseBody{
// 		Message: "Something went wrong internally",
// 	}
// 	appConfig := c.Locals("appConfig").(config.AppConfig)
// 	repoDb := appConfig.RepoDb

// 	repoDb.AddArtistPlatforms()




// }

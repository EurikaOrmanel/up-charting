package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"

	adminArtistControllers "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/artist"

	"github.com/gofiber/fiber/v2"
)

func adminArtistRouter(router fiber.Router) {

	router.Post("/add",
		middlewares.AddArtistInputValidator,
		adminArtistControllers.AddArtistController)

}

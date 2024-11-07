package v1

import (
	adminV1GenreControllers "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/genre"
	"EurikaOrmanel/up-charter/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func adminGenreRouter(router fiber.Router) {

	router.Get("/", middlewares.ValidateGetGenreQuery,
		adminV1GenreControllers.GetGenresController)

	router.Post("/add",
		middlewares.ValidateAddGenreInput,
		adminV1GenreControllers.AddGenreController)
}

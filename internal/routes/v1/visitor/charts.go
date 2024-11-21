package visitor

import (
	visitorSongsControllers "EurikaOrmanel/up-charter/internal/controllers/v1/visitor/songs"

	"github.com/gofiber/fiber/v2"
)

func chartsRouter(router fiber.Router) {
	router.Get("/sngs/tp100", visitorSongsControllers.GetChart)
}

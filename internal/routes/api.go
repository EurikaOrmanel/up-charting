package routes

import (
	"EurikaOrmanel/up-charter/internal/routes/v1"
	"github.com/gofiber/fiber/v2"
)

func API(app *fiber.App) {
	api := app.Group("/v1")
	v1.V1(api)
}

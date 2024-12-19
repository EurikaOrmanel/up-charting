package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"

	adminAlbumCtrlrs "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/album"
	"github.com/gofiber/fiber/v2"
)

func adminAlbumRouter(router fiber.Router) {
	router.Post("/add",
		middlewares.AddAlbumInputValidator,
		adminAlbumCtrlrs.AddAlbumController)
}

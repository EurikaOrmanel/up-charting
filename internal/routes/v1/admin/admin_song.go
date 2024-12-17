package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"

	adminSongCtrlrs "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/song"

	"github.com/gofiber/fiber/v2"
)

func adminSongRouter(router fiber.Router) {
	router.Post("/add",
		middlewares.AddSongInputValidate,
		adminSongCtrlrs.AddSongController,
	)
}

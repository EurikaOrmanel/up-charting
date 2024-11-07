package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"

	adminPltfrmCtrlrs "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/platform"

	"github.com/gofiber/fiber/v2"
)

func adminPlatformRouter(router fiber.Router) {
	router.Get("/",
		middlewares.ValidateGetPlatformQuery,
		adminPltfrmCtrlrs.GetPlatformsController)

	router.Post("/add",
		middlewares.AddPlatformInputValidator,
		adminPltfrmCtrlrs.AddPlatformController)

}

package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"

	adminV1PltfrmCtrlrs "EurikaOrmanel/up-charter/internal/controllers/v1/mgmt/platform"

	"github.com/gofiber/fiber/v2"
)

func adminPlatformRouter(router fiber.Router) {
	router.Post("/add", middlewares.AddPlatformInputValidator, adminV1PltfrmCtrlrs.AddPlatformController)

}

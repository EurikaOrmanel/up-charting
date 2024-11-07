package v1

import (
	"github.com/gofiber/fiber/v2"
)

func V1(router fiber.Router) {
	adminAuthRouteGroup := router.Group("/admin")
	adminAuthRouter(adminAuthRouteGroup)

	authRequiredEndpoints(router)

}

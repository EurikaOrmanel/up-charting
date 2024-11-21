package v1

import (
	adminRoutes "EurikaOrmanel/up-charter/internal/routes/v1/admin"
	visitorRoutes "EurikaOrmanel/up-charter/internal/routes/v1/visitor"

	"github.com/gofiber/fiber/v2"
)

func V1(router fiber.Router) {
	visitorRouterGroup := router.Group("/visitor")
	visitorRoutes.VisitorRouter(visitorRouterGroup)

	adminAuthRouteGroup := router.Group("/admin")
	adminRoutes.AdminAuthRouter(adminAuthRouteGroup)
	adminRoutes.AuthRequiredEndpoints(router)
}

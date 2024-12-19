package v1

import (
	"EurikaOrmanel/up-charter/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRequiredEndpoints(router fiber.Router) {
	jwt := middlewares.AuthMiddleware()
	router.All("*",
		jwt, /*middlewares.AuthVerifier,*/
		middlewares.InjectVerifiedUserAuth)

	adminMgmntActionRouter(router)

	// projectRouteGroup := router.Group("/projects")

}

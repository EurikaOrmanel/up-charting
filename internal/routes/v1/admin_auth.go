package v1

import (
	adminAuthControllers "EurikaOrmanel/up-charter/internal/controllers/v1/auth"
	"EurikaOrmanel/up-charter/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func adminAuthRouter(router fiber.Router) {
	router.Post("/register",
		middlewares.ValidateAdminReg,
		adminAuthControllers.AdminRegisterController)
	router.Post("/login", middlewares.ValidateAdminLogin,
		adminAuthControllers.AdminLoginController)
}

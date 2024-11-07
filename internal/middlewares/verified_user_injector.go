package middlewares

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func InjectVerifiedUserAuth(c *fiber.Ctx) error {
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb

	errResp := schemas.ErrorResponseBody{
		Message: "Something went wrong internally",
	}
	userLocal := c.Locals("user")

	if userLocal == nil {
		errResp.Message = "Please provide a valid authorization token"
		return c.Status(fiber.StatusUnauthorized).JSON(errResp)
	}
	userAuth := userLocal.(*jwt.Token)
	claims := userAuth.Claims.(jwt.MapClaims)
	authId := claims["id"].(string)

	foundUserAuth := repoDb.FindAdminByID(authId)
	if foundUserAuth.ID.String() != authId {
		errResp.Message = "User not found"

		return c.Status(fiber.StatusNotFound).JSON(errResp)

	}
	c.Locals("adminUser", foundUserAuth)
	return c.Next()

}

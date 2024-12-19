package middlewares

import (
	"EurikaOrmanel/up-charter/internal/schemas"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {

	tokenSecret := os.Getenv("TOKEN_SECRET")

	return jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{Key: []byte(tokenSecret)},
		ErrorHandler: jwtErrorHandler,
	})

}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(schemas.ErrorResponseBody{
		Message: err.Error(),
	})
}
func AuthVerifier(c *fiber.Ctx) error {

	// userAuth := c.Locals("user").(*jwt.Token)
	// claims := userAuth.Claims.(jwt.MapClaims)
	// authId := claims["id"].(string)
	// adminVerified := claims["verified"].(bool)
	adminVerified := true
	if !adminVerified {
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.ErrorResponseBody{Message: "Account not verified"})
	}
	return c.Next()
}

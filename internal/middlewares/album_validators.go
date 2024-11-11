package middlewares

import (
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func AddAlbumInputValidator(c *fiber.Ctx) error {
	body := new(schemas.AddAlbumInput)
	c.BodyParser(body)
	validatedBody := validate.Struct(*body)
	if !validatedBody.Validate() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validatedBody.Errors)
	}
	return c.Next()
}

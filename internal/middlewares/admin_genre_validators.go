package middlewares

import (
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func ValidateAddGenreInput(c *fiber.Ctx) error {

	body := new(schemas.AddGenreInput)

	c.BodyParser(body)
	validatedBody := validate.Struct(*body)
	if !validatedBody.Validate() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validatedBody.Errors)
	}
	return c.Next()

}

func ValidateGetGenreQuery(c *fiber.Ctx) error {
	query := new(schemas.GenreQuery)
	c.QueryParser(query)
	validatedBody := validate.Struct(*query)
	if !validatedBody.Validate() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validatedBody.Errors)
	}
	return c.Next()
}

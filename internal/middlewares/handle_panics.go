package middlewares

import (
	"EurikaOrmanel/up-charter/internal/schemas"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandlePanics(c *fiber.Ctx, err error) error {
	log.Println(err)
	return c.Status(fiber.StatusInternalServerError).JSON(schemas.ErrorResponseBody{Message: err.Error()})
}

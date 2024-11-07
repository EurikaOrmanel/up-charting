package middlewares

import (
	"EurikaOrmanel/up-charter/internal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func ValidateAdminReg(c *fiber.Ctx) error {
	body := new(schemas.AdminRegisterInput)
	c.BodyParser(body)
	validatedBody := validate.Struct(body)
	// validate.AddValidator("ghPhone", validators.GHPhoneValidator)

	if len(body.Phone) == 0 && len(body.Email) == 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"phone": map[string]any{
				"message": "Phone or email must be provided",
			},
			"email": map[string]any{
				"message": "Phone or email must be provided",
			},
		})
	}
	if !validatedBody.Validate() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validatedBody.Errors)
	}
	return c.Next()

}

func ValidateAdminLogin(c *fiber.Ctx) error {
	body := new(schemas.AdminLoginInput)
	c.BodyParser(body)
	validatedBody := validate.Struct(body)
	if !validatedBody.Validate() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validatedBody.Errors)
	}
	return c.Next()
}

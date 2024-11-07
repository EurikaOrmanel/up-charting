package v1

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"EurikaOrmanel/up-charter/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AdminLoginController(c *fiber.Ctx) error {
	body := new(schemas.AdminLoginInput)
	c.BodyParser(body)
	errRes := schemas.ErrorResponseBody{}
	appConfig := c.Locals("appConfig").(config.AppConfig)

	repoDb := appConfig.RepoDb
	cacheConfig := appConfig.CacheConfig
	foundAdmin := models.Admin{}
	if len(body.Email) > 0 {
		foundAdmin = repoDb.FindAdminByEmail(body.Email)

	} else if len(body.Email) > 0 {
		foundAdmin = repoDb.FindAdminByEmail(body.Email)
	}
	if foundAdmin.ID == uuid.Nil {
		errRes.Message = "Account not found"
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	passwordsMatch := utils.VerifyPassword(body.Password, foundAdmin.Password)

	if !passwordsMatch {
		errRes.Message = "Password mismatch"
		return c.Status(fiber.StatusUnauthorized).JSON(errRes)
	}

	respBody, err := utils.GenerateAuthResponse(cacheConfig, foundAdmin)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errRes)
	}

	return c.JSON(respBody)
}

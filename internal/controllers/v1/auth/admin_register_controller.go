package v1

import (
	"fmt"
	"log"

	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	"EurikaOrmanel/up-charter/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AdminRegisterController(c *fiber.Ctx) error {
	body := new(schemas.AdminRegisterInput)
	c.BodyParser(body)
	errRes := schemas.ErrorResponseBody{}
	appConfig := c.Locals("appConfig").(config.AppConfig)
	repoDb := appConfig.RepoDb
	cacheConfig := appConfig.CacheConfig
	foundAdmin := repoDb.FindAdminByPhone(body.Phone)

	if foundAdmin.ID != uuid.Nil {
		errRes.Message = "Phone number is already associated with an account."
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}
	foundAdmin = repoDb.FindAdminByEmail(body.Email)

	if foundAdmin.ID != uuid.Nil {
		errRes.Message = "Email is already associated with an account."
		return c.Status(fiber.StatusBadRequest).JSON(errRes)
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		log.Println(err.Error())
		errRes.Message = "Something went wrong internally."
		return c.Status(fiber.StatusInternalServerError).JSON(errRes)
	}
	body.Password = hashedPassword
	admin := models.Admin{Fullname: body.Fullname, Email: body.Email, Password: body.Password, Phone: body.Phone}
	err = repoDb.CreateAdmin(&admin)

	if err != nil {
		fmt.Println(err.Error())
		errRes.Message = "Failed to create account"
		return c.Status(fiber.StatusInternalServerError).JSON(errRes)
	}
	respBody, err := utils.GenerateAuthResponse(cacheConfig, admin)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errRes)
	}
	return c.JSON(respBody)
}

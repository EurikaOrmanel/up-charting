package main

import (
	"EurikaOrmanel/up-charter/config"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"

	cronjobs "EurikaOrmanel/up-charter/internal/cron_jobs"
	"EurikaOrmanel/up-charter/internal/middlewares"
	"EurikaOrmanel/up-charter/internal/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	appConfig := config.Config()
	appPort := os.Getenv("BACKEND_PORT")
	cronjobs.ConfigJobs(appConfig)
	fiberConfig := fiber.Config{
		ErrorHandler: middlewares.HandlePanics,
	}
	app := fiber.New(fiberConfig)
	app.Use(recover.New())
	app.Use(cors.New())

	app.Use(config.InjectAppConfig(appConfig))
	routes.API(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", appPort)))
}

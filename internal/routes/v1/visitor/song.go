package visitor

import "github.com/gofiber/fiber/v2"

func songChartRouter(router fiber.Router) {
	router.Get("/top100")
	
}

package visitor

import "github.com/gofiber/fiber/v2"

func VisitorRouter(router fiber.Router) {
	chartsRouterGroup := router.Group("/chrts")
	chartsRouter(chartsRouterGroup)
}

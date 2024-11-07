package v1

import "github.com/gofiber/fiber/v2"

func adminMgmntActionRouter(router fiber.Router) {
	routeGroup := router.Group("/admin/mgmt")

	adminGenreRouterGrp := routeGroup.Group("/gnrs")
	adminGenreRouter(adminGenreRouterGrp)

	adminPltfrmRouterGrp := routeGroup.Group("/pltfrms")
	adminPlatformRouter(adminPltfrmRouterGrp)

	adminArtstRouterGrp := routeGroup.Group("/artsts")
	adminArtistRouter(adminArtstRouterGrp)

	songRouterGrp := routeGroup.Group("/sngs")
	adminSongRouter(songRouterGrp)

}

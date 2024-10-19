package route

import (
	group2 "api/pkg/route/group"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {

	group2.RouteGroupUser(app)
	group2.RouteGroupProduct(app)

}

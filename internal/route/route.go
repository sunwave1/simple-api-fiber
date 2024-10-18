package route

import (
	"api/internal/route/group"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {

	group.RouteGroupUser(app)
	group.RouteGroupProduct(app)

}

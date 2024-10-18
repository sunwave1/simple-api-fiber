package main

import (
	"api/internal/route"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	route.InitializeRoutes(app)

	app.Listen(":8000")
}

package main

import (
	"api/internal/route"
	"api/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	if !config.CheckFileEnvExist() {
		panic("File env not exist")
	}

	app := fiber.New()

	route.InitializeRoutes(app)

	app.Listen(":8000")
}

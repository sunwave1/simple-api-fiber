package main

import (
	"api/pkg/config"
	"api/pkg/route"
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

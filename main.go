package main

import (
	"gok8s/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Router(app)

	app.Listen(":5000")
}

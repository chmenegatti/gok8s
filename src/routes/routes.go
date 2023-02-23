package routes

import (
	"gok8s/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", controller.HelloWorld)
	app.Get("/:name", controller.GetName)
}

package routes

import (
	"gok8s/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", controller.GetName)
	app.Get("/:name", controller.GetName)
}

package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetName(c *fiber.Ctx) error {
	name := c.Params("name")
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(fmt.Sprintf("<h1>Hello %s</h1>", name))
}

func HelloWorld(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString("<h1>Hello World do K8s</h1>")
}

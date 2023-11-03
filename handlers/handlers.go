package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.SendString("silkania v0.0.1")
}

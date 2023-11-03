package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/routes"
)

func main() {
	app := fiber.New()
	routes.Config(app)

	log.Fatal(app.Listen(":3000"))
}

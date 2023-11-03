package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/handlers"
)

func Config(app *fiber.App) {
	app.Get("/", handlers.Index)
	app.Get("/index/:name/search", handlers.Search)
	app.Post("/index", handlers.CreateIndex)
}

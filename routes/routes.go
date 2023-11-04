package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/handlers"
)

func Config(app *fiber.App) {
	app.Get("/", handlers.Index)
	index := app.Group("/index")

	index.Delete("/:name", handlers.DeleteIndex)
	index.Post("/", handlers.CreateIndex)

	index.Get("/:name/search", handlers.Search)
	index.Put("/:name/batch", handlers.Batch)
	index.Get("/:name/:id", handlers.GetDocument)
	index.Put("/:name/:id", handlers.IndexDocument)
	index.Delete("/:name/:id", handlers.DeleteDocument)
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
)

func Search(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	q := ctx.Query("q")

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	res, total, err := ix.Search(q)
	if err != nil {
		return serverErr(ctx, err)
	}

	return ctx.JSON(models.SearchResult{
		Total: total,
		Items: res,
	})
}

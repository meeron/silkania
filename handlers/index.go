package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
)

func CreateIndex(ctx *fiber.Ctx) error {
	body := models.CreateIndexReq{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return badRequest(ctx, "'Name' cannot be empty")
	}

	if index.Get(body.Name) != nil {
		return existsError(ctx)
	}

	if err := index.Create(body.Name, body.DefaultLang, body.Mapping); err != nil {
		return serverErr(ctx, err)
	}

	return created(ctx)
}

func DeleteIndex(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	if err := index.Drop(name); err != nil {
		return serverErr(ctx, err)
	}

	return ok(ctx)
}

func GetIndexStats(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	return ctx.JSON(ix.GetStats())
}

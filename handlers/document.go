package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
)

func IndexDocument(ctx *fiber.Ctx) error {
	var body any
	name := ctx.Params("name")
	id := ctx.Params("id")

	if err := ctx.BodyParser(&body); err != nil {
		return badRequestFromErr(ctx, err)
	}

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	if err := ix.IndexDocument(id, body); err != nil {
		return serverErr(ctx, err)
	}

	return ok(ctx)
}

func GetDocument(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	id := ctx.Params("id")

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	doc, err := ix.GetDocument(id)
	if err != nil {
		return serverErr(ctx, err)
	}

	if doc == nil {
		return noContent(ctx)
	}

	return ctx.JSON(doc)
}

func DeleteDocument(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	id := ctx.Params("id")

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	if err := ix.DeleteDocument(id); err != nil {
		return serverErr(ctx, err)
	}

	return ok(ctx)
}

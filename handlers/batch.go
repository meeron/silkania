package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
)

func Batch(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	var body models.BatchReq

	if err := ctx.BodyParser(&body); err != nil {
		return badRequestFromErr(ctx, err)
	}

	if body.IdField == "" {
		return badRequest(ctx, "'IdField' cannot be empty")
	}

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	if err := ix.Batch(body.IdField, body.Items); err != nil {
		return serverErr(ctx, err)
	}

	return ctx.SendStatus(http.StatusOK)
}

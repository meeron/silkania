package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/models"
)

func badRequestFromErr(ctx *fiber.Ctx, err error) error {
	return ctx.
		Status(http.StatusBadRequest).
		JSON(models.BadRequestError(err.Error()))
}

func badRequest(ctx *fiber.Ctx, message string) error {
	return ctx.
		Status(http.StatusBadRequest).
		JSON(models.BadRequestError(message))
}

func serverErr(ctx *fiber.Ctx, err error) error {
	return ctx.
		Status(http.StatusInternalServerError).
		JSON(models.ServerError(err))
}

func ok(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func noContent(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}

func created(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusCreated)
}

func notFoundError(ctx *fiber.Ctx) error {
	return ctx.
		Status(http.StatusUnprocessableEntity).
		JSON(models.NotFoundError())
}

func existsError(ctx *fiber.Ctx) error {
	return ctx.
		Status(http.StatusUnprocessableEntity).
		JSON(models.ExistsError())
}

package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
)

func Search(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	q := ctx.Query("q")

	ix := index.Get(name)
	if ix == nil {
		return ctx.Status(422).JSON(models.NotFoundError())
	}

	return ctx.SendString(fmt.Sprintf("Db: %s, Query: %s", name, q))
}

func CreateIndex(ctx *fiber.Ctx) error {
	body := models.CreateIndexReq{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return ctx.Status(400).JSON(models.BadRequestError("'Name' cannot be empty"))
	}

	if index.Get(body.Name) != nil {
		return ctx.Status(422).JSON(models.ExistsError())
	}

	if err := index.Create(body.Name, body.Mapping); err != nil {
		return ctx.Status(500).
			JSON(models.ServerError(err))
	}

	return ctx.SendStatus(201)
}

func DeleteIndex(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	if err := index.Drop(name); err != nil {
		return ctx.Status(500).
			JSON(models.ServerError(err))
	}

	return ctx.SendStatus(200)
}

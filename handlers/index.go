package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
)

func Search(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	q := ctx.Query("q")

	ix := index.Get(name)
	if ix == nil {
		return ctx.Status(422).JSON(Error{Code: "IndexNotFound"})
	}

	return ctx.SendString(fmt.Sprintf("Db: %s, Query: %s", name, q))
}

func CreateIndex(ctx *fiber.Ctx) error {
	body := CreateIndexReq{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return ctx.Status(400).JSON(Error{Code: "BadRequest", Message: "'Name' cannot be empty"})
	}

	if body.Lang == "" {
		return ctx.Status(400).JSON(Error{Code: "BadRequest", Message: "'Lang' cannot be empty"})
	}

	if index.Get(body.Name) != nil {
		return ctx.Status(422).JSON(Error{Code: "Exists"})
	}

	if err := index.Create(body.Name, body.Lang); err != nil {
		return ctx.Status(422).
			JSON(Error{Code: "ServerError", Message: fmt.Sprintf("%v", err)})
	}

	return ctx.SendStatus(201)
}

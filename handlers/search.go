package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
)

func Search(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	searchReq := &models.SearchReq{
		Q:            ctx.Query("q"),
		SortBy:       ctx.Query("sort_by"),
		Page:         ctx.QueryInt("page", 1),
		ItemsPerPage: ctx.QueryInt("items_per_page", 10),
	}

	if err := searchReq.Validate(); err != nil {
		return badRequestFromErr(ctx, err)
	}

	ix := index.Get(name)
	if ix == nil {
		return notFoundError(ctx)
	}

	res, total, err := ix.Search(searchReq)
	if err != nil {
		return serverErr(ctx, err)
	}

	return ctx.JSON(models.SearchResult{
		Total: total,
		Items: res,
	})
}

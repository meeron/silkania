package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/models"
	"github.com/meeron/silkania/server"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("silkania v0.0.1")
}

func GetStats(ctx *fiber.Ctx) error {
	serverStats := models.ServerStats{
		Uptime:  server.Uptime(),
		Indexes: make(map[string]map[string]any),
	}

	for name, ix := range index.All() {
		stats := ix.GetStats()
		indexStats, _ := stats["index"].(map[string]any)

		serverStats.Indexes[name] = make(map[string]any)
		serverStats.Indexes[name]["DocCount"] = stats["DocCount"]
		serverStats.Indexes[name]["CurOnDiskBytes"] = indexStats["CurOnDiskBytes"]
	}

	return ctx.JSON(serverStats)
}

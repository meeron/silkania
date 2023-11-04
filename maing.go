package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/routes"
	"github.com/meeron/silkania/server"
)

const (
	IndexBasePath = "./bin/index"
	Port          = 3000
)

func main() {
	if err := index.Load(IndexBasePath); err != nil {
		panic(err)
	}

	app := fiber.New()

	routes.Config(app)

	server.StartTime = time.Now()

	log.Fatal(app.Listen(fmt.Sprintf(":%d", Port)))
}

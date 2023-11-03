package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/meeron/silkania/index"
	"github.com/meeron/silkania/routes"
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

	log.Fatal(app.Listen(fmt.Sprintf(":%d", Port)))
}

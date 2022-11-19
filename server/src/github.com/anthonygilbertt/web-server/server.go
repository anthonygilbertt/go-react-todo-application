package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Jessica!")
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.SendString("TODOs")
	})

	log.Fatal(app.Listen(":3000"))
}

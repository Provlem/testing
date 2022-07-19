package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"testing/req"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		var payloads req.Payloads
		if err := c.BodyParser(&payloads); err != nil {
			return err
		}
		for _, payload := range payloads.Payloads {
			c.SendString(payload.Key + ": " + payload.Value)
		}
		return nil
	})

	app.Listen(":3000")
}

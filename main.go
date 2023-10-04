package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()
    // Initialize default config
    app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        //log.Info("[/]: Hello world")
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"testing": "Hello, World!"})
    })

    app.Listen(":3000")
}

package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	//"github.com/1garo/daedalus/pkg/database"
	"github.com/1garo/daedalus/pkg/env"
	"github.com/1garo/daedalus/pkg/router"
	"github.com/1garo/daedalus/types"
)



func NewApplication() *fiber.App {
	env.SetupEnvFile()
	//database.SetupDatabase()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	var validate = validator.New()
	myValidator := &types.XValidator{
		Validator: validate,
	}

	app.Use(func(c *fiber.Ctx) error {
	  c.Locals("validator", myValidator)
	  return c.Next()
	})

	app.Get("/dashboard", monitor.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"data": "HealthCheck!"})
	})

	router.InstallRouter(app)

	return app
}


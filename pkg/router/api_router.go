package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	controllers "github.com/1garo/daedalus/app/controllers"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())


	api.Post("recipe", controllers.CreateRecipe)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}

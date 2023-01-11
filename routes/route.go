package routes

import (
	"github.com/XinceChan/go-blog-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/blog", controllers.Index)
}

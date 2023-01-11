package routes

import (
	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/blog", controllers.Index)
	app.Get("/categories", controllers.Category)
	app.Get("/tags", controllers.Tag)
	app.Get("/article", controllers.Article)
	app.Get("/extra-nav", controllers.ExtraNav)
	app.Get(config.Cfg.DashboardEntrance, controllers.Dashboard)
	app.Post(config.Cfg.GitHookUrl, controllers.GithubHook)
}

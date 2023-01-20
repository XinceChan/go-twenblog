package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/XinceChan/go-blog-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	models.CompiledContent()
}

//go:embed views/*
var viewsfs embed.FS

func main() {

	// Create a new engine
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	engine.AddFunc("unescapeHTML", func(s string) template.HTML {
		return template.HTML(s)
	})

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		AppName: "Twen Go BLOG APP v0.0.1",
		Views:   engine,
	})

	// Get static html files
	app.Static("/public", "./public")
	app.Static("/assets", config.Cfg.DocumentAssetsDir)
	app.Static("/images", config.Cfg.CurrentDir+"/images")

	routes.Setup(app)

	app.Listen(":" + config.Cfg.Port)
}

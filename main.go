package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/XinceChan/go-blog-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func init() {

}

//go:embed views/*
var viewsfs embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// Create a new engine
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		AppName: "Twen Go BLOG APP v0.0.1",
		Views:   engine,
	})
	// Get static html files
	app.Static("/public", "./public")
	// Get uploaded file
	app.Static("/files", "./files")
	routes.Setup(app)

	app.Listen(":" + port)
}

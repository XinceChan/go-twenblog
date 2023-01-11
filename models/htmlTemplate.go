package models

import (
	"github.com/XinceChan/go-blog-backend/config"
	"github.com/gofiber/fiber/v2"
)

func BuildViewDate(title string, data interface{}) fiber.Map {
	return fiber.Map{
		"Title":  title,
		"Data":   data,
		"Config": config.Cfg,
		"Navs":   Navigation,
	}
}

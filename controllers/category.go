package controllers

import (
	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Category(c *fiber.Ctx) error {
	result := models.GroupByCategory(&models.ArticleList, config.Cfg.CategoryDisplayQuantity)

	return c.Render("views/categories", models.BuildViewDate("Category", result))
}

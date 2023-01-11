package controllers

import (
	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Tag(c *fiber.Ctx) error {
	result := models.GroupByTag(&models.ArticleList, config.Cfg.TagDisplayQuantity)

	return c.Render("views/tags", models.BuildViewDate("Tag", result))
}

package controllers

import (
	"strconv"

	"github.com/XinceChan/go-blog-backend/database"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	// which page to display xxxx/?page=xx
	page, _ := strconv.Atoi(c.Query("page", "1"))
	// every page limit
	limit := 3
	offset := (page - 1) * limit
	// all articles
	var articles models.Articles
	database.DB.Offset(offset).Limit(limit).Find(&articles)

	for _, article := range articles {
		var tags []models.Tag
		database.DB.Where("article_id=?", article.ID).Limit(limit).Find(&tags)

		article.Tags = tags
	}

	return c.Render("views/index", fiber.Map{
		"Title": "Blog",
		"Data":  articles,
	})
}

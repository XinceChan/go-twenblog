package controllers

import (
	"log"

	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Article(c *fiber.Ctx) error {
	key := c.Query("key")

	path := models.ArticleShortUrlMap[key]

	articleDetail, err := models.ReadArticleDetail(path)

	if err != nil {
		log.Println(err)
	}

	return c.Render("views/article", models.BuildViewDate("Article", articleDetail))
}

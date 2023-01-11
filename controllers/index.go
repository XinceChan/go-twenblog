package controllers

import (
	"strconv"

	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	var articles = models.ArticleList

	search := c.Query("search")
	category := c.Query("category")
	tag := c.Query("tag")

	// search function
	if category != "" || tag != "" {
		articles = models.SearchArticle(&articles, search, category, tag)
	}

	result := models.Pagination(&articles, page, config.Cfg.PageSize)

	// fmt.Println(result)
	return c.Render("views/index", models.BuildViewDate("Blog", result))
}

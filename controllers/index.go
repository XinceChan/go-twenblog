package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	// which page to display xxxx/?page=xx
	// page, _ := strconv.Atoi(c.Query("page", "1"))
	// every page limit
	// limit := 3

	// all articles
	// var articles models.Articles

	return c.Render("views/index", fiber.Map{
		"Title": "Blog",
	})
}

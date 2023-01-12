package controllers

import (
	"log"

	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func ExtraNav(c *fiber.Ctx) error {
	name := c.Query("name", "About")

	for _, nav := range models.Navigation {
		if nav.Title == name {
			articleDetail, err := models.ReadArticleDetail(nav.Path)
			// fmt.Println(articleDetail)
			if err != nil {
				log.Println(err)
			}
			return c.Render("views/extraNav", models.BuildViewDate(nav.Title, articleDetail))
		}
	}
	return nil
}

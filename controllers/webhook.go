package controllers

import (
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GithubHook(c *fiber.Ctx) error {
	models.CompiledContent()

	return c.JSON(fiber.Map{
		"msg": "Comment ok!",
	})
}

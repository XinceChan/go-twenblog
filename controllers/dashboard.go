package controllers

import (
	"strconv"

	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	var dashboardMsg []string

	index, err := strconv.Atoi(c.Query("theme"))
	if err == nil && index < len(config.Cfg.ThemeOption) {
		config.Cfg.ThemeColor = config.Cfg.ThemeOption[index]
		dashboardMsg = append(dashboardMsg, "颜色切换成功!")
	}

	action := c.Query("action")
	if action == "updateArticle" {
		models.CompiledContent()
		dashboardMsg = append(dashboardMsg, "文章更新成功!")
	}

	return c.Render("views/dashboard", models.BuildViewDate("Dashboard", map[string]interface{}{
		"msg": dashboardMsg,
	}))
}

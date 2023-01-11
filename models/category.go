package models

import (
	"strings"

	"github.com/XinceChan/go-blog-backend/config"
)

type Category struct {
	Name     string
	Quantity int
	Articles Articles
}

type Categories []Categories

func GetCategoryName(path string) string {
	var categoryName string
	newPath := strings.Replace(path, config.Cfg.DocumentContentDir+"/", "", 1)

	// 文件在根目录下(content/)没有分类名称
	if !strings.Contains(newPath, "/") {
		categoryName = "未分类"
	} else {
		categoryName = strings.Split(newPath, "/")[0]
	}
	return categoryName
}

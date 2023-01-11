package models

import (
	"sync"

	"github.com/XinceChan/go-blog-backend/config"
)

var Navigation Navs
var ArticleList Articles
var ArticleShortUrlMap map[string]string // 用来保证文章 shortUrl 唯一和快速定位文章

func CompiledContent() {
	config.Initial() // 克隆或者更新文档库
	// 对内容的生成
	wg := sync.WaitGroup{}
	var err error
	// 导航

	// 文章
	wg.Add(1)
	go func() {
		ArticleList, ArticleShortUrlMap, err = initArticlesAndImages(config.Cfg.DocumentContentDir)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

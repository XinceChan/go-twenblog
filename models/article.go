package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/XinceChan/go-blog-backend/config"
	"github.com/XinceChan/go-blog-backend/utils"
	"github.com/yuin/goldmark"
)

type Time time.Time

type Article struct {
	Title       string   `json:"title"`
	Date        Time     `json:"date"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
	MusicId     string   `json:"musicId"`
	Path        string
	ShortUrl    string
	Category    string
}

type Articles []Article

type ArticleDetail struct {
	Article
	Body string
}

func initArticlesAndImages(dir string) (Articles, map[string]string, error) {
	var articles Articles
	shortUrlMap := make(map[string]string)

	articles, err := RecursiveReadArticles(dir)
	if err != nil {
		return articles, shortUrlMap, err
	}
	sort.Sort(articles)
	for i := len(articles) - 1; i >= 0; i-- {
		// 这里必须使用倒序的方式生成 shortUrl,因为如果有相同的文章标题，
		// 倒序会将最老的文章优先生成shortUrl，保证和之前的 shortUrl一样
		article := articles[i]
		keyword := utils.GenerateShortUrl(article.Title, func(url, keyword string) bool {
			// 保证keyword统一
			_, ok := shortUrlMap[keyword]
			return !ok
		})
		articles[i].ShortUrl = keyword
		shortUrlMap[keyword] = article.Path
	}
	return articles, shortUrlMap, nil
}

func SearchArticle(articles *Articles, search, category, tag string) Articles {
	var articleList Articles
	for _, article := range *articles {
		pass := true

		if search != "" && !strings.Contains(article.Title, search) {
			pass = false
		}
		if category != "" && !strings.Contains(article.Category, category) {
			pass = false
		}
		if tag != "" {
			pass = false
			for _, tagItem := range article.Tags {
				if strings.Contains(tagItem, tag) {
					pass = true
					break
				}
			}
		}
		if pass {
			articleList = append(articleList, article)
		}
	}
	return articleList
}

func RecursiveReadArticles(dir string) (Articles, error) {
	var articles Articles

	difInfo, err := os.Stat(dir)
	if err != nil {
		return articles, err
	}

	if !difInfo.IsDir() {
		return articles, errors.New("目标不是一个目录")
	}

	fileOrDir, err := os.ReadDir(dir)
	if err != nil {
		return articles, err
	}

	for _, fileInfo := range fileOrDir {
		name := fileInfo.Name()
		path := dir + "/" + name
		upperName := strings.ToUpper(name)
		if fileInfo.IsDir() {
			subArticles, err := RecursiveReadArticles(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, subArticles...)
		} else if strings.HasSuffix(upperName, ".MD") {
			article, err := ReadArticle(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, article)
		} else if strings.HasSuffix(upperName, ".PNG") ||
			strings.HasSuffix(upperName, ".GIF") ||
			strings.HasSuffix(upperName, ".JPG") {

			dst := config.Cfg.CurrentDir + "/images/" + name
			if !utils.IsFile(dst) {
				_, _ = utils.CopyFile(path, dst)
			}
		}
	}
	return articles, nil
}

func ReadArticle(path string) (Article, error) {
	article, _, err := readMarkdown(path)
	if err != nil {
		return article, err
	}
	return article, nil
}

func ReadArticleDetail(path string) (ArticleDetail, error) {
	_, articleDetail, err := readMarkdown(path)
	if err != nil {
		return articleDetail, err
	}
	return articleDetail, nil
}

func readMarkdown(path string) (Article, ArticleDetail, error) {
	var article Article
	var articleDetail ArticleDetail
	mdFile, err := os.Stat(path)

	if err != nil {
		return article, articleDetail, err
	}
	if mdFile.IsDir() {
		return article, articleDetail, errors.New("this path is Dir")
	}
	markdown, err := os.ReadFile(path)

	if err != nil {
		return article, articleDetail, err
	}
	markdown = bytes.TrimSpace(markdown)

	article.Path = path
	article.Category = GetCategoryName(path)
	article.Title = strings.TrimSuffix(strings.ToUpper(mdFile.Name()), ".MD")
	article.Date = Time(mdFile.ModTime())

	var buf bytes.Buffer

	if !bytes.HasPrefix(markdown, []byte("```json")) {
		article.Description = cropDesc(markdown)
		articleDetail.Article = article
		if err := goldmark.Convert(markdown, &buf); err != nil {
			article.Title = "文章[" + article.Title + "]解析 markdown 出错，请检查。"
			return article, articleDetail, nil
		}
		articleDetail.Body = buf.String()
		return article, articleDetail, nil
	}

	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	markdownArrInfo := bytes.SplitN(markdown, []byte("```"), 2)

	article.Description = cropDesc(markdownArrInfo[1])

	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &article); err != nil {
		article.Title = "文章[" + article.Title + "]解析 JSON 出错，请检查。"
		article.Description = err.Error()
		return article, articleDetail, nil
	}
	article.Path = path
	article.Title = strings.ToUpper(article.Title)

	articleDetail.Article = article

	if err := goldmark.Convert(markdownArrInfo[1], &buf); err != nil {
		article.Title = "文章[" + article.Title + "]解析 markdown 出错，请检查。"
		return article, articleDetail, nil
	}

	articleDetail.Body = buf.String()

	// fmt.Println(articleDetail.Body)
	return article, articleDetail, nil
}

// Limit Description Len
func cropDesc(c []byte) string {
	content := []rune(string(c))
	contentLen := len(content)

	if contentLen <= config.Cfg.DescriptionLen {
		return string(content[0:contentLen])
	}

	return string(content[0:config.Cfg.DescriptionLen])
}

func (t *Time) UnmarshalJSON(b []byte) error {
	date, err := time.ParseInLocation(`"`+config.Cfg.TimeLayout+`"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*t = Time(date)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"` + config.Cfg.TimeLayout + `"`)), nil
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Less(i, j int) bool {
	return time.Time(a[i].Date).After(time.Time(a[j].Date))
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

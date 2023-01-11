package models

import (
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Nav struct {
	Title string
	Path  string
}

type Navs []Nav

func initExtraNav(dir string) (Navs, error) {
	var navigation Navs
	var extraNav Articles

	extraNav, err := RecursiveReadArticles(dir)
	if err != nil {
		return navigation, err
	}
	sort.Sort(extraNav)

	for _, article := range extraNav {
		c := cases.Title(language.English)
		title := c.String(strings.ToLower(article.Title))
		navigation = append(navigation, Nav{Title: title, Path: article.Path})
	}

	return navigation, nil
}

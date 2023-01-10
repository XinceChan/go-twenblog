package models

import "time"

type Article struct {
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Author      string    `json:"author"`
	MusicId     string    `json:"musicId"`
	Path        string
	ShortUrl    string
	Category    string
}

type Articles []Article

type ArticleDetail struct {
	Article
	Body string
}

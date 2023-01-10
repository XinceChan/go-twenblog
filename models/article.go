package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string    `json:"title"`
	UploadedAt  time.Time `json:"date"`
	Description string    `json:"description"`
	Tags        []Tag     `json:"tags"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Path        string    `json:"path"` // File Path
	UserID      string    `json:"userid"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
}

type Articles []Article

var ArticleList Articles

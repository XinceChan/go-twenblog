package models

type Tag struct {
	Name     string
	Quantity int
	Articles Articles
}

type Tags []Tag

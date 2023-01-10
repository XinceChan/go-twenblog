package models

type Category struct {
	Name     string
	Quantity int
	Articles Articles
}

type Categories []Categories

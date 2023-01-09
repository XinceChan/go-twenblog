package models

type Category struct {
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
	Articles []Article `json:"articles" gorm:"-"`
}

type Categories []Category

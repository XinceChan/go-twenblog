package models

type Tag struct {
	Name     string    `json:"name"`
	Quantity string    `json:"quantity"`
	Articles []Article `json:"articles" gorm:"-"`
}

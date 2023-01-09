package models

type User struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  []byte `json:"password" gorm:"-"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"` // Avatar path
}

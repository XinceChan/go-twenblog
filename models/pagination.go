package models

type PageResult struct {
	List      Articles `json:"list"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	TotalPage int
}

package models

type PageResult struct {
	List      Articles `json:"list"`
	Total     uint     `json:"total"`
	Page      uint     `json:"page"`
	PageSize  uint     `json:"page_size"`
	TotalPage uint
}

func Pagination() PageResult {
	return PageResult{}
}

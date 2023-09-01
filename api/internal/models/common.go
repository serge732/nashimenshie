package models

type Pagination struct {
	PageCount   int `json:"pageCount"`
	CurrentPage int `json:"currentPage"`
}

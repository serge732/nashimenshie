package models

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       int      `json:"price"`
}

type Assortment struct {
	Products   []Product  `json:"products"`
	Pagination Pagination `json:"pagination"`
}

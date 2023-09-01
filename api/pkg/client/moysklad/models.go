package moysklad

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SalePrices  []struct {
		Value float64 `json:"value"`
	} `json:"salePrices"`
}

type Assortment struct {
	Meta struct {
		Size int `json:"size"`
	} `json:"meta"`
	Rows []Product `json:"rows"`
}

type Images struct {
	Rows []struct {
		Meta struct {
			DownloadHref string `json:"downloadHref"`
		} `json:"meta"`
	} `json:"rows"`
}

package merchandiser_goods

type Attribute struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Goods struct {
	GoodsId    string      `json:"goodsId"`
	Title      string      `json:"title"`
	ImageLink  string      `json:"titleImage"`
	Attributes []Attribute `json:"attributes"`
}

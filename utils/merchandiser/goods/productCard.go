package merchandiser_goods

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ProductCardPayload struct {
	GoodsId string `json:"goodsId"`
}

type ProductCardItem struct {
	Goods          Goods  `json:"goods"`
	CollectionName string `json:"similarCollectionName"`
	Description    string `json:"description"`
}

type ProductCardResponse struct {
	Item ProductCardItem `json:"item"`
}

func ProductCard(payload ProductCardPayload) (*ProductCardResponse, error) {
	var productCardResponse ProductCardResponse

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("POST", "https://sbermegamarket.ru/api/mobile/v1/catalogService/catalog/productCard", &buf)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &productCardResponse)

	return &productCardResponse, nil
}

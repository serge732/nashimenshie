package merchandiser_goods

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type SearchPayload struct {
	SearchText string `json:"searchText"`
}

type SearchResponseItem struct {
	Goods Goods `json:"goods"`
}

type SearchResponse struct {
	Items []SearchResponseItem `json:"items"`
}

func Search(payload SearchPayload) (*SearchResponse, error) {
	var searchResponse SearchResponse

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("POST", "https://sbermegamarket.ru/api/mobile/v1/catalogService/catalog/search", &buf)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &searchResponse)

	return &searchResponse, nil
}

package merchandiser_moysklad

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ProductUpdatePayloadImage struct {
	FileName string `json:"filename"`
	Content  string `json:"content"`
}

type ProductUpdatePayload struct {
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Images      []ProductUpdatePayloadImage `json:"images"`
}

func ProductUpdate(token string, productId string, payload ProductUpdatePayload) (bool, error) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("PUT", "https://online.moysklad.ru/api/remap/1.2/entity/product/"+productId, &buf)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return true, nil
}

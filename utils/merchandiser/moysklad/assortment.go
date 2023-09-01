package merchandiser_moysklad

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type AssortmentResponseRow struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AssortmentResponse struct {
	Rows []AssortmentResponseRow `json:"rows"`
}

func Assortment(token string, offset int) (*AssortmentResponse, error) {
	var assortment AssortmentResponse

	req, _ := http.NewRequest("GET", "https://online.moysklad.ru/api/remap/1.2/entity/assortment?offset="+strconv.Itoa(offset), nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &assortment)

	return &assortment, nil
}

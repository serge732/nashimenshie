package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"nashimenshie_api/pkg/client/moysklad"
	"nashimenshie_api/pkg/logging"
	"net/http"
	"strconv"
)

type AssortmentRepository struct {
	moyskladClient *moysklad.MoyskladClient
	logger         *logging.Logger
}

func NewAssortmentRepository(moyskladClient *moysklad.MoyskladClient, logger *logging.Logger) *AssortmentRepository {
	return &AssortmentRepository{
		moyskladClient: moyskladClient,
		logger:         logger,
	}
}

func (r *AssortmentRepository) GetAssortment(limit, offset int) (moysklad.Assortment, error) {
	var assortment moysklad.Assortment

	url := fmt.Sprintf("https://online.moysklad.ru/api/remap/1.2/entity/assortment?limit=%s&offset=%s",
		strconv.Itoa(limit), strconv.Itoa(offset))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		r.logger.Error(err)
		return assortment, err
	}

	resp, err := r.moyskladClient.Do(req)
	if err != nil {
		r.logger.Error(err)
		return assortment, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err)
		return assortment, err
	}

	json.Unmarshal(data, &assortment)

	return assortment, nil
}

func (r *AssortmentRepository) GetProductImages(productId string) (moysklad.Images, error) {
	var images moysklad.Images

	url := fmt.Sprintf("https://online.moysklad.ru/api/remap/1.2/entity/product/%s/images", productId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		r.logger.Error(err)
		return images, err
	}

	resp, err := r.moyskladClient.Do(req)
	if err != nil {
		r.logger.Error(err)
		return images, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err)
		return images, err
	}

	json.Unmarshal(data, &images)

	return images, nil
}

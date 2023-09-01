package service

import (
	"nashimenshie_api/internal/models"
	"nashimenshie_api/internal/repository"
	"nashimenshie_api/pkg/logging"
	"strings"
)

type AssortmentService struct {
	repository *repository.Repository
	logger     *logging.Logger
}

func NewAssortmentService(repository *repository.Repository, logger *logging.Logger) *AssortmentService {
	return &AssortmentService{
		repository: repository,
		logger:     logger,
	}
}

func (s *AssortmentService) GetAssortment(limit, offset int) (models.Assortment, error) {
	var assortment models.Assortment

	moyskladAssortment, err := s.repository.GetAssortment(limit, offset)
	if err != nil {
		s.logger.Error(err)
		return assortment, err
	}

	for _, moyskladAssortmentRow := range moyskladAssortment.Rows {
		product := models.Product{
			Id:          moyskladAssortmentRow.Id,
			Name:        moyskladAssortmentRow.Name,
			Description: moyskladAssortmentRow.Description,
			Price:       int(moyskladAssortmentRow.SalePrices[0].Value) / 100,
		}

		productImages, err := s.repository.GetProductImages(moyskladAssortmentRow.Id)
		if err != nil {
			s.logger.Error(err)
			continue
		}

		for _, productImagesRow := range productImages.Rows {
			productImageDownloadHrefParts := strings.Split(productImagesRow.Meta.DownloadHref, "/")
			productImageId := productImageDownloadHrefParts[len(productImageDownloadHrefParts)-1]
			product.Images = append(product.Images, productImageId)
		}

		assortment.Products = append(assortment.Products, product)
	}

	assortment.Pagination.PageCount = moyskladAssortment.Meta.Size / limit
	assortment.Pagination.CurrentPage = (offset / limit) + 1

	return assortment, nil
}

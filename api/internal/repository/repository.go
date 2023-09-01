package repository

import (
	"nashimenshie_api/pkg/client/moysklad"
	"nashimenshie_api/pkg/logging"
)

type Assortment interface {
	GetAssortment(limit, offset int) (moysklad.Assortment, error)
	GetProductImages(productId string) (moysklad.Images, error)
}

type Multimedia interface {
	DownloadImage(imageId string) ([]byte, error)
}

type Repository struct {
	Assortment
	Multimedia
}

func NewRepository(moyskladClient *moysklad.MoyskladClient, logger *logging.Logger) *Repository {
	return &Repository{
		Assortment: NewAssortmentRepository(moyskladClient, logger),
		Multimedia: NewMultimediaRepository(moyskladClient, logger),
	}
}

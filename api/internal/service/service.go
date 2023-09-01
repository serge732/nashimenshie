package service

import (
	"nashimenshie_api/internal/models"
	"nashimenshie_api/internal/repository"
	"nashimenshie_api/pkg/logging"
)

type Assortment interface {
	GetAssortment(limit, offset int) (models.Assortment, error)
}

type Multimedia interface {
	DownloadImage(imageId string, size uint) ([]byte, error)
}

type Service struct {
	Assortment
	Multimedia
}

func NewService(repository *repository.Repository, logger *logging.Logger) *Service {
	return &Service{
		Assortment: NewAssortmentService(repository, logger),
		Multimedia: NewMultimediaService(repository, logger),
	}
}

package service

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"nashimenshie_api/internal/repository"
	"nashimenshie_api/pkg/logging"
	"net/http"

	"github.com/nfnt/resize"
)

type MultimediaService struct {
	repository *repository.Repository
	logger     *logging.Logger
}

func NewMultimediaService(repository *repository.Repository, logger *logging.Logger) *MultimediaService {
	return &MultimediaService{
		repository: repository,
		logger:     logger,
	}
}

func (s *MultimediaService) DownloadImage(imageId string, size uint) ([]byte, error) {
	var imageBuf bytes.Buffer
	var img image.Image

	data, err := s.repository.DownloadImage(imageId)
	if err != nil {
		s.logger.Error(err)
		return imageBuf.Bytes(), err
	}

	dataType := http.DetectContentType(data)

	switch dataType {
	case "image/jpeg":
		if img, err = jpeg.Decode(bytes.NewReader(data)); err != nil {
			s.logger.Error(err)
			return imageBuf.Bytes(), err
		}
	case "image/png":
		if img, err = png.Decode(bytes.NewReader(data)); err != nil {
			s.logger.Error(err)
			return imageBuf.Bytes(), err
		}
	}

	m := resize.Resize(0, size, img, resize.Lanczos3)

	if err := png.Encode(&imageBuf, m); err != nil {
		s.logger.Error(err)
		return imageBuf.Bytes(), err
	}

	return imageBuf.Bytes(), nil
}

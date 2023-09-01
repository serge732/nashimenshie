package repository

import (
	"fmt"
	"io"
	"nashimenshie_api/pkg/client/moysklad"
	"nashimenshie_api/pkg/logging"
	"net/http"
)

type MultimediaRepository struct {
	moyskladClient *moysklad.MoyskladClient
	logger         *logging.Logger
}

func NewMultimediaRepository(moyskladClient *moysklad.MoyskladClient, logger *logging.Logger) *MultimediaRepository {
	return &MultimediaRepository{
		moyskladClient: moyskladClient,
		logger:         logger,
	}
}

func (r *MultimediaRepository) DownloadImage(imageId string) ([]byte, error) {
	url := fmt.Sprintf("https://online.moysklad.ru/api/remap/1.2/download/%s", imageId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		r.logger.Error(err)
		return []byte{}, err
	}

	resp, err := r.moyskladClient.Do(req)
	if err != nil {
		r.logger.Error(err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err)
		return []byte{}, err
	}

	return data, nil
}

package handler

import (
	"nashimenshie_api/internal/service"
	"nashimenshie_api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	logger  *logging.Logger
}

func NewHandler(service *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	assortment := router.Group("/assortment")
	{
		assortment.GET("/", h.getAssortment)
	}

	multimedia := router.Group("/multimedia")
	{
		multimedia.GET("/image/:id", h.getImage)
	}

	return router
}

package handler

import (
	"nashimenshie_api/pkg/apperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getImage(c *gin.Context) {
	var size int
	var err error

	if c.Query("size") != "" {
		if size, err = strconv.Atoi(c.Query("size")); err != nil {
			appError := apperror.NewAppError(err, "incorrect value of the \"size\" parameter")
			h.logger.Error(appError.Unwrap())
			newErrorResponse(c, http.StatusOK, appError.Error())
			return
		}
	}

	image, err := h.service.Multimedia.DownloadImage(c.Param("id"), uint(size))
	if err != nil {
		h.logger.Error(err)
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	c.Data(http.StatusOK, "image/png", image)
}

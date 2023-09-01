package handler

import (
	"fmt"
	"nashimenshie_api/pkg/apperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAssortment(c *gin.Context) {
	var limit, offset int
	var err error

	if c.Query("limit") == "" {
		err := fmt.Errorf("missing required parameter \"limit\"")
		h.logger.Error(err)
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if limit, err = strconv.Atoi(c.Query("limit")); err != nil {
		appError := apperror.NewAppError(err, "incorrect value of the \"limit\" parameter")
		h.logger.Error(appError.Unwrap())
		newErrorResponse(c, http.StatusOK, appError.Error())
		return
	}

	if limit < 1 {
		err = fmt.Errorf("incorrect value of the \"limit\" parameter")
		h.logger.Error(err)
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if c.Query("offset") != "" {
		if offset, err = strconv.Atoi(c.Query("offset")); err != nil {
			appError := apperror.NewAppError(err, "incorrect value of the \"offset\" parameter")
			h.logger.Error(appError.Unwrap())
			newErrorResponse(c, http.StatusOK, appError.Error())
			return
		}
	}

	assortment, err := h.service.Assortment.GetAssortment(limit, offset)
	if err != nil {
		h.logger.Error(err)
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	newOkResponse(c, assortment)
}

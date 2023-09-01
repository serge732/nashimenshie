package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newOkResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok":    true,
		"data":  data,
		"error": nil,
	})
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, map[string]interface{}{
		"ok":    false,
		"data":  nil,
		"error": message,
	})
}

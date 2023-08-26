package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	reasonInvalidArgument = "invalid_argument"
)

type response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func successResponse(c *gin.Context, data string, message string) {
	obj := &response{
		Success: true,
		Data:    data,
		Message: message,
	}
	c.JSON(http.StatusOK, obj)
}

func invalidArgumentResponse(c *gin.Context, message string) {
	obj := &response{
		Success: false,
		Data:    "",
		Reason:  reasonInvalidArgument,
		Message: message,
	}
	c.JSON(http.StatusBadRequest, obj)
}

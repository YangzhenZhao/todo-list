package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	reasonInvalidArgument     = "invalid_argument"
	reasonUnauthorized        = "unauthorized"
	reasonInternalServerError = "internall_server_error"
)

type response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func SuccessResponse(c *gin.Context, data string, message string) {
	obj := &response{
		Success: true,
		Data:    data,
		Message: message,
	}
	c.JSON(http.StatusOK, obj)
}

func InvalidArgumentResponse(c *gin.Context, message string) {
	obj := &response{
		Success: false,
		Data:    "",
		Reason:  reasonInvalidArgument,
		Message: message,
	}
	c.JSON(http.StatusBadRequest, obj)
}

func InternalServerErrorResponse(c *gin.Context, message string) {
	obj := &response{
		Success: false,
		Data:    "",
		Reason:  reasonInternalServerError,
		Message: message,
	}
	c.JSON(http.StatusInternalServerError, obj)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	obj := &response{
		Success: false,
		Data:    "",
		Reason:  reasonUnauthorized,
		Message: message,
	}
	c.JSON(http.StatusUnauthorized, obj)
}

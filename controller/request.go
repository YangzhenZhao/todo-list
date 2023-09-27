package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/YangzhenZhao/todo-list/common/response"
	"github.com/gin-gonic/gin"
)

func unmarshalRequest(c *gin.Context, v any) error {
	if c == nil || c.Request == nil {
		return errors.New("invalid request")
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func validteToken(c *gin.Context, reqUserID uint, funcName string) bool {
	tokenUserID, exists := c.Get("userID")
	if !exists {
		logger.Info(funcName, " token don't have userID")
		response.InvalidArgumentResponse(c, "token不合法")
		return false
	}

	if reqUserID != tokenUserID {
		logger.Info(funcName, " invalid userID, reqUserID, ", reqUserID, "tokenUserID", tokenUserID)
		response.InvalidArgumentResponse(c, "token不合法")
		return false
	}

	return true
}

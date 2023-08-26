package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"

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

package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/logic/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	registerReq := &dto.RegisterRequest{}
	err = json.Unmarshal(body, registerReq)
	log.Printf("register request... body: %s", string(body))
	if err != nil {
		log.Printf("register request... err: %v\n", err)
		return
	}
	log.Printf("register request... email:%s, password:%s\n", registerReq.Email, registerReq.Password)
	err = user.CreateUser(registerReq.Email, registerReq.Password)
	if err != nil {
		log.Printf("[Register] create user err: %v\n", err)
	}
	c.JSON(200, "")
}

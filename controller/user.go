package controller

import (
	"log"

	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/logic/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerReq := &dto.RegisterRequest{}
	if err := unmarshalRequest(c, registerReq); err != nil {
		log.Printf("register request... err: %v\n", err)
		invalidArgumentResponse(c, "注册参数不合法")
		return
	}

	err := user.CreateUser(registerReq.Email, registerReq.Password)
	if err != nil {
		log.Printf("[Register] create user err: %v\n", err)
		c.JSON(400, err.Error())
		return
	}
	successResponse(c, "", "注册成功")
}

func Login(c *gin.Context) {
	loginReq := &dto.LoginRequest{}
	if err := unmarshalRequest(c, loginReq); err != nil {
		log.Printf("login request... err: %v\n", err)
		invalidArgumentResponse(c, "登录参数不合法")
		return
	}

	userID, err := user.Login(loginReq)
	if err != nil {
		invalidArgumentResponse(c, "账号或密码不正确")
		return
	}
	res := &dto.LoginResponse{
		UserID: userID,
	}
	successResponse(c, res.JsonDumps(), "登录成功")
}

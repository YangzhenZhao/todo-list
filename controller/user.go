package controller

import (
	"github.com/YangzhenZhao/todo-list/common/log"
	"github.com/YangzhenZhao/todo-list/common/response"
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/logic/user"
	"github.com/gin-gonic/gin"
)

var logger = log.NewLogger()

func Register(c *gin.Context) {
	registerReq := &dto.RegisterRequest{}
	if err := unmarshalRequest(c, registerReq); err != nil {
		logger.Info("register request... err: %v\n", err)
		response.InvalidArgumentResponse(c, "注册参数不合法")
		return
	}

	err := user.CreateUser(registerReq.Email, registerReq.Password)
	if err != nil {
		logger.Info("[Register] create user err: ", err)
		c.JSON(400, err.Error())
		return
	}
	response.SuccessResponse(c, "", "注册成功")
}

func Login(c *gin.Context) {
	loginReq := &dto.LoginRequest{}
	if err := unmarshalRequest(c, loginReq); err != nil {
		logger.Info("login request... err: ", err)
		response.InvalidArgumentResponse(c, "登录参数不合法")
		return
	}

	loginRes, err := user.Login(loginReq)
	if err != nil {
		response.InvalidArgumentResponse(c, "账号或密码不正确")
		return
	}
	response.SuccessResponse(c, loginRes.JsonDumps(), "登录成功")
}

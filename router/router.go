package router

import (
	"github.com/YangzhenZhao/todo-list/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouters() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/todos", controller.GetTodolist)

	addUserRouters(v1)
	addTodoRouters(v1)
}

func Run() {
	router.Run(":5000")
}

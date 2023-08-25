package router

import (
	"github.com/YangzhenZhao/todo-list/controller"
	"github.com/gin-gonic/gin"
)

func addUserRouters(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.POST("/register", controller.Register)
}

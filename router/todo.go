package router

import (
	"github.com/YangzhenZhao/todo-list/controller"
	"github.com/YangzhenZhao/todo-list/middleware"
	"github.com/gin-gonic/gin"
)

func addTodoRouters(rg *gin.RouterGroup) {
	users := rg.Group("/todo")

	users.GET("", middleware.AuthJWTMiddleware, controller.GetTodo)
	users.POST("", middleware.AuthJWTMiddleware, controller.InsertTodo)
	users.PUT("", middleware.AuthJWTMiddleware, controller.UpdateTodo)
	users.DELETE("", middleware.AuthJWTMiddleware, controller.DeleteTodo)
}

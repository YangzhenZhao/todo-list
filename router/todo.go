package router

import (
	"github.com/YangzhenZhao/todo-list/controller"
	"github.com/YangzhenZhao/todo-list/middleware"
	"github.com/gin-gonic/gin"
)

func addTodoRouters(rg *gin.RouterGroup) {
	users := rg.Group("/todo")

	users.GET("", middleware.AuthJWTMiddleware, controller.GetTodo)
	users.POST("", middleware.AuthJWTMiddleware, controller.CreateTodo)
	users.PUT("", middleware.AuthJWTMiddleware, controller.UpdateTodo)
	users.DELETE("", middleware.AuthJWTMiddleware, controller.DeleteTodo)

	users.POST("star", middleware.AuthJWTMiddleware, controller.StarTodo)
	users.POST("unstar", middleware.AuthJWTMiddleware, controller.UnstarTodo)
}

package router

import (
	"github.com/YangzhenZhao/todo-list/common/log"
	"github.com/YangzhenZhao/todo-list/common/monitor"
	"github.com/YangzhenZhao/todo-list/controller"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var router *gin.Engine

func InitRouters() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.NewLogger().Out

	router = gin.Default()

	registry := monitor.NewMonitoryRegistry()
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	)))

	v1 := router.Group("/v1")

	v1.GET("/todos", controller.GetTodolist)

	addUserRouters(v1)
	addTodoRouters(v1)
}

func Run() {
	router.Run(":5000")
}

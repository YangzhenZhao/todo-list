package router

import "github.com/gin-gonic/gin"

var router *gin.Engine

func InitRouters() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	v1 := router.Group("/v1")
	addUserRouters(v1)
}

func Run() {
	router.Run(":5000")
}

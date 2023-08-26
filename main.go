package main

import (
	"github.com/YangzhenZhao/todo-list/dao"
	"github.com/YangzhenZhao/todo-list/router"
)

func main() {
	dao.Init()
	router.InitRouters()

	router.Run()
}

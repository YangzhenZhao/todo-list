package main

import (
	"github.com/YangzhenZhao/todo-list/common/log"
	"github.com/YangzhenZhao/todo-list/dao"
	"github.com/YangzhenZhao/todo-list/router"
)

func main() {
	log.InitLogger()
	dao.Init()
	router.InitRouters()

	router.Run()
}

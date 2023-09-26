package dao

import (
	"github.com/YangzhenZhao/todo-list/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger = log.NewLogger()

func Init() {
	dsn := "root:woaini123@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.Error(err)
		panic("数据库连接失败!")
	}
	UserDao = NewUserDao(db)
	TodoDao = NewTodoDao(db)
}

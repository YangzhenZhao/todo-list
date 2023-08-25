package user

import (
	"github.com/YangzhenZhao/todo-list/dao"
)

func CreateUser(email, password string) error {
	return dao.UserDao.CreateUser(email, password)
}
